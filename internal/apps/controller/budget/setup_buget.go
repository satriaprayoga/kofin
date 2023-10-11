package budget

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
)

func BudgetSetupRoute(router *gin.RouterGroup) {
	budget := router.Group("/setup")
	budget.POST("/", BudgetSetupCreate)

}

func BudgetSetupCreate(c *gin.Context) {
	service.GetServices().BudgetSetup.Setup(c)
}
