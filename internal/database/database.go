package database

import (
	"fmt"

	"github.com/satriaprayoga/kofin/internal/config"
)

func NewConfigDB(cfg config.Config) string {
	var conn string
	conn = fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Name,
		cfg.Database.Port, "disable",
		cfg.Database.TimeZone)
	if cfg.Database.Password != "" {
		conn = fmt.Sprintf("%s password=%s", conn, cfg.Database.Password)
	}

	return conn
}
