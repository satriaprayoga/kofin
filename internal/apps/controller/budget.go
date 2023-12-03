package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/service"
)

func BudgetRoute(router *gin.RouterGroup) {
	budget := router.Group("/budgets")
	budget.GET("/", BudgetList)
	budget.POST("/", BudgetCreate)
	budget.GET("/:id", BudgetIndex)
	budget.DELETE("/:id", BudgetDelete)
	budget.PUT("/:id", BudgetUpdate)

	setup := budget.Group("/setup")
	setup.POST("/", BudgetSetupCreate)
}

func BudgetCreate(c *gin.Context) {
	service.GetServices().Budget.Create(c)
}

func BudgetDelete(c *gin.Context) {
	service.GetServices().Budget.Delete(c)
}

func BudgetUpdate(c *gin.Context) {
	service.GetServices().Budget.Update(c)
}

func BudgetIndex(c *gin.Context) {
	service.GetServices().Budget.Get(c)
}

func BudgetSetupCreate(c *gin.Context) {
	service.GetServices().BudgetSetup.Setup(c)
}

func BudgetList(c *gin.Context) {
	service.GetServices().Budget.FindAll(c)
}
