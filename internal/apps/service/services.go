package service

import (
	"time"

	"github.com/satriaprayoga/kofin/internal/apps/service/budget"
	"github.com/satriaprayoga/kofin/internal/config"
)

type Services struct {
	Account        AccountService
	Unit           UnitService
	Program        KProgramService
	Kegiatan       KegiatanService
	Budget         BudgetService
	Expend         ExpendService
	ExpendProgram  ExpendProgramService
	ExpendKegiatan ExpendKegiatanService
	ExpendAccount  ExpendAccountService

	BudgetSetup    budget.BudgetSetupService
	ProgramBudget  budget.ProgramBudgetService
	KegiatanBudget budget.KegiatanBudgetService
	ExpendObject   ExpendObjectService
	ObjectBudget   budget.ObjectBudgetService
}

var services *Services

func SetupServices(cfg config.Config) {
	services = &Services{
		Account:        NewAccountService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		Unit:           NewUnitService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		Program:        NewKProgramService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		Kegiatan:       NewKegiatanService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		Budget:         NewBudgetService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		Expend:         NewExpendService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		ExpendProgram:  NewExpendProgramService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		ExpendKegiatan: NewExpendKegiatanService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		ExpendAccount:  NewExpendAccountService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		BudgetSetup:    budget.NewBudgetSetupService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		ProgramBudget:  budget.NewProgramBudgetService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		KegiatanBudget: budget.NewKegiatanBudgetService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		ObjectBudget:   budget.NewObjectBudgetService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
	}
}

func GetServices() *Services {
	return services
}
