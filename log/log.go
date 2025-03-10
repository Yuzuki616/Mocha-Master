package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

var logger *zap.Logger

func init() {
	logger = newMainLogger(zapcore.InfoLevel)
}

func newMainLogger(level zapcore.Level) *zap.Logger {
	if logger != nil {
		logger.Sync()
	}
	c := zap.NewProductionConfig()
	c.EncoderConfig.TimeKey = "time"
	c.EncoderConfig.CallerKey = ""
	c.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ = c.Build()
	c.Level = zap.NewAtomicLevelAt(level)
	logger = logger.Named("Main")
	return logger
}

func SetLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		logger = newMainLogger(zap.DebugLevel)
	case "info":
		logger = newMainLogger(zap.InfoLevel)
	case "warn":
		logger = newMainLogger(zap.WarnLevel)
	case "error":
		logger = newMainLogger(zap.ErrorLevel)
	case "panic":
		logger = newMainLogger(zap.PanicLevel)
	}
}

func SubLogger(name string) *zap.Logger {
	return logger.Named(name)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
