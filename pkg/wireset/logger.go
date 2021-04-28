package wireset

import (
	"shortURL/config"
	"shortURL/pkg/logger"

	"github.com/rs/zerolog"
)

func InitLogger(config *config.Config) (zerolog.Logger, error) {
	return logger.NewLogger(config.Logger.Level, config.Logger.Format)
}
