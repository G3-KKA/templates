package zappconfigless

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// This is 0 dependency , hardcoded version
// Should be used in cases:
// - Configuration is incopatible with config-dependent version
// - Need fast drop-in zap logger for small apps

// Should be initialised via InitGlobalLogger()
var globalLogger *zap.SugaredLogger

// Little wrapper for future ease of identification
type LevelWithName struct {
	zap.AtomicLevel
	Name string
}

func withName(name string, level zap.AtomicLevel) LevelWithName {
	return LevelWithName{level, name}
}

// []LevelWithName may be used to change specific output destination log levels
// Thread safe
func InitLogger() (*zap.SugaredLogger, []LevelWithName) {
	log.Println("Logger initialization started")

	// May dynamicly change log levels in runtime, will be returned from InitLogger()

	// Creating cores NON dynamic from Hardcoded constants
	// stderr/stdout supported, network not supported
	// TODO: Add network support

	// Hardcode
	cores := []zapcore.Core{
		zapcore.NewCore(mustSetEncoder("production"), getLogDest("stderr"), zap.NewAtomicLevelAt(zapcore.InfoLevel)),
		zapcore.NewCore(mustSetEncoder("production"), getLogDest("./logs"), zap.NewAtomicLevelAt(zapcore.DebugLevel)),
	}
	levels := []LevelWithName{
		withName("stderr", zap.NewAtomicLevelAt(zapcore.InfoLevel)),
		withName("logfile", zap.NewAtomicLevelAt(zapcore.DebugLevel)),
	}
	// Creating zap.Cores
	// And merging them
	core := zapcore.NewTee(cores...)

	// Creating Logger from cores
	// And sugaring
	logger := zap.New(core)
	sugarlogger := logger.Sugar()
	fmt.Println(sugarlogger.Level())

	// First log message
	// That tells us that logger construction succeeded
	defer sugarlogger.Sync()
	sugarlogger.Debug("Logger construction succeeded")

	return sugarlogger, levels
}

// Useful for small apps where you want to log a bit
// Not sure about async
func InitGlobalLogger() {
	//Ignoring ability to change level in runtime for global usecase
	//TODO: it is not hard to add this feature, mb next time
	globalLogger, _ = InitLogger()
}

func Debug(args ...any) {
	globalLogger.Debug(args...)
	globalLogger.Sync()
}

func Info(args ...any) {
	globalLogger.Info(args...)
	globalLogger.Sync()
}
