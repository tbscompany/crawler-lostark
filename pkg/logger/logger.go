package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func InitLogger(version string) {
	host, _ := os.Hostname()
	Logger = zerolog.
		New(os.Stderr).
		With().
		Caller().
		Timestamp().
		Str("host", host).
		Str("version", version).
		Logger()
}
