package api

import (
	"fmt"
	"net/http"
	"time"

	repository "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateFrontendClientRequestDTO struct {
	Region string `json:"region" binding:"required"`
}

type FrontendClientResponseDTO struct {
	ClientUid uuid.UUID `json:"client_uid"`
	Ip        string    `json:"ip"`
	Region    string    `json:"region"`
	CreatedAt time.Time `json:"created_at"`
}

func (server *Server) CreateFrontendClient(ctx *gin.Context) {
	var req CreateFrontendClientRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ip := ctx.ClientIP()
	if ip == "" {
		err := fmt.Errorf("client ip is empty")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	} else if !util.IsValidIP(ip) {
		err := fmt.Errorf("invalid ip format")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := repository.CreateFrontendClientParams{
		Ip:     ip,
		Region: req.Region,
	}

	entity, err := server.dbDao.CreateFrontendClient(ctx, arg)
	if err != nil {
		errCode := repository.ErrorCode(err)
		if errCode == repository.UniqueViolation || errCode == repository.ForeginKeyViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := FrontendClientResponseDTO{
		ClientUid: entity.ClientUid,
		Ip:        entity.Ip,
		Region:    entity.Region,
		CreatedAt: entity.CreatedAt,
	}

	ctx.JSON(http.StatusAccepted, res)
}

type deleteFrontendClientRequest struct {
	ClientUID string `uri:"client_uid" binding:"required"`
}

func (server *Server) DeleteFrontendClient(ctx *gin.Context) {
	var req deleteFrontendClientRequest
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
	err = server.dbDao.DeleteFrontendClient(ctx, clientId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}

func (server *Server) GetFrontendClientByIP(ctx *gin.Context) {
	ip := ctx.ClientIP()
	if ip == "" {
		err := fmt.Errorf("ip is empty")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	} else {
		if !util.IsValidIP(ip) {
			err := fmt.Errorf("invalid ip format")
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}
	fc, err := server.dbDao.GetFrontendClientByIP(ctx, ip)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			err = fmt.Errorf("frontend client not found, %w", err)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusAccepted, fc)
}
