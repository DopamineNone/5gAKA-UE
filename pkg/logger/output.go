package logger

import "go.uber.org/zap"

func (logger *Logger) Debug(msg, source string, fields ...zap.Field) {
	logger.Logger.Debug(msg, append([]zap.Field{zap.String("Source", source)}, fields...)...)
}

func (logger *Logger) Info(msg, source string, fields ...zap.Field) {
	logger.Logger.Info(msg, append([]zap.Field{zap.String("Source", source)}, fields...)...)
}

func (logger *Logger) Warn(msg, source string, fields ...zap.Field) {
	logger.Logger.Warn(msg, append([]zap.Field{zap.String("Source", source)}, fields...)...)
}

func (logger *Logger) Error(msg, source string, err error, fields ...zap.Field) {
	logger.Logger.Error(msg, append([]zap.Field{zap.String("Source", source), zap.Error(err)}, fields...)...)
}
