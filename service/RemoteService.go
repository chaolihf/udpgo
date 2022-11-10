package service

import (
	"context"

	remote "github.com/chaolihf/udpgo/com.chinatelecom.udp.grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CallToken struct {
	token string
}

func (c CallToken) GetRequestMetadata(ctx context.Context,
	uri ...string) (map[string]string, error) {
	return map[string]string{
		"token": c.token,
	}, nil
}

func (c CallToken) RequireTransportSecurity() bool {
	return false
}

type RemoteClient struct {
	client    remote.DataServiceClient
	Token     string
	authToken CallToken
}

func (c *RemoteClient) Connect(address string) error {
	c.authToken = CallToken{token: c.Token}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(c.authToken))
	if err != nil {
		return err
	}
	c.client = remote.NewDataServiceClient(conn)
	return nil
}

func (c *RemoteClient) CallService(code string, method string, version uint32, sequence uint32, data []byte) (string, error) {
	request := remote.ServiceRequest{}
	request.ServiceName = code
	request.MethodName = method
	request.Version = version
	request.Sequence = sequence
	request.Data = data
	response, err := c.client.CallService(context.Background(), &request)
	if err != nil {
		return "", err
	}
	return string(response.Data), nil
}
