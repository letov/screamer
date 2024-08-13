package logger

import "go.uber.org/zap"

var sugar zap.SugaredLogger

func Init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	sugar = *logger.Sugar()
}

func GetSugar() *zap.SugaredLogger {
	return &sugar
}
