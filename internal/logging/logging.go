package logging

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

const (
	logsDir = "logs"
	logName = "kofin_prod.log"
)

var logFilePath = filepath.Join(logsDir, logName)

func SetGinLoginToFile() {
	gin.SetMode(gin.ReleaseMode)
	logFile, err := os.Create(logFilePath)
	if err != nil {
		log.Panic().Err(err).Msg("Error Opening log file")
	}
	gin.DefaultWriter = io.MultiWriter(logFile)
}

func createLogDir() {
	if err := os.Mkdir(logsDir, 0744); err != nil && !os.IsExist(err) {
		log.Fatal().Err(err).Msg("Unable to create logs directory.")
	}
}

func ConfigureLogger(env string) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	switch env {
	case "dev":
		stdOutWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05.000"}
		logger := zerolog.New(stdOutWriter).With().Timestamp().Logger()
		log.Logger = logger
	case "prod":
		createLogDir()
		backupLastLog()
		logFile := openLogFile()
		logFileWriter := zerolog.ConsoleWriter{Out: logFile, NoColor: true, TimeFormat: "15:04:05.000"}
		logger := zerolog.New(logFileWriter).With().Timestamp().Logger()
		log.Logger = logger
	default:
		fmt.Printf("Env not valid: %s\n", env)
		os.Exit(2)
	}
}

func backupLastLog() {
	timeStamp := time.Now().Format("20060201_15_04_05")
	base := strings.TrimSuffix(logName, filepath.Ext(logName))
	bkpLogName := base + "_" + timeStamp + "." + filepath.Ext(logName)
	bkpLogPath := filepath.Join(logsDir, bkpLogName)

	logFile, err := os.ReadFile(logFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Panic().Err(err).Msg("Error reading log file for backup")
	}

	if err = os.WriteFile(bkpLogPath, logFile, 0644); err != nil {
		log.Panic().Err(err).Msg("Error writing backup log file")
	}

}

func openLogFile() *os.File {
	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Panic().Err(err).Msg("Error while opening log file")
	}
	return logFile
}

// func curentDir() string {
// 	path, err := os.Executable()
// 	if err != nil {
// 		log.Panic().Err(err).Msg("Can't get current directory.")
// 	}
// 	return filepath.Dir(path)
// }
