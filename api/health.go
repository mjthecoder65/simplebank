package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) CheckHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (server *Server) CheckReadiness(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "Ok"})
}
