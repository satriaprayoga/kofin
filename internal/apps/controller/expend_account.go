package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
)

func ExpendAccountRoute(router *gin.RouterGroup) {
	expend_account := router.Group("/expend_accounts")
	expend_account.POST("/", ExpendAccountCreate)
	expend_account.GET("/:id", ExpendAccountIndex)
	expend_account.DELETE("/:id", ExpendAccountDelete)
	expend_account.PUT("/:id", ExpendAccountUpdate)
}

func ExpendAccountCreate(c *gin.Context) {
	service.GetServices().ExpendAccount.Create(c)
}

func ExpendAccountDelete(c *gin.Context) {
	service.GetServices().ExpendAccount.Delete(c)
}

func ExpendAccountUpdate(c *gin.Context) {
	service.GetServices().ExpendAccount.Update(c)
}

func ExpendAccountIndex(c *gin.Context) {
	service.GetServices().ExpendAccount.Get(c)
}
