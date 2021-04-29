package redis

import (
	"context"
	"shortURL/config"
)

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	originalUrl, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}

func (c *Cache) Set(ctx context.Context, key, value string) error {
	expireTime := config.Get().ShortURL.ExpireTime
	if _, err := c.Client.Set(ctx, key, value, expireTime).Result(); err != nil {
		return err
	}

	return nil
}

func (c *Cache) Del(ctx context.Context, key string) error {
	if _, err := c.Client.Del(ctx, key).Result(); err != nil {
		return err
	}

	return nil
}
