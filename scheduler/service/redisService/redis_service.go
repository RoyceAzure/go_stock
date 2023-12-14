package redisService

import jredis "github.com/RoyceAzure/go-stockinfo-scheduler/repository/redis"

type RedisService interface {
	RedisServiceSPR
}

type JRedisService struct {
	redisDao jredis.JRedisDao
}

/*
use redisDao to manage spr data in redis
in bs logic level, integrate with cronworker and gapi level
*/
func NewJRedisService(redisDao jredis.JRedisDao) *JRedisService {
	return &JRedisService{
		redisDao: redisDao,
	}
}
