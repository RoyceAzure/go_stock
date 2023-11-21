package gapi

import (
	"github.com/RoyceAzure/go-stockinfo-schduler/api/pb"
	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/service"
	"github.com/RoyceAzure/go-stockinfo-schduler/service/redisService"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
)

type Server struct {
	pb.UnimplementedStockInfoSchdulerServer
	config       config.Config
	service      service.SyncDataService
	dao          repository.Dao
	redisService redisService.RedisService
}

/*
建立gapi server
task distributir包含在service裡面  有需要獨立出來成server的field??
*/
func NewServer(config config.Config, dao repository.Dao, service service.SyncDataService, redisService redisService.RedisService) (*Server, error) {
	server := &Server{
		config:       config,
		service:      service,
		dao:          dao,
		redisService: redisService,
	}
	return server, nil
}
