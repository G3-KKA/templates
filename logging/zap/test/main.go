package main

import (
	zapp "yet-again-templates/logging/zap"
	"yet-again-templates/logging/zap/config"
)

func main() {
	config.InitConfig()
	zapp.InitGlobalLogger()
	zapp.Debug("Debug")
	zapp.Info("Info")
}
