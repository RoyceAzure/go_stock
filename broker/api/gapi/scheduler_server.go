package gapi

import (
	"github.com/RoyceAzure/go-stockinfo-broker/api"
	scheduler_dao "github.com/RoyceAzure/go-stockinfo-broker/repository/remote_dao/schedulerDao"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
)

type SchedulerServer struct {
	pb.UnimplementedStockInfoSchdulerServer
	schedulerDao scheduler_dao.ISchedulerDao
	authorizer   api.IAuthorizer
}

func NewSchedulerServer(schedulerDao scheduler_dao.ISchedulerDao, authorizer api.IAuthorizer) (*SchedulerServer, error) {
	server := &SchedulerServer{
		schedulerDao: schedulerDao,
		authorizer:   authorizer,
	}
	return server, nil
}
