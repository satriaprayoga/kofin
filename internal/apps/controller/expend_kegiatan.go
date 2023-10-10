package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
)

func ExpendKegiatanRoute(router *gin.RouterGroup) {
	expend_kegiatan := router.Group("/expend_kegiatans")
	expend_kegiatan.POST("/", ExpendKegiatanCreate)
	expend_kegiatan.GET("/:id", ExpendKegiatanIndex)
	expend_kegiatan.DELETE("/:id", ExpendKegiatanDelete)
	expend_kegiatan.PUT("/:id", ExpendKegiatanUpdate)
}

func ExpendKegiatanCreate(c *gin.Context) {
	service.GetServices().ExpendKegiatan.Create(c)
}

func ExpendKegiatanDelete(c *gin.Context) {
	service.GetServices().ExpendKegiatan.Delete(c)
}

func ExpendKegiatanUpdate(c *gin.Context) {
	service.GetServices().ExpendKegiatan.Update(c)
}

func ExpendKegiatanIndex(c *gin.Context) {
	service.GetServices().ExpendKegiatan.Get(c)
}
