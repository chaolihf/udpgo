package influxdb

import (
	"fmt"
	"testing"
	"time"
)

func TestSending(t *testing.T) {
	currentTime := time.Now().UnixMilli()
	line, err := GetLine("a", []TagType{
		{Key: "t1", Value: "t1Value"},
		{Key: "t2", Value: "t2Value"},
	},
		[]FieldType{
			{Key: "f1", Value: "fieldValue1"},
			{Key: "f2", Value: 1},
			{Key: "f3", Value: 1.12},
			{Key: "f4", Value: true},
		}, currentTime)
	if err != nil {
		t.Error(err.Error())
	}
	expected := fmt.Sprintf("a,t1=t1Value,t2=t2Value f1=\"fieldValue1\",f2=1i,f3=1.12000,f4=t %d", currentTime)
	if expected != line {
		t.Error("content is not equal")
	}
	t.Log("OK")
}
