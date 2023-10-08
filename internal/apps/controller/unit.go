package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
)

func UnitRoute(router *gin.RouterGroup) {
	unit := router.Group("/units")
	unit.POST("/", UnitCreate)
	unit.GET("/:id", UnitIndex)
	unit.DELETE("/:id", UnitDelete)
	unit.PUT("/:id", UnitUpdate)
}

func UnitCreate(c *gin.Context) {
	service.GetServices().Unit.Create(c)
}

func UnitDelete(c *gin.Context) {
	service.GetServices().Unit.Delete(c)
}

func UnitUpdate(c *gin.Context) {
	service.GetServices().Unit.Update(c)
}

func UnitIndex(c *gin.Context) {
	service.GetServices().Unit.Get(c)
}
