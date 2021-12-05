package logger

import "go.uber.org/zap"

func NewLogger(context string) *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.InitialFields = map[string]interface{}{
		"context": context,
	}
	zapLogger, _ := config.Build()

	defer zapLogger.Sync()

	return zapLogger.Sugar()
}
