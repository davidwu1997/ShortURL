package migration

import (
	"shortURL/pkg/repository/mysql/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var v20210423 = &gormigrate.Migration{
	ID: "20210423",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&model.ShortUrl{}); err != nil {
			return err
		}
		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable(&model.ShortUrl{}); err != nil {
			return err
		}
		return nil
	},
}
