package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func New() zerolog.Logger {
	logger := zerolog.
		New(os.Stderr).
		With().
		Timestamp().
		Logger()

	logger = logger.With().Timestamp().Caller().Logger()

	level := viper.GetString("log-level")
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	return logger.Level(lvl)
}
