package service

import (
	"shortURL/pkg/repository/mysql/database"
	"shortURL/pkg/repository/redis"
)

type ShortURL struct {
	redisClient *redis.Cache
	DataBase    *database.Database
}

func New(redis *redis.Cache, database *database.Database) *ShortURL {
	s := &ShortURL{
		redisClient: redis,
		DataBase:    database,
	}
	return s
}
