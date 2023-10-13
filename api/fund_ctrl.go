package api

import (
	"database/sql"
	"net/http"

	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility/constants"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// gin使用go-playground/validator做驗證  這裡的驗證function使用,分隔  不能使用空白建
type createFundRequest struct {
	UserID       int64  `json:"user_id"  binding:"required"`
	Balance      string `json:"balance"  binding:"required,gte=0"`
	CurrencyType string `json:"currency_type"  binding:"required,Currency"`
}

// 不符合單職責原則  應該要區分不同的Controller
func (server *Server) createFund(ctx *gin.Context) {
	var request createFundRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFundParams{
		UserID:       request.UserID,
		Balance:      request.Balance,
		CurrencyType: request.CurrencyType,
		CrUser:       "SYSTEM",
	}

	user, err := server.store.CreateFund(ctx, arg)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			switch pgErr.Code.Name() {
			case constants.ForeignKeyViolation, constants.UniqueViolation:
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, user)
}

type getFundRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getFund(ctx *gin.Context) {
	var request getFundRequest
	//S從 HTTP 請求的 JSON Body 中獲取和驗證資料，並將其填充到 Go 程式中的物件 且只會填充有匹配的
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//注意  這裡的ctx是由gin.Context提供，這就表示要不要中止process是由gin框架控制
	user, err := server.store.GetFund(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
