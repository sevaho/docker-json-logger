package main

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.LevelFieldName = "levelname"
	zerolog.TimestampFieldName = "asctime"

	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return strings.ToUpper(l.String())
	}
	Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
}
