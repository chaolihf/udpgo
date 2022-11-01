package kafka

import "testing"

func TestSending(t *testing.T) {
	p := initKafka([]string{"127.0.0.1:8450"})
	p.SendMesssage([]byte(""), "aa", "key")
}
