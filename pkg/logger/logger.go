package logger

import (
	"context"
	"sync"

	"github.com/mohdjishin/sportsphere/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once   sync.Once
	logger *LoggerClass
)

type LoggerClass struct {
	zapLogger *zap.Logger
}

func Run(ctx context.Context) {
	once.Do(func() {
		logLevel := getLogLevelFromEnv(config.Get().LogLevel)

		config := zap.Config{
			Level:            zap.NewAtomicLevelAt(logLevel),
			Development:      true,
			Encoding:         "json",
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "time",
				LevelKey:       "level",
				NameKey:        "logger",
				MessageKey:     "msg",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
		}

		zapLogger, err := config.Build()
		if err != nil {
			panic("failed to initialize logger: " + err.Error())
		}

		logger = &LoggerClass{zapLogger: zapLogger}

		go func() {
			<-ctx.Done()
			logger.Sync()
		}()
	})
}

func getLogLevelFromEnv(logLevel string) zapcore.Level {
	switch logLevel {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func (l *LoggerClass) Info(msg string, fields ...zap.Field) {
	l.zapLogger.Info(msg, fields...)
}

func (l *LoggerClass) Warn(msg string, fields ...zap.Field) {
	l.zapLogger.Warn(msg, fields...)
}

func (l *LoggerClass) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

func (l *LoggerClass) Error(msg string, fields ...zap.Field) {
	l.zapLogger.Error(msg, fields...)
}

func (l *LoggerClass) Sync() {
	l.zapLogger.Sync()
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Sync() {
	logger.Sync()
}
