package budget

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/repository"
)

type ProgramBudgetService interface {
	//Setup(c *gin.Context)
}

type ProgramBudgetServiceImpl struct {
	r  repository.BudgetRepo
	u  repository.UnitRepo
	e  repository.ExpendRepo
	ep repository.ExpendProgramRepo
	t  time.Duration
}

func NewProgramBudgetService(timeout time.Duration) ProgramBudgetService {
	return &ProgramBudgetServiceImpl{t: timeout}
}

func (pb *ProgramBudgetServiceImpl) Setup(c *gin.Context) {

}
