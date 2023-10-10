package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
)

func ExpendRoute(router *gin.RouterGroup) {
	expend := router.Group("/expends")
	expend.POST("/", ExpendCreate)
	expend.GET("/:id", ExpendIndex)
	expend.DELETE("/:id", ExpendDelete)
	expend.PUT("/:id", ExpendUpdate)
}

func ExpendCreate(c *gin.Context) {
	service.GetServices().Expend.Create(c)
}

func ExpendDelete(c *gin.Context) {
	service.GetServices().Expend.Delete(c)
}

func ExpendUpdate(c *gin.Context) {
	service.GetServices().Expend.Update(c)
}

func ExpendIndex(c *gin.Context) {
	service.GetServices().Expend.Get(c)
}
