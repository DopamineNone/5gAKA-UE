package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	LoggerLevel zapcore.LevelEnabler = zap.DebugLevel
)

type LogFileConfig struct {
	FileName     string
	MaxMegaBytes int
	MaxBackups   int
	MaxAge       int
}

func InitLogger(cfg ...*LogFileConfig) *zap.Logger {
	var core zapcore.Core = InitConsoleLoggerCore()
	if len(cfg) > 0 {
		for _, config := range cfg {
			core = zapcore.NewTee(
				core,
				InitFileLoggerCore(config),
			)
		}
	}
	return zap.New(core, zap.AddCaller())
}

func InitConsoleLoggerCore() zapcore.Core {
	EncoderConfig := zap.NewProductionEncoderConfig()
	EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewCore(zapcore.NewConsoleEncoder(EncoderConfig), zapcore.Lock(os.Stdout), LoggerLevel)
}

func InitFileLoggerCore(cfg *LogFileConfig) zapcore.Core {
	WriteSyncer := &lumberjack.Logger{
		Filename:   cfg.FileName,
		MaxSize:    cfg.MaxMegaBytes,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
	}
	EncoderConfig := zap.NewProductionEncoderConfig()
	EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewCore(zapcore.NewJSONEncoder(EncoderConfig), zapcore.AddSync(WriteSyncer), LoggerLevel)
}
