package service

import (
	"context"
	"errors"
	"fmt"
)

type ShortUrlResponse struct {
	Url_ID   string
	ShortUrl string
}

func (s *ShortURL) Upload(ctx context.Context, url string) error {
	if url == "" {
		return errors.New("url is empty")
	}

	id, err := s.DataBase.UploadURL(url)
	if err != nil {
		return err
	}

	encodeString := Encode(uint64(id))
	fmt.Println(encodeString)

	if err := s.redisClient.Set(ctx, encodeString, url); err != nil {
		return err
	}

	return nil
}

func (s *ShortURL) Delete(context context.Context, key string) error {
	if err := s.redisClient.Del(context, key); err != nil {
		return err
	}

	id, err := Decode(key)
	if err != nil {
		return err
	}

	return s.DataBase.DeleteURL(int64(id))
}

func (s *ShortURL) Redirect(context context.Context, url string) (string, error) {
	originalUrl, err := s.redisClient.Get(context, url)
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}
