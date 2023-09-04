package kafka

import (
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
