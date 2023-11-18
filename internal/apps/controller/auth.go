package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
	"github.com/satriaprayoga/kofin/internal/pkg"
)

func AuthRoute(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	auth.POST("/login", AuthLogin)

}

func AuthLogin(c *gin.Context) {
	defer pkg.PanicHandler(c)
	service.GetServices().Auth.Login(c)
}
