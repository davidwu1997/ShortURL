package database

import (
	"errors"
	"shortURL/pkg/repository/mysql/model"
	"time"

	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (database *Database) UploadURL(url string) (int64, error) {
	id, err := database.GetURL(url)
	if err == nil {
		return id, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	newID, err := database.insertURL(url)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (database *Database) GetURL(url string) (int64, error) {
	result := &model.ShortUrl{}
	if err := database.DB.Where("url = ?", url).First(result).Error; err != nil {
		return 0, err
	}

	return result.ID, nil
}

func (database *Database) insertURL(url string) (int64, error) {
	shortUrl := &model.ShortUrl{
		Url:     url,
		Created: time.Now(),
		Updated: time.Now(),
	}

	if err := database.DB.Create(shortUrl).Error; err != nil {
		return 0, err
	}

	return shortUrl.ID, nil
}

func (database *Database) DeleteURL(id int64) error {
	condition := &model.ShortUrl{
		ID: id,
	}

	return database.DB.Delete(condition).Error
}
