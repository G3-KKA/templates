package zapp

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	PRODUCTION_FROM_CONFIG  = true
	DEVELOPMENT_FROM_CONFIG = false
)
const PATH_TO_LOG_FILE_FROM_CONFIG = "./tmp/logs"

func InitLogger() (*zap.SugaredLogger, error) {
	log.Println("Logger initialization started")
	encoder := setEncoder()

	logfile, err := os.OpenFile(PATH_TO_LOG_FILE_FROM_CONFIG, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("Cannot create log file", err)
		log.Println("Trying to create directory")
		os.Mkdir(PATH_TO_LOG_FILE_FROM_CONFIG[:strings.LastIndex(PATH_TO_LOG_FILE_FROM_CONFIG, "/")], 0777)
		logfile, err = os.OpenFile(PATH_TO_LOG_FILE_FROM_CONFIG, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Println("Unsuccessful logger initialization, cannot create log file ", err)
			return nil, err
		}
	}
	//Creating zap.Cores
	//And merging them
	logfilecore := zapcore.NewCore(encoder, logfile, zapcore.InfoLevel)
	stdout := zapcore.NewCore(encoder, os.Stdout, zapcore.InfoLevel)
	stderr := zapcore.NewCore(encoder, os.Stderr, zapcore.DebugLevel)
	core := zapcore.NewTee(stdout, logfilecore, stderr)
	//Creating Logger
	//And sugaring
	logger := zap.New(core)
	sugarlogger := logger.Sugar()
	//First log message
	//That tells us that logger construction succeeded
	defer sugarlogger.Sync()
	sugarlogger.Debug("Logger construction succeeded")

	return sugarlogger, nil
}
func setEncoder() zapcore.Encoder {
	var encoder zapcore.Encoder
	if PRODUCTION_FROM_CONFIG {
		encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	} else if DEVELOPMENT_FROM_CONFIG {
		encoder = zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	}
	return encoder
}
func __DEPRECATED__InitLogger() *zap.SugaredLogger {
	// For some users, the presets offered by the NewProduction, NewDevelopment,
	// and NewExample constructors won't be appropriate. For most of those
	// users, the bundled Config struct offers the right balance of flexibility
	// and convenience. (For more complex needs, see the AdvancedConfiguration
	// example.)
	//
	// See the documentation for Config and zapcore.EncoderConfig for all the
	// available options.
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "./tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "levelEncoder": "lowercase"
	  }
	}`)
	/*
		"initialFields": {"foo": "bar"},
		just key:value that will be added to every log record
	*/

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	logger, err := cfg.Build()
	sugerlogger := logger.Sugar()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugerlogger.Info("logger construction succeeded")
	return sugerlogger
}
