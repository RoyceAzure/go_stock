package api

import (
	"database/sql"
	"net/http"

	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"github.com/gin-gonic/gin"
)

// gin使用go-playground/validator做驗證  這裡的驗證function使用,分隔  不能使用空白建
type createUserRequest struct {
	UserID          int64  `json:"user_id" binding:"required,min=1"`
	StockID         int64  `json:"stock_id"  binding:"required,min=1"`
	TransactionType string `json:"transaction_type"  binding:"required,oneof=buy sell"`
	TransationAmt   int64  `json:"transation_amt"  binding:"required,gt=0"`
	TransationProcePerShare
}

// 不符合單職責原則  應該要區分不同的Controller
func (server *Server) createUser(ctx *gin.Context) {
	var request createUserRequest
	//S從 HTTP 請求的 JSON Body 中獲取和驗證資料，並將其填充到 Go 程式中的物件 且只會填充有匹配的
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		UserName:     request.UserName,
		Email:        request.Email,
		Password:     utility.StringToSqlNiStr(request.Password),
		SsoIdentifer: utility.StringToSqlNiStr(request.SsoIdentifer),
		CrUser:       "SYSTEM",
	}

	//注意  這裡的ctx是由gin.Context提供，這就表示要不要中止process是由gin框架控制
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, user)
}
