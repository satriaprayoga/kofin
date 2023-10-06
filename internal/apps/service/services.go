package service

import (
	"time"

	"github.com/satriaprayoga/kofin/internal/config"
)

type Services struct {
	Account AccountService
}

var services *Services

func SetupServices(cfg config.Config) {
	services = &Services{
		Account: NewAccountService(time.Duration(cfg.Server.ReadTimeOut) * time.Second),
	}
}

func GetServices() *Services {
	return services
}
