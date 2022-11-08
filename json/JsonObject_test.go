package jjson

import (
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	data := `{"id":524042,"name":"酷旅-mob-otv-2","male":true,"other":null}`
	object, err := FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		//object.GetJsonObject("id").Value = "aaaa"
		fmt.Println(object.String())
	}
}
