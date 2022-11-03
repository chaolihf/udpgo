package lang

import "go.uber.org/zap"

func InitLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}

//@title return left string ,support utf-8 and without panic when outbound
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
