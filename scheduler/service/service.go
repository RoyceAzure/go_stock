package service

import (
	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	worker "github.com/RoyceAzure/go-stockinfo-schduler/worker"
)

type Service interface {
	SyncDataService
}

type SchdulerService struct {
	dao             repository.Dao
	taskDistributor worker.TaskDistributor
}

var _ SyncDataService = (*SchdulerService)(nil)

func NewService(dao repository.Dao, taskDistributor worker.TaskDistributor) *SchdulerService {
	return &SchdulerService{
		dao:             dao,
		taskDistributor: taskDistributor,
	}
}
