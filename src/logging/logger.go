package logging

import (
	"github.com/rs/zerolog"

	"crypto-viewer/src/config"
)

func SetLogger(c config.Configs) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	setLogLevel(c)
}

func setLogLevel(configs config.Configs) {
	switch configs.LogLevel {
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
