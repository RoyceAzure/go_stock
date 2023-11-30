package api

import (
	"fmt"
	"net/http"
	"time"

	repository "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateClientRegisterRequestDTO struct {
	ClientUid *uuid.UUID `json:"client_uid,omitempty"`
	StockCode string     `json:"stock_code" binding:"required`
}

type ClientRegisterResponseDTO struct {
	ClientUid uuid.UUID  `json:"client_uid,omitempty"`
	StockCode string     `json:"stock_code"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (server *Server) CreateClientRegister(ctx *gin.Context) {
	var req CreateClientRegisterRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var fc repository.FrontendClient
	var err error
	if req.ClientUid == nil {
		clientIp := ctx.ClientIP()
		fc, err = server.dbDao.GetFrontendClientByIP(ctx, clientIp)
	} else {
		fc, err = server.dbDao.GetFrontendClientByID(ctx, *req.ClientUid)
	}
	if err != nil {
		err = fmt.Errorf("can't find client with ip, err : %w", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//要不要檢查stockCode??

	crParm := repository.CreateClientRegisterParams{
		ClientUid: fc.ClientUid,
		StockCode: req.StockCode,
	}

	entity, err := server.dbDao.CreateClientRegister(ctx, crParm)
	if err != nil {
		errCode := repository.ErrorCode(err)
		if errCode == repository.UniqueViolation || errCode == repository.ForeginKeyViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := cvCrEntity2CrResDTO(entity)
	ctx.JSON(http.StatusAccepted, res)
}

type DeleteClientRegisterRequestDTO struct {
	ClientUid uuid.UUID `json:"client_uid" binding:"required"`
	StockCode string    `json:"stock_code" binding:"required"`
}

func (server *Server) DeleteClientRegister(ctx *gin.Context) {
	var req DeleteClientRegisterRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := repository.DeleteClientRegisterParams{
		ClientUid: req.ClientUid,
		StockCode: req.StockCode,
	}

	err := server.dbDao.DeleteClientRegister(ctx, arg)
	if err != nil {
		err = fmt.Errorf("delete client register get err : %w", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"result": fmt.Sprintf("success deleted client with id : %s", req.ClientUid)})
}

type GetClientRegisterByClientUIDRequest struct {
	ClientUID string `uri:"client_uid" binding:"required"`
}

type GetClientRegisterByClientUIDResponse struct {
	Result []ClientRegisterResponseDTO `json:"result"`
}

func (server *Server) GetClientRegisterByClientUID(ctx *gin.Context) {
	var req GetClientRegisterByClientUIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var clientId uuid.UUID
	var err error
	if clientId, err = uuid.Parse(req.ClientUID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	entity, err := server.dbDao.GetClientRegisterByClientUID(ctx, clientId)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := GetClientRegisterByClientUIDResponse{}

	for _, item := range entity {
		res.Result = append(res.Result, cvCrEntity2CrResDTO(item))
	}

	ctx.JSON(http.StatusAccepted, res)
}

func cvCrEntity2CrResDTO(value repository.ClientRegister) ClientRegisterResponseDTO {
	var updated_at *time.Time
	if value.UpdatedAt.Valid {
		updated_at = &value.UpdatedAt.Time
	} else {
		updated_at = nil
	}

	return ClientRegisterResponseDTO{
		ClientUid: value.ClientUid,
		StockCode: value.StockCode,
		CreatedAt: value.CreatedAt,
		UpdatedAt: updated_at,
	}
}
