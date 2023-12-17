package gapi

import (
	"context"
	"errors"
	"net"

	"github.com/RoyceAzure/go-stockinfo-distributor/api/token"
	sqlc "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	remote_repo "github.com/RoyceAzure/go-stockinfo-distributor/repository/remote_repo"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/pb"
	"google.golang.org/grpc/peer"
)

/*

1.要能註冊需要哪些股票資料
2.資料要丟給kafuka
3.使用grpc跟schduler 通信  取回資料

*/

type Server struct {
	pb.UnimplementedStockInfoDistributorServer
	dbDao       sqlc.DistributorDao
	schdulerDao remote_repo.SchdulerInfoDao
	tokenMaker  token.Maker
}

func NewServer(dbDao sqlc.DistributorDao, schdulerDao remote_repo.SchdulerInfoDao, tokenMaker token.Maker) *Server {
	server := Server{
		dbDao:       dbDao,
		schdulerDao: schdulerDao,
		tokenMaker:  tokenMaker,
	}
	return &server
}

func (s *Server) GetClientIP(ctx context.Context) (string, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return "", errors.New("cannot find peer info")
	}

	if p.Addr == net.Addr(nil) {
		return "", errors.New("peer address is nil")
	}

	// IP 字串格式為 "ip:port"，這裡我們只關心 IP 部分
	host, _, err := net.SplitHostPort(p.Addr.String())
	if err != nil {
		return "", err
	}

	return host, nil
}
