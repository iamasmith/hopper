package config

import "go.uber.org/zap/zapcore"

const (
	listenBind = ":8000"
	logLevel   = zapcore.InfoLevel
	otel       = false
)
