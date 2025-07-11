package kafka

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"os"
	"strings"

	"github.com/Shopify/sarama"

	keystore "github.com/pavlo-v-chernykh/keystore-go/v4"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

/*
send message and return partition,offset
*/
func (t *KafkaProducer) SendMesssage(content []byte, topic string, keyType string) (int32, int64, error) {
	message := &sarama.ProducerMessage{}
	message.Topic = topic
	message.Key = sarama.StringEncoder(keyType)
	message.Value = sarama.ByteEncoder(content)
	partition, offset, err := t.producer.SendMessage(message)
	return partition, offset, err
}

func InitKafka(servers []string) KafkaProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	var err error
	producer, err := sarama.NewSyncProducer(servers, config)
	if err != nil {
		panic(err)
	}
	return KafkaProducer{producer: producer}
}

// 📁 d:\idea project\udpgo\kafka\kafkaClient.go
func InitKafkaWithAuth(
	servers []string,
	saslUser string,
	saslPassword string,
	saslMechanism string,
	keystorePath string,
	keystorePassword string,
	truststorePath string,
) KafkaProducer {
	config := sarama.NewConfig()

	// 基础生产者配置 (原有逻辑保留)
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	// SASL认证配置 (新增)
	if saslUser != "" && saslPassword != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = saslUser
		config.Net.SASL.Password = saslPassword
		switch saslMechanism {
		case "SCRAM-SHA-256":
			config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256
		case "SCRAM-SHA-512":
			config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
		default:
			config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
		}
	}

	// SSL配置（JKS读取）
	if keystorePath != "" && truststorePath != "" && keystorePassword != "" {
		// 读取密码
		pwBytes, err := os.ReadFile(keystorePassword)
		if err != nil {
			panic(err)
		}
		password := strings.TrimSpace(string(pwBytes))

		// 读取 keystore.jks
		ksFile, err := os.ReadFile(keystorePath)
		if err != nil {
			panic(err)
		}
		ks := keystore.New()
		if err := ks.Load(bytes.NewReader(ksFile), []byte(password)); err != nil {
			panic(err)
		}
		// 获取第一个私钥别名
		var keyAlias string
		for _, alias := range ks.Aliases() {
			if _, err := ks.GetPrivateKeyEntry(alias, []byte(password)); err == nil {
				keyAlias = alias
				break
			}
		}
		if keyAlias == "" {
			panic("no private key entry found in keystore")
		}
		privKeyEntry, err := ks.GetPrivateKeyEntry(keyAlias, []byte(password))
		if err != nil {
			panic(err)
		}
		// 转为PEM
		privKeyBlock := pem.Block{Type: "PRIVATE KEY", Bytes: privKeyEntry.PrivateKey}
		privKeyPEM := pem.EncodeToMemory(&privKeyBlock)
		certPEM := []byte{}
		for _, cert := range privKeyEntry.CertificateChain {
			certBlock := pem.Block{Type: "CERTIFICATE", Bytes: cert.Content}
			certPEM = append(certPEM, pem.EncodeToMemory(&certBlock)...)
		}

		// 读取 truststore.jks
		trustFile, err := os.ReadFile(truststorePath)
		if err != nil {
			panic(err)
		}
		ts := keystore.New()
		if err := ts.Load(bytes.NewReader(trustFile), []byte(password)); err != nil {
			panic(err)
		}
		caPEM := []byte{}
		for _, alias := range ts.Aliases() {
			entry, _ := ts.GetTrustedCertificateEntry(alias)
			if len(entry.Certificate.Content) > 0 {
				caBlock := pem.Block{Type: "CERTIFICATE", Bytes: entry.Certificate.Content}
				caPEM = append(caPEM, pem.EncodeToMemory(&caBlock)...)
			}
		}

		// 构造 tls.Certificate
		keyPair, err := tls.X509KeyPair(certPEM, privKeyPEM)
		if err != nil {
			panic(err)
		}
		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(caPEM)

		config.Net.TLS.Enable = true
		config.Net.TLS.Config = &tls.Config{
			Certificates:       []tls.Certificate{keyPair},
			RootCAs:            certPool,
			InsecureSkipVerify: false,
		}
	}

	producer, err := sarama.NewSyncProducer(servers, config)
	if err != nil {
		panic(err)
	}
	return KafkaProducer{producer: producer}
}
