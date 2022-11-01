package kafka

import (
	"github.com/Shopify/sarama"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func (t *KafkaProducer) SendMesssage(content []byte, topic string, keyType string) error {
	message := &sarama.ProducerMessage{}
	message.Topic = topic
	message.Key = sarama.StringEncoder(keyType)
	message.Value = sarama.ByteEncoder(content)
	_, _, err := t.producer.SendMessage(message)
	return err
}

func initKafka(servers []string) KafkaProducer {
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
