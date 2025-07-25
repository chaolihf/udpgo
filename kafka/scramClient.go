package kafka

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"os"

	"github.com/xdg/scram"
)

var (
	SHA256 scram.HashGeneratorFcn = func() hash.Hash { return sha256.New() }
	SHA512 scram.HashGeneratorFcn = func() hash.Hash { return sha512.New() }
)

type XDGSCRAMClient struct {
	*scram.Client
	*scram.ClientConversation
	scram.HashGeneratorFcn
}

func (x *XDGSCRAMClient) Begin(userName, password, authzID string) (err error) {
	x.Client, err = x.HashGeneratorFcn.NewClient(userName, password, authzID)
	if err != nil {
		return err
	}
	x.ClientConversation = x.Client.NewConversation()
	return nil
}

func (x *XDGSCRAMClient) Step(challenge string) (response string, err error) {
	return x.ClientConversation.Step(challenge)
}

func (x *XDGSCRAMClient) Done() bool {
	return x.ClientConversation.Done()
}

func readFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("文件读取失败: " + path)
	}
	return data
}
