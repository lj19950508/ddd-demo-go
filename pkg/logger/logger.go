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

	skipFrameCount := 3
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
	//执行玩需要关闭系统 但是zerolog已经实现
	// os.Exit(1)
}

// func (l *Logger) log(message string, args ...interface{}) {
// 	if len(args) == 0 {
// 		l.logger.Info().Msg(message)
// 	} else {
// 		l.logger.Info().Msgf(message, args...)
// 	}
// }

// func (l *Logger) msg(level string, message interface{}, args ...interface{}) {
// 	switch msg := message.(type) {
// 	case error:
// 		l.log(msg.Error(), args...)
// 	case string:
// 		l.log(msg, args...)
// 	default:
// 		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
// 	}
// }
