package main

import (
	"log"
	"os"

	"github.com/satriaprayoga/kofin/internal/apps/repository"
	"github.com/satriaprayoga/kofin/internal/apps/service"
	"github.com/satriaprayoga/kofin/internal/config"
	"github.com/satriaprayoga/kofin/internal/database"
	"github.com/satriaprayoga/kofin/internal/logging"
	"github.com/satriaprayoga/kofin/internal/server"
	"github.com/satriaprayoga/kofin/internal/store"
)

func main() {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to read working directory: %v", err)
	}
	configPath := rootPath + "/configs/config.json"
	conf, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}
	logging.ConfigureLogger(conf.App.LogLevel)
	if conf.App.LogLevel == "prod" {
		logging.SetGinLoginToFile()
	}
	connString := database.NewConfigDB(conf)
	store.SetConnDB(connString)
	defer store.CloseDB()

	repository.SetupRepo()
	service.SetupServices(conf)
	server.Start(conf)
}
