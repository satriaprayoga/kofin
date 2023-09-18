package server

import (
	"fmt"

	"github.com/satriaprayoga/kofin/internal/config"
)

func Start(cfg config.Config) {
	router := setRouter()
	addrs := fmt.Sprintf(":%d", cfg.Server.HTTPPort)
	router.Run(addrs)
}
