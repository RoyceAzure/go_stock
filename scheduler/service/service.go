package service

import (
	jredis "github.com/RoyceAzure/go-stockinfo-scheduler/repository/redis"
	repository "github.com/RoyceAzure/go-stockinfo-scheduler/repository/sqlc"
	worker "github.com/RoyceAzure/go-stockinfo-scheduler/worker"

)

type Service interface {
	SyncDataService
}

type SchdulerService struct {
	dao             repository.Dao
	taskDistributor worker.TaskDistributor
	redisDao        jredis.JRedisDao
}

var _ SyncDataService = (*SchdulerService)(nil)

/*
內部組件支持異步
*/
func NewService(dao repository.Dao, taskDistributor worker.TaskDistributor, redisDao jredis.JRedisDao) *SchdulerService {
	return &SchdulerService{
		dao:             dao,
		taskDistributor: taskDistributor,
		redisDao:        redisDao,
	}
}
