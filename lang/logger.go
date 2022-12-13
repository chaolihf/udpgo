package lang

import (
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FormatStringer struct {
	parts []*StringPart
}

type StringPart struct {
	partData   []rune
	isParam    bool
	newValue   string
	isNewValue bool
}

// parse content to parts
func (f *FormatStringer) Parse(content string) {
	escapeMode := false
	paramMode := false
	allParts := []*StringPart{}
	currentPart := &StringPart{}
	for _, char := range content {
		switch char {
		case '\\':
			{
				if escapeMode {
					currentPart.partData = append(currentPart.partData, '\\')
					escapeMode = false
				} else {
					escapeMode = true
				}
			}
		case '{':
			{
				if escapeMode {
					currentPart.partData = append(currentPart.partData, '\\')
					escapeMode = false
				} else {
					currentPart.isParam = false
					if len(currentPart.partData) > 0 {
						allParts = append(allParts, currentPart)
						currentPart = &StringPart{}
					}
					paramMode = true
				}
			}
		case '}':
			{
				if escapeMode {
					currentPart.partData = append(currentPart.partData, '\\')
					escapeMode = false
				} else {
					if paramMode {
						currentPart.isParam = true
						allParts = append(allParts, currentPart)
						currentPart = &StringPart{}
						paramMode = false
					} else {
						currentPart.partData = append(currentPart.partData, '}')
					}
				}
			}
		default:
			{
				currentPart.partData = append(currentPart.partData, char)
			}
		}

	}
	currentPart.isParam = false
	if len(currentPart.partData) > 0 {
		allParts = append(allParts, currentPart)
	}
	f.parts = allParts
}

func (f *FormatStringer) GetParamNames() []string {
	result := []string{}
	for _, part := range f.parts {
		if part.isParam {
			result = append(result, string(part.partData))
		}
	}
	return result
}

func (f *FormatStringer) String() string {
	buffer := strings.Builder{}
	for _, part := range f.parts {
		if !part.isParam {
			buffer.WriteString(string(part.partData))
		} else if part.isNewValue {
			buffer.WriteString(string(part.newValue))
		} else {
			buffer.WriteRune('{')
			buffer.WriteString(string(part.partData))
			buffer.WriteRune('}')
		}
	}
	return buffer.String()
}

func (f *FormatStringer) Replace(paramName string, index int, newValue string) {
	for _, part := range f.parts {
		if part.isParam && string(part.partData) == paramName {
			if index == 0 {
				part.newValue = newValue
				part.isNewValue = true
				return
			}
			index--
		}
	}
}

func (f *FormatStringer) ReplaceAll(paramName string, newValue string) {
	for _, part := range f.parts {
		if part.isParam && string(part.partData) == paramName {
			part.newValue = newValue
			part.isNewValue = true
		}
	}
}

func InitLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}

// @param logFileName 日志文件存放目录
// param maxSize 文件大小限制,单位MB
// param maxBackups 最大保留日志文件数量
// param maxAge 日志文件保留天数
func InitProductLogger(logFileName string, maxSize int, maxBackups int, maxAge int) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   false,
	})
	fileCore := zapcore.NewCore(encoder, fileWriteSyncer, zapcore.InfoLevel)
	logger := zap.New(fileCore, zap.AddCaller())
	return logger
}

// @title return left string ,support utf-8 and without panic when outbound
func LeftString(data string, count int) string {
	if len(data) <= count {
		return data
	}
	utfData := []rune(data)
	if len(utfData) <= count {
		return data
	} else {
		return string(utfData[0:count])
	}
}
