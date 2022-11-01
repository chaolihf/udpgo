package lang

import "go.uber.org/zap"

var logger *zap.Logger

func InitLogger() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
}
