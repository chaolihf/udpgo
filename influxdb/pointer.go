package influxdb

import (
	"errors"
	"strconv"
	"strings"
)

type ParseState int32

const (
	Measurement ParseState = 0
	Tag         ParseState = 1
	Field       ParseState = 2
	Timestamp   ParseState = 3
)

type FieldValueType int32

const (
	StringType    FieldValueType = 0
	BooleanType   FieldValueType = 1
	UndefinedType FieldValueType = 2
	NumberType    FieldValueType = 3
)

const (
	Comma       rune = ','
	Space       rune = ' '
	EqualSign   rune = '='
	Backslash   rune = '/'
	DoubleQuote rune = '"'
)

type TagType struct {
	Key   string
	Value string
}

type FieldType struct {
	Key   string
	Value any
}

// @title 输出行数据
func GetLine(measureName string, tags []TagType, fields []FieldType, ptime int64) (string, error) {
	var line strings.Builder
	outputPart(&line, measureName, Measurement, false)
	for _, item := range tags {
		line.WriteRune(Comma)
		outputPart(&line, item.Key, Tag, false)
		line.WriteRune(EqualSign)
		outputPart(&line, item.Value, Tag, false)
	}
	line.WriteRune(Space)
	outputPart(&line, fields[0].Key, Field, false)
	line.WriteRune(EqualSign)
	outputPart(&line, fields[0].Value, Field, true)
	for _, item := range fields[1:] {
		line.WriteRune(Comma)
		outputPart(&line, item.Key, Field, false)
		line.WriteRune(EqualSign)
		outputPart(&line, item.Value, Field, true)
	}
	line.WriteRune(Space)
	outputPart(&line, ptime, Timestamp, true)
	return line.String(), nil
}

func outputPart(line *strings.Builder, value any, stage ParseState, isValue bool) error {
	switch stage {
	case Timestamp:
		{
			line.WriteString(strconv.FormatInt(value.(int64), 10))
		}
	case Measurement:
		{
			name := value.(string)
			for _, ch := range name {
				if ch == Comma || ch == Space {
					line.WriteRune(Backslash)
				}
				line.WriteRune(ch)
			}
		}
	default:
		{
			if stage == Field && isValue {
				switch v := value.(type) {
				case string:
					content := string(v)
					line.WriteRune(DoubleQuote)
					for _, ch := range content {
						if ch == DoubleQuote || ch == Backslash {
							line.WriteRune(Backslash)
						}
						line.WriteRune(ch)
					}
					line.WriteRune(DoubleQuote)
				case int:
					line.WriteString(strconv.Itoa(int(v)))
					line.WriteRune('i')
				case int64:
					line.WriteString(strconv.FormatInt(int64(v), 10))
					line.WriteRune('u')
				case float64:
					line.WriteString(strconv.FormatFloat(float64(v), 'f', 5, 64))
				case bool:
					if value.(bool) {
						line.WriteRune('t')
					} else {
						line.WriteRune('f')
					}
				default:
					return errors.New("error format")
				}
			} else {
				content := value.(string)
				for _, ch := range content {

					if ch == Comma || ch == Space || ch == EqualSign {
						line.WriteRune(Backslash)
					}
					line.WriteRune(ch)
				}
			}
			break
		}
	}
	return nil
}
