package zapp

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.SugaredLogger

type LevelWithName struct {
	zap.AtomicLevel
	Name string
}

func withName(name string, level zap.AtomicLevel) LevelWithName {
	return LevelWithName{level, name}
}
func InitLogger() (*zap.SugaredLogger, []LevelWithName) {
	log.Println("Logger initialization started")

	// Creating encoders
	encoder := setEncoder()

	// Creating log file
	logfile := getLogFile()

	// Creating levels to be able to change level in runtime
	levels := []LevelWithName{
		withName("stdout", zap.NewAtomicLevelAt(zapcore.InfoLevel)),   // 0
		withName("logfile", zap.NewAtomicLevelAt(zapcore.DebugLevel)), // 1
	}

	// Creating zap.Cores
	// And merging them
	logfilecore := zapcore.NewCore(encoder, logfile, levels[0])
	stderr := zapcore.NewCore(encoder, os.Stderr, levels[1])
	core := zapcore.NewTee(logfilecore, stderr)

	// Creating Logger from cores
	// And sugaring

	logger := zap.New(core)
	sugarlogger := logger.Sugar()

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

// Logs to stderr
func Debug(args ...any) {
	globalLogger.Debug(args...)
	globalLogger.Sync()
}

// Logs to stderr and logfile
func Info(args ...any) {
	globalLogger.Info(args...)
	globalLogger.Sync()
}
