package lang

import "go.uber.org/zap"

func InitLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}
