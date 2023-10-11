package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
)

func ExpendProgramRoute(router *gin.RouterGroup) {
	expend_program := router.Group("/expend_programs")
	expend_program.POST("/", ExpendProgramCreate)
	expend_program.GET("/:id", ExpendProgramIndex)
	expend_program.DELETE("/:id", ExpendProgramDelete)
	expend_program.PUT("/:id", ExpendProgramUpdate)
	expend_program.GET("/available/:year", ExpendProgramAvailable)
}

func ExpendProgramCreate(c *gin.Context) {
	service.GetServices().ExpendProgram.Create(c)
}

func ExpendProgramDelete(c *gin.Context) {
	service.GetServices().ExpendProgram.Delete(c)
}

func ExpendProgramUpdate(c *gin.Context) {
	service.GetServices().ExpendProgram.Update(c)
}

func ExpendProgramIndex(c *gin.Context) {
	service.GetServices().ExpendProgram.Get(c)
}

func ExpendProgramAvailable(c *gin.Context) {
	service.GetServices().ExpendProgram.GetAvailable(c)
}
