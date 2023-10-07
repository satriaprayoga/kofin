package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
)

func AccountRoute(router *gin.RouterGroup) {
	account := router.Group("/accounts")
	account.POST("/", Create)
	account.GET("/:id", Index)
	account.DELETE("/:id", Delete)
	account.PUT("/:id", Update)
}

func Create(c *gin.Context) {
	service.GetServices().Account.Create(c)
}

func Delete(c *gin.Context) {
	service.GetServices().Account.Delete(c)
}

func Update(c *gin.Context) {
	service.GetServices().Account.Update(c)
}

func Index(c *gin.Context) {
	service.GetServices().Account.Get(c)
}
