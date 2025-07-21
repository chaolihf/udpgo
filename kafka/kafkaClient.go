package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"github.com/Shopify/sarama"
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
			config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
				return &XDGSCRAMClient{HashGeneratorFcn: SHA256}
			}
		case "SCRAM-SHA-512":
			config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
			config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
				return &XDGSCRAMClient{HashGeneratorFcn: SHA512}
			}
		default:
			config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
		}
	}

	if keystorePath != "" && truststorePath != "" && keystorePassword != "" {

		// 加载证书和私钥
		keyPair, err := tls.LoadX509KeyPair(keystorePath, keystorePassword)
		if err != nil {
			panic("加载证书和私钥失败: " + err.Error())
		}

		// 加载CA证书
		caPEM, err := os.ReadFile(truststorePath)
		if err != nil {
			panic("加载CA证书失败: " + err.Error())
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
