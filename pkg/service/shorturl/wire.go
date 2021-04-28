//+build wireinject

//The build tag makes sure the stub is not built in the final build.
package shorturl

import (
	"shortURL/config"
	"shortURL/internal/delivery/http"
	"shortURL/pkg/repository/mysql/database"
	"shortURL/pkg/repository/redis"
	"shortURL/pkg/service"
	"shortURL/pkg/wireset"

	"github.com/google/wire"
)

func Initialize() (Application, error) {
	wire.Build(
		newApplication,
		config.New,
		wireset.InitLogger,
		http.NewHttpServer,
		redis.InitRedisClient,
		database.InitConnections,
		service.New,
	)
	return Application{}, nil
}
