package infrastructure

import (
	"github.com/ezio1119/fishapp-user/conf"
	"github.com/go-redis/redis/v7"
)

func NewRedisClient() (*redis.Client, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     conf.C.Kvs.Host + ":" + conf.C.Kvs.Port,
		Password: conf.C.Kvs.Pass,
		DB:       conf.C.Kvs.Db,
		Network:  conf.C.Kvs.Net,
	})

	if _, err := c.Ping().Result(); err != nil {
		return nil, err
	}

	return c, nil
}
