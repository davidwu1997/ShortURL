package database

import (
	"shortURL/config"

	"github.com/pkg/errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConnections(config *config.Config) (*Database, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.MySQL.DSN,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		return nil, errors.Wrap(err, "gorm init")
	}

	return &Database{DB: db}, nil
}
