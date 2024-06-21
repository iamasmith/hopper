package config

import (
	"errors"
	"flag"

	"go.uber.org/zap/zapcore"
)

type ConfigT struct {
	flagSet    *flag.FlagSet
	LogLevel   zapcore.Level
	ListenBind string
	Otel       bool
}

var Config = ConfigT{
	flagSet: flag.NewFlagSet("standard", flag.ExitOnError),
}

func (c *ConfigT) setDefaults() {
	c.LogLevel = logLevel
	c.ListenBind = listenBind
	c.Otel = otel
}

// Allows us to call ParseArgs from unit tests over and over
func (c *ConfigT) ResetForTest() {
	c.flagSet = flag.NewFlagSet("test", flag.ContinueOnError)
}

func (c *ConfigT) ParseArgs(args []string) {
	c.setDefaults()
	flag := c.flagSet
	flag.Func(
		"level", "debug|info|warn|error (defaults info)",
		func(s string) error {
			level, has := map[string]zapcore.Level{
				"debug": zapcore.DebugLevel,
				"info":  zapcore.InfoLevel,
				"warn":  zapcore.WarnLevel,
				"error": zapcore.ErrorLevel,
			}[s]
			if !has {
				return errors.New("INVALID LOG LEVEL")
			}
			c.LogLevel = level
			return nil
		},
	)
	flag.StringVar(&c.ListenBind, "listen", ":8000", "Host/Port binding for server")
	flag.BoolVar(&c.Otel, "otel", false, "Enable OTEL exporter configured via environment variables")
	flag.Parse(args)
}
