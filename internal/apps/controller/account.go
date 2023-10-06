package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
)

func AccountRoute(router *gin.RouterGroup) {
	account := router.Group("/accounts")
	account.POST("/", Create)
}

func Create(c *gin.Context) {
	service.GetServices().Account.Create(c)
}
