package budget

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/satriaprayoga/kofin/internal/apps/repository"
	"github.com/satriaprayoga/kofin/internal/constant"
	dto "github.com/satriaprayoga/kofin/internal/dto/budget"
	"github.com/satriaprayoga/kofin/internal/pkg"
	"github.com/satriaprayoga/kofin/internal/store"
)

type BudgetSetupService interface {
	Setup(c *gin.Context)
	//Verify(c *gin.Context) error
}

type BudgetSetupServiceImpl struct {
	r  repository.BudgetRepo
	u  repository.UnitRepo
	p  repository.KProgramRepo
	e  repository.ExpendRepo
	a  repository.AccountRepo
	ep repository.ExpendProgramRepo
	ea repository.ExpendAccountRepo
	t  time.Duration
}

func NewBudgetSetupService(timeout time.Duration) BudgetSetupService {
	budgetRepo := repository.GetRepo().BudgetRepo
	programRepo := repository.GetRepo().ProgramRepo
	unitRepo := repository.GetRepo().UnitRepo
	expendRepo := repository.GetRepo().ExpendRepo
	accountRepo := repository.GetRepo().AccountRepo
	expendProgramRepo := repository.GetRepo().ExpendProgramRepo
	extendAccountRepo := repository.GetRepo().ExpendAccountRepo
	return &BudgetSetupServiceImpl{r: budgetRepo, u: unitRepo, e: expendRepo, p: programRepo, a: accountRepo, ep: expendProgramRepo, ea: extendAccountRepo, t: timeout}
}

func (a *BudgetSetupServiceImpl) Setup(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(a.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	var setup dto.BudgetSetup
	var budget *store.Budget
	if err := c.ShouldBindJSON(&setup); err != nil {
		log.Err(err).Msg("Error when mapping request for budget creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	today := time.Now()

	if setup.Year < today.Year() {
		log.Err(errors.New("error checking year")).Msg("Error when cheking budget year. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	budget = &store.Budget{
		BudgetYear:   setup.Year,
		BudgetDate:   setup.Date,
		BudgetStatus: 0,
	}

	err := a.r.Create(budget)
	if err != nil {
		log.Err(err).Msg("Error when saving budget setup. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	go a.initializeBudget(setup)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))

}

func (s *BudgetSetupServiceImpl) initializeBudget(b dto.BudgetSetup) {
	var (
		expend   store.Expend
		root     *store.Unit
		err      error
		accounts *[]store.Account
		programs *[]store.KProgram
	)

	root, err = s.u.GetFirstRootUnit()
	if err != nil {
		log.Err(err).Msg("Error when finding root unit on setup budget proccess. Error")
		pkg.PanicException(constant.DataNotFound)
	}

	expend = store.Expend{
		ExpendStatus: 0,
		ExpendYear:   b.Year,
		UnitID:       root.UnitID,
		UnitKode:     root.Kode,
		UnitName:     root.Name,
		Root:         true,
		ExpendPagu:   0.0,
	}

	err = s.e.Create(&expend)
	if err != nil {
		log.Err(err).Msg("Error when creating root expend on setup budget proccess. Error")
		pkg.PanicException(constant.UnknownError)
	}

	accounts, err = s.a.GetRoot("LRA")
	if err != nil {
		log.Err(err).Msg("Error when getting root account LRA on setup budget proccess. Error")
		pkg.PanicException(constant.UnknownError)
	}

	for _, cc := range *accounts {
		err = s.ea.Create(&store.ExpendAccount{
			AccountID:   cc.AccountID,
			AccKode:     cc.AccKode,
			AccName:     cc.AccName,
			Root:        cc.Root,
			Report:      cc.Report,
			AccGroup:    cc.AccGroup,
			AccType:     cc.AccType,
			UnitID:      root.UnitID,
			BudgetYear:  b.Year,
			AccountPagu: 0.0,
		})
		if err != nil {
			log.Err(err).Msg("Error when creating root expend LRA on setup budget proccess. Error")
			pkg.PanicException(constant.UnknownError)
		}
	}

	programs, err = s.p.GetAll()
	if err != nil {
		log.Err(err).Msg("Error when getting programs setup budget proccess. Error")
		pkg.PanicException(constant.UnknownError)
	}

	for _, pp := range *programs {
		err = s.ep.Create(&store.ExpendProgram{
			ProgramID:   pp.ProgramID,
			ProgramKode: pp.ProgramKode,
			ProgramName: pp.ProgramName,
			ProgramPagu: 0.0,
			BudgetYear:  b.Year,
			UnitID:      pp.UnitID,
			UnitKode:    pp.UnitKode,
			UnitName:    pp.UnitName,
			Included:    false,
		})
		if err != nil {
			log.Err(err).Msg("Error when creating expend programs on setup budget proccess. Error")
			pkg.PanicException(constant.UnknownError)
		}
	}

}
