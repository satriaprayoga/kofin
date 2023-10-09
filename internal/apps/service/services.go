package service

import (
	"time"

	"github.com/satriaprayoga/kofin/internal/config"
)

type Services struct {
	Account  AccountService
	Unit     UnitService
	Program  KProgramService
	Kegiatan KegiatanService
}

var services *Services

func SetupServices(cfg config.Config) {
	services = &Services{
		Account:  NewAccountService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		Unit:     NewUnitService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		Program:  NewKProgramService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
		Kegiatan: NewKegiatanService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
	}
}

func GetServices() *Services {
	return services
}
