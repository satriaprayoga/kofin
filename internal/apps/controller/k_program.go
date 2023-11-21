package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
)

func ProgramRoute(router *gin.RouterGroup) {
	program := router.Group("/programs")
	program.POST("/", ProgramCreate)
	program.GET("/:id", ProgramIndex)
	program.DELETE("/:id", ProgramDelete)
	program.PUT("/:id", ProgramUpdate)
	program.GET("/search", ProgramSearch)
	program.POST("/paginate", ProgramPaginate)
	program.GET("/kegiatan/:id", ProgramKegiatan)
}

func ProgramCreate(c *gin.Context) {
	service.GetServices().Program.Create(c)
}

func ProgramDelete(c *gin.Context) {
	service.GetServices().Program.Delete(c)
}

func ProgramUpdate(c *gin.Context) {
	service.GetServices().Program.Update(c)
}

func ProgramIndex(c *gin.Context) {
	service.GetServices().Program.Get(c)
}

func ProgramSearch(c *gin.Context) {
	service.GetServices().Program.TextSearch(c)
}

func ProgramPaginate(c *gin.Context) {
	service.GetServices().Program.PaginateSearch(c)
}

func ProgramKegiatan(c *gin.Context) {
	service.GetServices().Kegiatan.GetByProgramID(c)
}
