package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	logger *Logger
)

type Config struct {
	FileName     string
	MaxMegaBytes int
	MaxBackups   int
	MaxAge       int
}

func InitLogger(cfg ...*Config) *Logger {
	var core = InitConsoleLoggerCore()
	if len(cfg) > 0 {
		for _, config := range cfg {
			core = zapcore.NewTee(
				core,
				InitFileLoggerCore(config),
			)
		}
	}
	return &Logger{
		Logger: zap.New(core, zap.AddCaller()),
	}
}

func InitConsoleLoggerCore() zapcore.Core {
	EncoderConfig := zap.NewProductionEncoderConfig()
	EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewCore(zapcore.NewConsoleEncoder(EncoderConfig), zapcore.Lock(os.Stdout), Level)
}

func InitFileLoggerCore(cfg *Config) zapcore.Core {
	WriteSyncer := &lumberjack.Logger{
		Filename:   cfg.FileName,
		MaxSize:    cfg.MaxMegaBytes,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
	}
	EncoderConfig := zap.NewProductionEncoderConfig()
	EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewCore(zapcore.NewJSONEncoder(EncoderConfig), zapcore.AddSync(WriteSyncer), Level)
}

type Logger struct {
	*zap.Logger
}

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

func (logger *Logger) Stop() {
	_ = logger.Sync()
}
