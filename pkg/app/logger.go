package app

import "go.uber.org/zap"

type LoggerOption struct {
}

func createAppLogger(option *LoggerOption) *zap.Logger {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", "log/ue.log"}
	config.ErrorOutputPaths = []string{"log/error.log"}

	if logger, err := config.Build(); err == nil {
		return logger
	} else {
		return nil
	}
}

func syncLogger(logger *zap.Logger) {
	if err := logger.Sync(); err != nil {
		logger.Error("Fail to sync local logger when app closed")
	}
}
