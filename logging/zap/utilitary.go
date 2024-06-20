package zapp

import (
	"log"
	"os"
	"path/filepath"
	"yet-again-templates/logging/zap/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setEncoder() zapcore.Encoder {
	if config.C.EncoderLevel == "production" {
		return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}
	if config.C.EncoderLevel == "development" {
		return zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	}
	log.Fatal("Unknown encoder level: ", config.C.EncoderLevel)
	return nil
}
func getLogFile() *os.File {
	// Trying to create log file
	logfile, err := os.OpenFile(config.C.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		// There is common case that directory doesn't exist
		// So we try to create it
		log.Println("Cannot create log file", err)
		log.Println("Trying to create directory")
		os.Mkdir(filepath.Dir(config.C.LogFile), 0777)

		// Retry to create log file
		logfile, err = os.OpenFile(config.C.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Println("Unsuccessful logger initialization, cannot create log file ", err)
			return nil
		}
	}
	return logfile
}
