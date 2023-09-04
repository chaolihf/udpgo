package jjson

import (
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	data := `{"id":1,"name":"中文","isOk":true,"f":1.2,"t":null}`
	object, _ := FromBytes([]byte(data))
	fmt.Println(object.String())

	json1 := new(JsonObject).New()
	fmt.Println(json1.String())
	json2 := new(JsonObject).New()
	fmt.Println(json2.String())
	fmt.Println(ArrayString([]*JsonObject{json1, json2}))
	if object.GetString("name") != "中文" {
		t.Fatal("name compare")
	}

}

func TestJsonArray(t *testing.T) {
	data := `[{"id":1,"name":"中文","isOk":true,"f":1.2,"t":null}]`
	array, _ := NewJsonArray([]byte(data))
	for _, item := range array {
		fmt.Println(item.GetInt("id"))
	}
}
