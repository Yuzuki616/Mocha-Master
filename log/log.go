package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	c := zap.NewProductionConfig()
	c.EncoderConfig.TimeKey = "time"
	c.EncoderConfig.CallerKey = ""
	c.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ = c.Build()
	logger = logger.Named("Main")
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
