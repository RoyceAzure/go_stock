package api

import (
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/gin-gonic/gin"
)

// 為何gin.Engine要使用*?
// 需要改變Engine內部設置  效率
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	//gin.Default() 也是回傳指標
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.GET("/user/:id", server.getUser)
	router.GET("/users", server.listUser)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// gin.H 是個map[string]any
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
