package app

import (
	"testing"

	"github.com/iamasmith/hopper/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestDefaultSetup(t *testing.T) {
	assert := assert.New(t)
	config.Config.ResetForTest()
	config.Config.ParseArgs([]string{})
	server, app := Setup()
	assert.NotNil(server)
	assert.NotNil(app)
	app.Stop()
}
