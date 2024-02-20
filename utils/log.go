package utils

import (
	"os"

	"github.com/rs/zerolog"
)

func log() zerolog.Logger{
	return zerolog.New(os.Stdout).With().Timestamp().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})

}