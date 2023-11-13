package service

import (
	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
)

type Service interface {
	SyncDataService
}

type SchdulerService struct {
	dao repository.Dao
}

var _ SyncDataService = (*SchdulerService)(nil)

func NewService(dao repository.Dao) *SchdulerService {
	return &SchdulerService{
		dao: dao,
	}
}
