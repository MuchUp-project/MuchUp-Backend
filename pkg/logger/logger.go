package logger
import (
	"context"
	"os"
	"strings"
	"github.com/rs/zerolog"
)
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	WithContext(ctx context.Context) Logger
	WithError(err error) Logger
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}
type LoggerStruct struct {
	logger *zerolog.Logger
}
var _ Logger = (*LoggerStruct)(nil)
func New(level string) *LoggerStruct {
	var l zerolog.Level
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
	return &LoggerStruct{
		logger: &logger,
	}
}
func NewLogger() Logger {
	return New("info")
}
func (l *LoggerStruct) Debug(msg string, args ...interface{}) {
	l.logger.Debug().Msgf(msg, args...)
}
func (l *LoggerStruct) Info(msg string, args ...interface{}) {
	l.logger.Info().Msgf(msg, args...)
}
func (l *LoggerStruct) Warn(msg string, args ...interface{}) {
	l.logger.Warn().Msgf(msg, args...)
}
func (l *LoggerStruct) Error(msg string, args ...interface{}) {
	l.logger.Error().Msgf(msg, args...)
}
func (l *LoggerStruct) Fatal(msg string, args ...interface{}) {
	l.logger.Fatal().Msgf(msg, args...)
	os.Exit(1)
}
func (l *LoggerStruct) WithContext(ctx context.Context) Logger {
	return &LoggerStruct{
		logger: l.logger,
	}
}
func (l *LoggerStruct) WithError(err error) Logger {
	newLogger := l.logger.With().Err(err).Logger()
	return &LoggerStruct{
		logger: &newLogger,
	}
}
func (l *LoggerStruct) WithField(key string, value interface{}) Logger {
	newLogger := l.logger.With().Interface(key, value).Logger()
	return &LoggerStruct{
		logger: &newLogger,
	}
}
func (l *LoggerStruct) WithFields(fields map[string]interface{}) Logger {
	newLogger := l.logger.With().Fields(fields).Logger()
	return &LoggerStruct{
		logger: &newLogger,
	}
}
func (l *LoggerStruct) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}
func (l *LoggerStruct) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}
func (l *LoggerStruct) Warnf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}
func (l *LoggerStruct) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}
func (l *LoggerStruct) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal().Msgf(format, args...)
	os.Exit(1)
}
