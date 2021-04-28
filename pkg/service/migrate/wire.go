//+build wireinject

//The build tag makes sure the stub is not built in the final build.

package migrate

import (
	"shortURL/config"
	"shortURL/pkg/repository/mysql/database"
	"shortURL/pkg/wireset"

	"github.com/google/wire"
)

func Initialize(configPath string) (Application, error) {
	wire.Build(
		newApplication,
		database.InitConnections,
		wireset.InitLogger,
		config.New,
	)
	return Application{}, nil
}
