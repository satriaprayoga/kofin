package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/controller"
	handlers "github.com/satriaprayoga/kofin/internal/handlers/home"
)

func setRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		handlers.HomeRoute(api)
		controller.AccountRoute(api)
		controller.UnitRoute(api)
		controller.ProgramRoute(api)
		controller.KegiatanRoute(api)
		controller.BudgetRoute(api)
		controller.ExpendRoute(api)
		controller.ExpendProgramRoute(api)
		controller.ExpendKegiatanRoute(api)
		controller.ExpendAccountRoute(api)
	}
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })
	return router
}
