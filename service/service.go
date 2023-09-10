package service

import (
	"go-shorten-link/repository"
	"go-shorten-link/utils"
)

type ServiceRedis interface {
	SetShortenLink(url string) (string, error)
	GetShortenLink(key string) (string, error)
}

type ServiceRedisImplement struct {
	redisRepository repository.RedisRepository
}

func NewRedisService(redisRepository repository.RedisRepository) ServiceRedis {
	return &ServiceRedisImplement{redisRepository: redisRepository}
}

func (svc ServiceRedisImplement) SetShortenLink(url string) (string, error) {
	generateKey := utils.RandomStringWithNumber()

	_, err := svc.redisRepository.Set(generateKey, url)
	if err != nil {
		return "failed to shorten link", err
	}
	return generateKey, nil

}
func (svc ServiceRedisImplement) GetShortenLink(key string) (string, error) {
	val, err := svc.redisRepository.Get(key)
	if err != nil {
		return "value with this key ' " + key + " ' is not found", err
	}
	return val, nil
}
