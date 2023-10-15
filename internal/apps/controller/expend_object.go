package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
)

func ExpendObjectRoute(router *gin.RouterGroup) {
	expend_object := router.Group("/expend_objects")
	expend_object.POST("/", ExpendObjectCreate)
	expend_object.GET("/:id", ExpendObjectIndex)
	expend_object.DELETE("/:id", ExpendObjectDelete)
	expend_object.PUT("/:id", ExpendObjectUpdate)
	expend_object.POST("/rincian", ExpendObjectImport)
	expend_object.PUT("/ubah/:id", ExpendObjectUpdateRincian)
}

func ExpendObjectCreate(c *gin.Context) {
	service.GetServices().ExpendObject.Create(c)
}

func ExpendObjectDelete(c *gin.Context) {
	service.GetServices().ExpendObject.Delete(c)
}

func ExpendObjectUpdate(c *gin.Context) {
	service.GetServices().ExpendObject.Update(c)
}

func ExpendObjectIndex(c *gin.Context) {
	service.GetServices().ExpendObject.Get(c)
}

func ExpendObjectImport(c *gin.Context) {
	service.GetServices().ObjectBudget.AddToAccount(c)
}
func ExpendObjectUpdateRincian(c *gin.Context) {
	service.GetServices().ObjectBudget.UpdateOnAccount(c)
}
