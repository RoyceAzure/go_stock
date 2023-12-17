package gapi

import (
	"fmt"

	"github.com/RoyceAzure/go-stockinfo/api/token"
	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	"github.com/RoyceAzure/go-stockinfo/shared/util/config"
	"github.com/RoyceAzure/go-stockinfo/worker"

)

type Server struct {
	pb.UnimplementedStockInfoServer
	config          config.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
	clientFactory   SchedulerClientFactory
}

const (
	DEFAULT_PAGE      = 1
	DEFAULT_PAGE_SIZE = 10
)

func NewServer(config config.Config, store db.Store, taskDistributor worker.TaskDistributor, clientFactory SchedulerClientFactory) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create tokenMaker %w", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
		clientFactory:   clientFactory,
	}

	return server, nil
}
