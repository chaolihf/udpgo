package jjson

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type JsonObject struct {
	Attributes map[string]*JsonObject
	Value      interface{}
	VType      reflect.Kind
}

func (j *JsonObject) getJsonArray() []*JsonObject {
	if j == nil && j.VType != reflect.Slice {
		return nil
	}
	return j.Value.([]*JsonObject)
}

/*
new json object from bytes
*/
func NewJsonObject(data []byte) (*JsonObject, error) {
	return FromBytes(data)
}

/*
new json array from bytes
*/
func NewJsonArray(data []byte) ([]*JsonObject, error) {
	object, err := FromBytes(data)
	return object.getJsonArray(), err
}

func (j *JsonObject) GetJsonArray(key string) []*JsonObject {
	jsonArray := j.GetJsonObject(key)
	if jsonArray != nil {
		return jsonArray.getJsonArray()
	} else {
		return nil
	}
}

func (j *JsonObject) New() *JsonObject {
	j.VType = reflect.Struct
	return j
}

func (j *JsonObject) GetString(key string) string {
	if j == nil {
		return ""
	}
	itemValue := j.Attributes[key]
	if itemValue == nil || itemValue.VType != reflect.String {
		return ""
	}
	return itemValue.Value.(string)
}

func (j *JsonObject) PutString(key string, value string) {
	item := j.GetJsonObject(key)
	if item != nil {
		item.VType = reflect.String
		item.Value = value
	} else {
		j.Attributes[key] = &JsonObject{VType: reflect.String, Value: value}
	}
}

func (j *JsonObject) GetInt(key string) int {
	if j == nil {
		return 0
	}
	itemValue := j.Attributes[key]
	if itemValue == nil {
		return 0
	}
	switch itemValue.VType {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return int(itemValue.Value.(int64))
	default:
		return 0
	}

}

func (j *JsonObject) PutInt(key string, value int) {
	item := j.GetJsonObject(key)
	if item != nil {
		item.VType = reflect.Int
		item.Value = value
	} else {
		j.Attributes[key] = &JsonObject{VType: reflect.Int, Value: value}
	}
}

func (j *JsonObject) GetLong(key string) int64 {
	if j == nil {
		return 0
	}
	itemValue := j.Attributes[key]
	if itemValue == nil || (itemValue.VType != reflect.Int64 && itemValue.VType != reflect.Uint64) {
		return 0
	}
	return itemValue.Value.(int64)
}

func (j *JsonObject) PutLong(key string, value int64) {
	item := j.GetJsonObject(key)
	if item != nil {
		item.VType = reflect.Int64
		item.Value = value
	} else {
		j.Attributes[key] = &JsonObject{VType: reflect.Int64, Value: value}
	}
}

func (j *JsonObject) GetFloat(key string) float64 {
	if j == nil {
		return 0
	}
	itemValue := j.Attributes[key]
	if itemValue == nil || itemValue.VType != reflect.Float64 {
		return 0
	}
	return itemValue.Value.(float64)
}

func (j *JsonObject) PutFloat(key string, value float64) {
	item := j.GetJsonObject(key)
	if item != nil {
		item.VType = reflect.Float64
		item.Value = value
	} else {
		j.Attributes[key] = &JsonObject{VType: reflect.Float64, Value: value}
	}
}

func (j *JsonObject) GetBool(key string) bool {
	if j == nil {
		return false
	}
	itemValue := j.Attributes[key]
	if itemValue == nil || itemValue.VType != reflect.Bool {
		return false
	}
	return itemValue.Value.(bool)
}

func (j *JsonObject) PutBool(key string, value bool) {
	item := j.GetJsonObject(key)
	if item != nil {
		item.VType = reflect.Bool
		item.Value = value
	} else {
		j.Attributes[key] = &JsonObject{VType: reflect.Bool, Value: value}
	}
}

func (j *JsonObject) GetJsonObject(key string) *JsonObject {
	if j == nil {
		return nil
	}
	return j.Attributes[key]
}

func (j *JsonObject) PutJsonObject(key string, value *JsonObject) {
	item := j.GetJsonObject(key)
	if item != nil {
		item.VType = reflect.Struct
		item.Value = nil
		item.Attributes[key] = value
	} else {
		j.Attributes[key] = &JsonObject{VType: reflect.Struct, Value: value}
	}
}

func (j *JsonObject) PutJsonArray(key string, value []JsonObject) {
	item := j.GetJsonObject(key)
	if item != nil {
		item.VType = reflect.Slice
		item.Value = value
	} else {
		j.Attributes[key] = &JsonObject{VType: reflect.Slice, Value: value}
	}
}

func (j *JsonObject) GetKeys() []string {
	var keys = []string{}
	for key, _ := range j.Attributes {
		keys = append(keys, key)
	}
	return keys
}

func (j *JsonObject) String() string {
	var buffer strings.Builder
	switch j.VType {
	case reflect.Slice:
		{
			items := reflect.ValueOf(j.Value)
			buffer.WriteString("[")
			for i := 0; i < items.Len(); i++ {
				item := items.Index(i)
				if item.Kind() == reflect.Ptr {
					childJson, _ := item.Elem().Interface().(JsonObject)
					buffer.WriteString(childJson.String())
				} else {
					buffer.WriteString(encodeValue(item))
				}
				buffer.WriteString(",")
			}
			if buffer.Len() > 1 {
				return buffer.String()[0:buffer.Len()-1] + "]"
			} else {
				return "[]"
			}
		}
	case reflect.Struct:
		{
			buffer.WriteString("{")
			for key, value := range j.Attributes {
				buffer.WriteString("\"")
				buffer.WriteString(key)
				buffer.WriteString("\"")
				buffer.WriteString(":")
				buffer.WriteString(value.String())
				buffer.WriteString(",")
			}
			result := buffer.String()
			if result[len(result)-1:] == "," {
				return result[0:len(result)-1] + "}"
			} else {
				return result + "}"
			}
		}
	default:
		{
			return encodeValue(j.Value)
		}
	}
}

func encodeValue(j interface{}) string {
	v := reflect.ValueOf(j)
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return "\"" + v.String() + "\""
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", v.Float())
	case reflect.Interface, reflect.Ptr:
		return "null"
	}
	return "null"
}

func ArrayString(jsonArray []*JsonObject) string {
	var buffer strings.Builder
	buffer.WriteString("[")
	for _, jsonObject := range jsonArray {
		buffer.WriteString(jsonObject.String())
		buffer.WriteString(",")
	}
	if buffer.Len() > 1 {
		return buffer.String()[0:buffer.Len()-1] + "]"
	} else {
		return "[]"
	}
}
