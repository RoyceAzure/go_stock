package redisService

import jredis "github.com/RoyceAzure/go-stockinfo-scheduler/repository/redis"

type RedisService interface {
	RedisServiceSPR
}

type JRedisService struct {
	redisDao jredis.JRedisDao
}

func NewJRedisService(redisDao jredis.JRedisDao) *JRedisService {
	return &JRedisService{
		redisDao: redisDao,
	}
}
