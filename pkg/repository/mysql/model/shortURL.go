package model

import (
	"time"
)

type ShortUrl struct {
	ID      int64
	Url     string `gorm:"type:varchar(1024)"`
	Created time.Time
	Updated time.Time
}

func (s *ShortUrl) TableName() string {
	return "short_urls"
}
