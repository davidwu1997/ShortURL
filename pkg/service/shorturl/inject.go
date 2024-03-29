package shorturl

import (
	"fmt"
	"shortURL/config"
	"shortURL/internal/delivery/http"

	"github.com/rs/zerolog"
)

type Application struct {
	logger     zerolog.Logger
	config     *config.Config
	httpServer *http.HttpServer
}

func (application Application) Start() error {
	application.logger.Info().Msgf("http server listens :%d", application.config.HTTP.Port)
	return application.httpServer.Run(fmt.Sprintf(":%d", application.config.HTTP.Port))
}

func newApplication(
	logger zerolog.Logger,
	config *config.Config,
	httpServer *http.HttpServer,
) Application {
	return Application{
		logger:     logger,
		config:     config,
		httpServer: httpServer,
	}
}
