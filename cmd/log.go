package cmd

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// MustNewLogger creates a new logger from defaults or panics.
func MustNewLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.DisableStacktrace = true
	config.DisableCaller = true
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	log, err := config.Build()
	if err != nil {
		panic(err.Error())
	}
	return log
}
