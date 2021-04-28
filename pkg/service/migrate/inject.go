package migrate

import (
	"fmt"
	"shortURL/pkg/repository/mysql/database"

	"shortURL/deployment/migration"

	"shortURL/config"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"github.com/rs/zerolog"
)

type Application struct {
	config *config.Config
	logger zerolog.Logger
	mysql  *database.Database
}

func (application Application) Start() error {
	m := gormigrate.New(application.mysql.DB, gormigrate.DefaultOptions, migration.Migrations)
	if err := m.Migrate(); err != nil {
		return err
	}
	fmt.Println("migration complete")
	return nil
}

func newApplication(config *config.Config, db *database.Database, logger zerolog.Logger) Application {
	return Application{
		config: config,
		logger: logger,
		mysql:  db,
	}
}
