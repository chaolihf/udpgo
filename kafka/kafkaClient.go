package kafka

import (
	"github.com/Shopify/sarama"
)

var producer sarama.SyncProducer

func initKafka(servers []string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	var err error
	producer, err = sarama.NewSyncProducer(servers, config)
	if err != nil {
		panic(err)
	}
}

func SendMesssage(content []byte, topic string, keyType string) error {
	message := &sarama.ProducerMessage{}
	message.Topic = topic
	message.Key = sarama.StringEncoder(keyType)
	message.Value = sarama.ByteEncoder(content)
	_, _, err := producer.SendMessage(message)
	return err
}
