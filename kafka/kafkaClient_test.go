package kafka

import "testing"

func TestSending(t *testing.T) {
	initKafka([]string{"127.0.0.1:8450"})
}
