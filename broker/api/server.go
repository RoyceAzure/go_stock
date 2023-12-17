package api

import (
	"fmt"

	"github.com/RoyceAzure/go-stockinfo-broker/api/token"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     config.Config
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config config.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create tokenMaker %w", err)
	}
	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
	}
	server.SetupRouter()
	return server, nil
}

func (server *Server) SetupRouter() {
	//gin.Default() 也是回傳指標
	router := gin.Default()
	server.router = router
}

func (server *Server) Start(address string) {
	//gin.Default() 也是回傳指標
	server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
