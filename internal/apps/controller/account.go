package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
)

func AccountRoute(router *gin.RouterGroup) {
	account := router.Group("/accounts")
	account.POST("/", AccountCreate)
	account.GET("/:id", AccountIndex)
	account.DELETE("/:id", AccountDelete)
	account.PUT("/:id", AccountUpdate)
}

func AccountCreate(c *gin.Context) {
	service.GetServices().Account.Create(c)
}

func AccountDelete(c *gin.Context) {
	service.GetServices().Account.Delete(c)
}

func AccountUpdate(c *gin.Context) {
	service.GetServices().Account.Update(c)
}

func AccountIndex(c *gin.Context) {
	service.GetServices().Account.Get(c)
}
