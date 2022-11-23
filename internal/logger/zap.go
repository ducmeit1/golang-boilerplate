package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a new zap.Logger instance
func NewLogger() *zap.Logger {
	var (
		level       = zapcore.InfoLevel
		development = false
	)

	// Determines if the zapcore.DebugLevel should be used.
	if os.Getenv("DEBUG") != "" {
		level = zapcore.DebugLevel
		development = true
	}

	zapLogger := zap.Must(zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: development,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		// Configures the zapcore encoder config
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "name",
			MessageKey:    "message",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
			EncodeName:    zapcore.FullNameEncoder,
		},
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stdout"},
		DisableStacktrace: true,
	}.Build())

	// Replace globals and redirect stdLog.
	zap.ReplaceGlobals(zapLogger)
	zap.RedirectStdLog(zapLogger)

	return zapLogger
}
