package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"github.com/RoyceAzure/go-stockinfo-shared/utility/constants"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// gin使用go-playground/validator做驗證  這裡的驗證function使用,分隔  不能使用空白建
type createUserRequest struct {
	UserName     string `json:"user_name" binding:"required,alphanum"`
	Email        string `json:"email"  binding:"required,email"`
	Password     string `json:"password"  binding:"required,min=6"`
	SsoIdentifer string `json:"sso_identifer"  binding:"required,SSO"`
}
type UserResponseDTO struct {
	UserID       int64     `json:"user_id"`
	UserName     string    `json:"user_name"`
	Email        string    `json:"email"`
	SsoIdentifer string    `json:"sso_identifer"`
	CrDate       time.Time `json:"cr_date"`
	CrUser       string    `json:"cr_user"`
}

func newUserResponse(user db.User) UserResponseDTO {
	return UserResponseDTO{
		UserID:       user.UserID,
		UserName:     user.UserName,
		Email:        user.Email,
		SsoIdentifer: user.SsoIdentifer.String,
		CrDate:       user.CrDate,
		CrUser:       user.CrUser,
	}
}

// 不符合單職責原則  應該要區分不同的Controller
func (server *Server) createUser(ctx *gin.Context) {
	var request createUserRequest
	//S從 HTTP 請求的 JSON Body 中獲取和驗證資料，並將其填充到 Go 程式中的物件 且只會填充有匹配的
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashed_password, err := utility.HashPassword(request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		UserName:       request.UserName,
		Email:          request.Email,
		HashedPassword: hashed_password,
		SsoIdentifer:   utility.StringToSqlNiStr(request.SsoIdentifer),
		CrUser:         "SYSTEM",
	}

	//注意  這裡的ctx是由gin.Context提供，這就表示要不要中止process是由gin框架控制
	user, err := server.store.CreateUser(ctx, arg)
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

	res := newUserResponse(user)

	ctx.JSON(http.StatusAccepted, res)
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

	res := newUserResponse(user)

	ctx.JSON(http.StatusOK, res)
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
	users, err := server.store.Getusers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var responses []UserResponseDTO
	for _, user := range users {
		res := newUserResponse(user)
		responses = append(responses, res)
	}

	ctx.JSON(http.StatusOK, responses)
}

type loginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginUserResponse struct {
	User        UserResponseDTO `json:"user"`
	AccessToken string          `json:"access_token"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	err = utility.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(req.Email, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	rsp := loginUserResponse{
		User:        newUserResponse(user),
		AccessToken: accessToken,
	}
	ctx.JSON(http.StatusOK, rsp)
}
