package api

import (
	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/service"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config  config.Config
	service service.SyncDataService
	dao     repository.Dao
	router  *gin.Engine
}

func NewServer(config config.Config, dao repository.Dao, service service.SyncDataService) (*Server, error) {
	server := &Server{
		config:  config,
		service: service,
		dao:     dao,
	}
	server.SetupRouter()
	return server, nil
}

func (server *Server) SetupRouter() {
	//gin.Default() 也是回傳指標
	router := gin.Default()

	synRouter := router.Group("/v1/syncData")

	synRouter.GET("/stock_day_avg_all", server.syncSVAA)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
