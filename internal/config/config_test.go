package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestLogLevelFlag(t *testing.T) {
	assert := assert.New(t)
	Config.ResetForTest()
	Config.ParseArgs([]string{})
	assert.Equal(zapcore.InfoLevel, Config.LogLevel)
	Config.ResetForTest()
	Config.ParseArgs([]string{"-level=debug"})
	assert.Equal(zapcore.DebugLevel, Config.LogLevel)
	Config.ResetForTest()
	Config.ParseArgs([]string{"-level=unknown"})
	assert.Equal(zapcore.InfoLevel, Config.LogLevel)
}

func TestListenFlag(t *testing.T) {
	assert := assert.New(t)
	Config.ResetForTest()
	Config.ParseArgs([]string{})
	assert.Equal(":8000", Config.ListenBind)
	Config.ResetForTest()
	Config.ParseArgs([]string{"-listen=127.0.0.1:8080"})
	assert.Equal("127.0.0.1:8080", Config.ListenBind)
}
