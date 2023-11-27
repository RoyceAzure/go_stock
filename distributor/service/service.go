package service

import (
	"context"
	"sync"

	jkafka "github.com/RoyceAzure/go-stockinfo-distributor/jkafka"
	sqlc "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	remote_repo "github.com/RoyceAzure/go-stockinfo-distributor/repository/remote_repo"
	dto "github.com/RoyceAzure/go-stockinfo-distributor/shared/model/dto"
)

type IDistributorService interface {
	GetFilterSPRByIP(ctx context.Context, ip string) ([]dto.StockPriceRealTimeDTO, error)
	GetPreSprtime(ctx context.Context) string
	SetPreSprtime(ctx context.Context, preSprtime string)
}

type DistributorService struct {
	schdulerDao remote_repo.SchdulerInfoDao
	dbDao       sqlc.DistributorDao
	jkafkaWrite jkafka.KafkaWriter
	preSprtime  string
	sprLock     sync.RWMutex
}

func NewDistributorService(schdulerDao remote_repo.SchdulerInfoDao, dbDao sqlc.DistributorDao, jkafkaWrite jkafka.KafkaWriter) IDistributorService {
	return &DistributorService{
		schdulerDao: schdulerDao,
		dbDao:       dbDao,
		jkafkaWrite: jkafkaWrite,
	}
}
