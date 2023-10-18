package api

import (
	"fmt"

	"github.com/RoyceAzure/go-stockinfo-api/token"
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 為何gin.Engine要使用*?
// 需要改變Engine內部設置  效率
type Server struct {
	config     utility.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

const (
	DEFAULT_PAGE      = 1
	DEFAULT_PAGE_SIZE = 10
)

func NewServer(config utility.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create tokenMaker %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("SSO", validSSO)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("Currency", validCurrency)
	}
	server.SetupRouter()
	return server, nil
}

func (server *Server) SetupRouter() {
	//gin.Default() 也是回傳指標
	router := gin.Default()
	router.GET("/users", server.listUser)
	router.POST("/user/login", server.loginUser)

	//路由前缀: router.Group()的第一個參數是前缀。
	//在此例中，前缀是"/"，這意味著它沒有添加任何特定的前缀到群組內的路由。
	//如果前缀是/api，那麼/user路由就會變成/api/user。
	authRouter := router.Group("/", authMiddleware(server.tokenMaker, &server.store))

	authRouter.POST("/user", server.createUser)
	authRouter.GET("/user/:id", server.getUser)
	authRouter.POST("/fund", server.createFund)
	authRouter.GET("/fund/:id", server.getFund)
	authRouter.POST("/stockTransfer", server.createStockTransaction)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// gin.H 是個map[string]any
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
