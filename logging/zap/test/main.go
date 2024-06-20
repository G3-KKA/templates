package main

import (
	zapp "yet-again-templates/logging/zap"
	"yet-again-templates/logging/zap/config"
)

// Lower the level of logging - more logs gonna make through
// Debug accumulates everything that -1 and above ( like .Info which is 0 )
func main() {
	config.InitConfig()
	zapp.InitGlobalLogger()
	zapp.Debug("Debug")
	zapp.Info("Info")
	zapp.Info("STDERR")
}
