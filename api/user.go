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
	UserName     string `json:"user_name" binding:"required"`
	Email        string `json:"email"  binding:"required,email"`
	Password     string `json:"password"  binding:"required"`
	SsoIdentifer string `json:"sso_identifer"  binding:"required,oneof=MS GOOGLE AWS FB"`
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

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var request getUserRequest
	//S從 HTTP 請求的 JSON Body 中獲取和驗證資料，並將其填充到 Go 程式中的物件 且只會填充有匹配的
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//注意  這裡的ctx是由gin.Context提供，這就表示要不要中止process是由gin框架控制
	user, err := server.store.GetUser(ctx, request.ID)
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

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listUser(ctx *gin.Context) {
	var request listUserRequest
	//S從 HTTP 請求的 JSON Body 中獲取和驗證資料，並將其填充到 Go 程式中的物件 且只會填充有匹配的
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetusersParams{
		Limit:  request.PageSize,
		Offset: (request.PageID - 1) * request.PageSize,
	}

	//注意  這裡的ctx是由gin.Context提供，這就表示要不要中止process是由gin框架控制
	user, err := server.store.Getusers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
