package logger

import (
	"os"

	"github.com/rs/zerolog"

	"crypto-viewer/src/config"
)

var Log = zerolog.New(os.Stdout).With().Timestamp().Logger()

func SetLogLevel(c config.Configs) {
	switch c.LogLevel {
	case "DEBUG":
		Log.Level(zerolog.DebugLevel)
	case "INFO":
		Log = Log.Level(zerolog.InfoLevel)
	case "TRACE":
		Log = Log.Level(zerolog.TraceLevel)
	case "PANIC":
		Log = Log.Level(zerolog.PanicLevel)
	case "NOLEVEL":
		Log = Log.Level(zerolog.NoLevel)
	case "ERROR":
		Log = Log.Level(zerolog.ErrorLevel)
	default:
		Log.Level(zerolog.InfoLevel)
	}

	Log.Info().Any("log level", Log.GetLevel().String()).Msg("Log level setted")
}
