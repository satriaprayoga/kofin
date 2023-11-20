package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/controller"
	handlers "github.com/satriaprayoga/kofin/internal/handlers/home"
)

func setRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
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
		controller.ExpendObjectRoute(api)
		controller.AuthRoute(api)
	}
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })
	return router
}
