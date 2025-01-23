package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    encoderConfig,
		InitialFields: map[string]interface{}{
			"app": "calculation-backend",
			"pid": os.Getpid(),
		},
	}

	logger := zap.Must(config.Build())
	logger.Info("logger construction succeeded")

	return logger
}

var (
	Logger = New()
)
