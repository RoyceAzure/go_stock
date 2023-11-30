package api

import (
	sqlc "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	remote_repo "github.com/RoyceAzure/go-stockinfo-distributor/repository/remote_repo"
	"github.com/gin-gonic/gin"
)

type Server struct {
	dbDao       sqlc.DistributorDao
	schdulerDao remote_repo.SchdulerInfoDao
	router      *gin.Engine
}

func NewServer(dbDao sqlc.DistributorDao, schdulerDao remote_repo.SchdulerInfoDao) *Server {
	server := Server{
		dbDao:       dbDao,
		schdulerDao: schdulerDao,
	}
	server.setUpRouter()
	return &server
}

func (server *Server) setUpRouter() {
	router := gin.Default()

	router.POST("/client_register", server.CreateClientRegister)
	router.GET("/client_register/:client_uid", server.GetClientRegisterByClientUID)
	router.DELETE("/client_register", server.DeleteClientRegister)
	router.GET("/frontend_client", server.GetFrontendClientByIP)
	router.POST("/frontend_client", server.CreateFrontendClient)
	router.DELETE("/frontend_client/:client_uid", server.DeleteFrontendClient)

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
