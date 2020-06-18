package repository

import (
	"time"

	"github.com/ezio1119/fishapp-auth/usecase/repository"
	"github.com/go-redis/redis/v7"
)

type blackListRepository struct {
	client *redis.Client
}

func NewBlackListRepository(client *redis.Client) repository.BlackListRepository {
	return &blackListRepository{client}
}

func (r *blackListRepository) SetNX(token string, exp time.Duration) (bool, error) {
	return r.client.SetNX(token, "", exp).Result()
}

func (r *blackListRepository) Exists(t string) (int64, error) {
	return r.client.Exists(t).Result()
}
