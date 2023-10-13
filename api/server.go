package api

import (
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("SSO", validSSO)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("Currency", validCurrency)
	}
	router.POST("/user", server.createUser)
	router.GET("/user/:id", server.getUser)
	router.GET("/users", server.listUser)

	router.POST("/fund", server.createFund)
	router.GET("/fund", server.getFund)
	router.POST("/stockTransfer", server.createStockTransaction)
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
