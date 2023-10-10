package service

import (
	"time"

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
	}
}

func GetServices() *Services {
	return services
}
