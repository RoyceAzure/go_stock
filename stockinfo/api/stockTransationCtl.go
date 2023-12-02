package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/RoyceAzure/go-stockinfo/api/token"
	db "github.com/RoyceAzure/go-stockinfo/project/db/sqlc"
	"github.com/gin-gonic/gin"
)

// gin使用go-playground/validator做驗證  這裡的驗證function使用,分隔  不能使用空白建
type createStockTransactionRequest struct {
	UserID          int64  `json:"user_id" binding:"required,min=1"`
	StockID         int64  `json:"stock_id"  binding:"required,min=1"`
	TransactionType string `json:"transaction_type"  binding:"required,oneof=buy sell"`
	TransationAmt   int32  `json:"transation_amt"  binding:"required,gt=0"`
}

func (server *Server) createStockTransaction(ctx *gin.Context) {
	var request createStockTransactionRequest
	//S從 HTTP 請求的 JSON Body 中獲取和驗證資料，並將其填充到 Go 程式中的物件 且只會填充有匹配的
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var stock *db.Stock
	var isValidated bool
	user, isValidated := server.validateUser(ctx, request.UserID)
	if !isValidated {
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if user.UserID != authPayload.UserId {
		err := errors.New("userid doesn't math authorizd user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	stock, isValidated = server.validateStock(ctx, request.StockID)
	if !isValidated {
		return
	}

	arg := db.CreateStockTransactionParams{
		UserID:                  request.UserID,
		StockID:                 request.StockID,
		TransactionType:         request.TransactionType,
		TransactionDate:         time.Now().UTC(),
		TransationAmt:           request.TransationAmt,
		TransationPricePerShare: stock.CurrentPrice,
		CrUser:                  "SYSTEM",
	}

	transfer, err := server.store.CreateStockTransaction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, transfer)
}

// 應該放在service 層
func (server *Server) validateUser(ctx *gin.Context, userID int64) (*db.User, bool) {
	user, err := server.store.GetUser(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return nil, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return nil, false
	}
	return &user, true
}

func (server *Server) validateStock(ctx *gin.Context, stockID int64) (*db.Stock, bool) {
	stock, err := server.store.GetStock(ctx, stockID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return nil, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return nil, false
	}
	return &stock, true
}
