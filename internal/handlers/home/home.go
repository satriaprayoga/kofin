package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Home struct {
}

func HomeRoute(router *gin.RouterGroup) {
	h := &Home{}
	home := router.Group("/home")
	home.GET("/", h.Index)
}

func (h *Home) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "world"})
}
