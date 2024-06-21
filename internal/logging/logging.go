package logging

import (
	"github.com/iamasmith/logtozap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Setup(level zapcore.Level) *zap.SugaredLogger {
	z := zap.NewProductionConfig()
	z.Level.SetLevel(level)
	logger := zap.Must(z.Build()).Sugar()
	logtozap.ToSugared(logger, zapcore.WarnLevel)
	return logger
}
