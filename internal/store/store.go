package store

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func SetConnDB(connString string) {
	var (
		err      error
		logLevel logger.LogLevel = logger.Error
		gormCfg  *gorm.Config
	)
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logLevel,    // Log level
			Colorful:      false,       // Disable color
		},
	)
	gormCfg = &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: dbLogger,
	}
	db, err = gorm.Open(postgres.Open(connString), gormCfg)
	if err != nil {
		log.Panicln("failed to set DB Connection")
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Panicln("failed to set DB Pool")

	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	autoMigrate()

}

func CloseDB() {
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDb.Close()
}

func DB() *gorm.DB {
	return db
}
