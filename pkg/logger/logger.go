package logger

import (
	"os"
	"strings"
	"github.com/rs/zerolog"
)

// 统一日志接口，不管什么日志框架都符合这个规范
type Interface interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}

// Logger -.
type Logger struct {
	logger *zerolog.Logger
}

// var Instance Interface = (*Logger)(nil)

// New -.
//cfg如何抽象出来
func New(level string) Interface {
	var l zerolog.Level

	//日志级别控制
	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(l)

	//在默认跳过2帧stack的情况下 基于logger本身框架再跳一层
	skipFrameCount := 1
	logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()
	return &Logger{
		logger: &logger,
	}
}

// Debug -.
func (l *Logger) Debug(message string, args ...interface{}) {
	l.logger.Debug().Msgf(message, args...)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {
	l.logger.Info().Msgf(message, args...)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.logger.Warn().Msgf(message, args)
}

// Error -.
func (l *Logger) Error(message string, args ...interface{}) {
	l.logger.Error().Msgf(message, args)
}

// Fatal -.
func (l *Logger) Fatal(message string, args ...interface{}) {
	l.logger.Fatal().Msgf(message, args...)
}

