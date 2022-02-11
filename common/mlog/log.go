package mlog

import (
	"github.com/rs/zerolog"
	"os"
)

var _logger zerolog.Logger

func InitLog() {
	_logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func Info(msg string) {
	_logger.Info().Caller(1).Msg(msg)
}
