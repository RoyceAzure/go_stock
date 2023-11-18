package jredis

import (
	"context"
	"errors"
	"fmt"

	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
	"github.com/redis/go-redis/v9"
)

type Jredis struct {
	client *redis.Client
}

var ErrNotExists = errors.New("data not exists")

func NewJredis(config config.Config) *Jredis {
	redisAddr := config.RedisSchdulerAddress
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return &Jredis{
		client: client,
	}
}

func (jredis *Jredis) Start(ctx context.Context) error {
	err := jredis.client.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect redis-schduler : %w", err)
	}
	return nil
}

func (jredis *Jredis) Close(ctx context.Context) error {
	return jredis.client.Close()
}
