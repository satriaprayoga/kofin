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
