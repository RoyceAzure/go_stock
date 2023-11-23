package service

import (
	"context"

	sqlc "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	remote_repo "github.com/RoyceAzure/go-stockinfo-distributor/repository/remote_repo"
	dto "github.com/RoyceAzure/go-stockinfo-distributor/shared/model/dto"
)

type IDistributorService interface {
	GetFilterSPRByIP(ctx context.Context, ip string) ([]dto.StockPriceRealTimeDTO, error)
}

type DistributorService struct {
	schdulerDao remote_repo.SchdulerInfoDao
	dbDao       sqlc.DistributorDao
}

func NewDistributorService(schdulerDao remote_repo.SchdulerInfoDao, dbDao sqlc.DistributorDao) IDistributorService {
	return &DistributorService{
		schdulerDao: schdulerDao,
		dbDao:       dbDao,
	}
}
