package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() (*Server, error) {
	server := &Server{}
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
