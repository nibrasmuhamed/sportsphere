package logger

import (
	"sync"

	"github.com/mohdjishin/sportsphere/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once   sync.Once
	Logger *LoggerClass
)

// Logger is a wrapper for zap.Logger that provides receiver methods for logging
type LoggerClass struct {
	zapLogger *zap.Logger
}

// Init initializes the logger with a log level based on the LOG_LEVEL environment variable
func Init() {
	logLevel := getLogLevelFromEnv(config.Config.LogLevel)

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(logLevel), // Set the log level dynamically
		Development:      true,
		Encoding:         "json", // or "console" for pretty output
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:  "time",
			LevelKey: "level",
			NameKey:  "logger",
			// CallerKey:      "caller",
			MessageKey: "msg",
			// StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder, // INFO, ERROR, etc.
			EncodeTime:     zapcore.ISO8601TimeEncoder,  // Human-readable time format
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	zapLogger, err := config.Build()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}

	Logger = &LoggerClass{zapLogger: zapLogger}
}

// getLogLevelFromEnv reads the LOG_LEVEL from environment and returns the appropriate zapcore.Level
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
		return zapcore.InfoLevel // Default log level if not set
	}
}

// Info logs a message at InfoLevel
func (l *LoggerClass) Info(msg string, fields ...zap.Field) {
	l.zapLogger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel
func (l *LoggerClass) Warn(msg string, fields ...zap.Field) {
	l.zapLogger.Warn(msg, fields...)
}

// Debug logs a message at DebugLevel
func (l *LoggerClass) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

// Error logs a message at ErrorLevel
func (l *LoggerClass) Error(msg string, fields ...zap.Field) {
	l.zapLogger.Error(msg, fields...)
}

// Sync flushes any buffered log entries
func (l *LoggerClass) Sync() {
	l.zapLogger.Sync()
}
