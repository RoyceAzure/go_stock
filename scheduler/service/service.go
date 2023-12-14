package service

import (
	"context"

	logger "github.com/RoyceAzure/go-stockinfo-scheduler/repository/logger_distributor"
	jredis "github.com/RoyceAzure/go-stockinfo-scheduler/repository/redis"
	repository "github.com/RoyceAzure/go-stockinfo-scheduler/repository/sqlc"
	worker "github.com/RoyceAzure/go-stockinfo-scheduler/worker"
)

type Service interface {
	SyncDataService
	InitSDA(ctx context.Context)
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
	service := &SchdulerService{
		dao:             dao,
		taskDistributor: taskDistributor,
		redisDao:        redisDao,
	}
	return service
}

/*
下載SDA
*/
func (server *SchdulerService) InitSDA(ctx context.Context) {
	select {
	case <-ctx.Done():
		logger.Logger.Warn().Msg("InitStock get cancel")
	default:
		res, err := server.DownloadAndInsertDataSVAA(ctx)
		if err != nil || res == 0 {
			logger.Logger.Warn().Err(err).Msg("download SDA failed")
		}
	}
}
