package gapi

import (
	"fmt"

	"github.com/RoyceAzure/go-stockinfo-api/pb"
	"github.com/RoyceAzure/go-stockinfo-api/token"
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	worker "github.com/RoyceAzure/go-stockinfo-worker"
)

type Server struct {
	pb.UnimplementedStockInfoServer
	config          utility.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

const (
	DEFAULT_PAGE      = 1
	DEFAULT_PAGE_SIZE = 10
)

func NewServer(config utility.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create tokenMaker %w", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
