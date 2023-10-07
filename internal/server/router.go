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
	}
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })
	return router
}
