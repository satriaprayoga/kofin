package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
)

func KegiatanRoute(router *gin.RouterGroup) {
	kegiatan := router.Group("/kegiatans")
	kegiatan.POST("/", KegiatanCreate)
	kegiatan.GET("/:id", KegiatanIndex)
	kegiatan.DELETE("/:id", KegiatanDelete)
	kegiatan.PUT("/:id", KegiatanUpdate)
	kegiatan.POST("/paginate/:id", KegiatanPaginate)
	//kegiatan.GET("/program/:id", KegiatanProgram)
}

func KegiatanCreate(c *gin.Context) {
	service.GetServices().Kegiatan.Create(c)
}

func KegiatanDelete(c *gin.Context) {
	service.GetServices().Kegiatan.Delete(c)
}

func KegiatanUpdate(c *gin.Context) {
	service.GetServices().Kegiatan.Update(c)
}

func KegiatanIndex(c *gin.Context) {
	service.GetServices().Kegiatan.Get(c)
}

func KegiatanProgram(c *gin.Context) {
	service.GetServices().Kegiatan.GetByProgramID(c)
}

func KegiatanPaginate(c *gin.Context) {
	service.GetServices().Kegiatan.PaginateSearch(c)
}
