package main

import (
	"log"
	"time"
	"yet-again-templates/logging/zap/internal/config"
	zapp "yet-again-templates/logging/zap/internal/logger"
)

func main() {
	config.InitConfig()
	err := zapp.InitGlobalLogger(config.Get())
	if err != nil {
		log.Fatal(err)
	}
	zapp.Info("Hello, World!")
	zapp.Debug("Hello, World!")
	time.Sleep(time.Second)

}
