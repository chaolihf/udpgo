package service

import (
	"fmt"
	"testing"
)

func TestCallService(t *testing.T) {
	client := RemoteClient{Token: ""}
	client.Connect("localhost:8803")
	content, err := client.CallService("workflowExecutor", "add", 0, 0, []byte("{}"))
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(string(content))
}
