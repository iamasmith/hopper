package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestSetup(t *testing.T) {
	assert := assert.New(t)
	logger := Setup(zapcore.InfoLevel)
	assert.NotNil(logger, "must create")
}
