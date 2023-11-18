package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	token "github.com/satriaprayoga/kofin/internal/apps/middlewares/jwt"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := token.ValidateToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			ctx.Abort()
			return
		}
		ctx.Next()

	}
}
