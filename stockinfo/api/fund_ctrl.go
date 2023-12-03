package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/RoyceAzure/go-stockinfo/api/token"
	db "github.com/RoyceAzure/go-stockinfo/project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/utility/constants"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// gin使用go-playground/validator做驗證  這裡的驗證function使用,分隔  不能使用空白建
type createFundRequest struct {
	Balance      string `json:"balance"  binding:"required,gte=0"`
	CurrencyType string `json:"currency_type"  binding:"required,Currency"`
}

type FundResponseDTO struct {
	FundID       int64  `json:"fund_id"`
	UserID       int64  `json:"user_id"`
	Balance      string `json:"balance"`
	CurrencyType string `json:"currency_type"`
}

func FundToResponseDTO(fund db.Fund) FundResponseDTO {
	return FundResponseDTO{
		FundID:       fund.FundID,
		UserID:       fund.UserID,
		Balance:      fund.Balance,
		CurrencyType: fund.CurrencyType,
	}
}

// 不符合單職責原則  應該要區分不同的Controller
func (server *Server) createFund(ctx *gin.Context) {
	var request createFundRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateFundParams{
		UserID:       authPayload.UserId,
		Balance:      request.Balance,
		CurrencyType: request.CurrencyType,
		CrUser:       "SYSTEM",
	}

	fund, err := server.store.CreateFund(ctx, arg)
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

	ctx.JSON(http.StatusAccepted, fund)
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
	fund, err := server.store.GetFund(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if fund.UserID != authPayload.UserId {
		err := errors.New("fund doesn't belong to the authorized user")
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}
	res := FundToResponseDTO(fund)
	ctx.JSON(http.StatusOK, res)
}

type getFundsRequest struct {
	ID     int64 `uri:"id" binding:"required,min=1"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (server *Server) getFunds(ctx *gin.Context) {
	var request getFundsRequest
	//S從 HTTP 請求的 JSON Body 中獲取和驗證資料，並將其填充到 Go 程式中的物件 且只會填充有匹配的
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//注意  這裡的ctx是由gin.Context提供，這就表示要不要中止process是由gin框架控制

	if request.Limit <= 0 {
		request.Limit = DEFAULT_PAGE - 1
	}
	if request.Offset <= 0 {
		request.Offset = DEFAULT_PAGE_SIZE
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	funds, err := server.store.GetfundByUserId(ctx, db.GetfundByUserIdParams{
		UserID: authPayload.UserId,
		Limit:  request.Limit,
		Offset: request.Offset,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var res []FundResponseDTO
	for i := range funds {
		fund := funds[i]
		res = append(res, FundToResponseDTO(fund))
	}
	ctx.JSON(http.StatusOK, res)
}
