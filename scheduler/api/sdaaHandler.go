package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) syncSVAA(ctx *gin.Context) {
	res, err := server.service.DownloadAndInsertDataSVAA(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Insert rows": res})
}
