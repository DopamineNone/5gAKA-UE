package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

type WriterConfig struct {
	FileName     string
	MaxMegaBytes int
	MaxBackups   int
	MaxAge       int
}

func NewConfig() *WriterConfig {
	return &WriterConfig{
		FileName:     "output.log",
		MaxAge:       30,
		MaxBackups:   30,
		MaxMegaBytes: 10,
	}
}

func (cfg *WriterConfig) SetFileName(filename string) *WriterConfig {
	cfg.FileName = filename
	return cfg
}

func (cfg *WriterConfig) SetMaxAge(input int) *WriterConfig {
	cfg.MaxAge = input
	return cfg
}

func (cfg *WriterConfig) SetMaxBackups(input int) *WriterConfig {
	cfg.MaxBackups = input
	return cfg
}

func (cfg *WriterConfig) SetMaxMegaBytes(input int) *WriterConfig {
	cfg.MaxMegaBytes = input
	return cfg
}

func (cfg *WriterConfig) Build() io.Writer {
	return &lumberjack.Logger{
		Filename:   cfg.FileName,
		MaxSize:    cfg.MaxMegaBytes,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
	}
}

type Logger struct {
	*zap.Logger
}

func NewLogger(syncWriter ...io.Writer) *Logger {
	var core = initConsoleLoggerCore()
	for _, writer := range syncWriter {
		core = zapcore.NewTee(
			core,
			initFileLoggerCore(writer),
		)
	}
	return &Logger{
		Logger: zap.New(core, zap.AddCaller()),
	}
}

func initConsoleLoggerCore() zapcore.Core {
	EncoderConfig := zap.NewProductionEncoderConfig()
	EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewCore(zapcore.NewConsoleEncoder(EncoderConfig), zapcore.Lock(os.Stdout), zap.DebugLevel)
}

func initFileLoggerCore(syncWriter io.Writer) zapcore.Core {
	EncoderConfig := zap.NewProductionEncoderConfig()
	EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewCore(zapcore.NewJSONEncoder(EncoderConfig), zapcore.AddSync(syncWriter), zap.InfoLevel)
}
