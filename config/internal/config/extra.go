package config

import (
	"sync"

	"github.com/spf13/viper"
)

// Global config instance
var c Config

// Returns global config instance
func Get() Config {
	return c
}

// Initialaise config process
// Every path in service works around single env WORKSPACE
func InitConfig() {

	once := sync.Once{}
	once.Do(func() {
		pipeline := []initPhase{
			setEnv,
			setFlags,
			handleConfigFile,
			bindFlags,
			fillGlobalConfig,
			setElse,
			doOverride,
		}
		// panics only here
		execute(pipeline)
	})
}

// Two main functions you should change in config code are:
// setEnv() and setFlags()
// See ./example/example.go for additional hints

// Set ENV
// Immediately validate thorough utilitary register*()
func setEnv() error {
	for _, env := range environment {
		err := registerENV(env)
		if err != nil {
			return err
		}
	}
	return nil
}

// Set flags and explicitly define defaults
// Defaults, as stated in constraints, should be *negative
func setFlags() (err error) {
	for _, flag := range flags {
		flag()
	}
	return nil
}

// Callback on config change , aliases etc.
func setElse() (err error) {
	for _, els := range elses {
		err = els()
		if err != nil {
			return err
		}
	}
	return nil
}

// Do not use, this violates constraints
// If there any way to not override - do not override (C) Me
func doOverride() error {
	for _, over := range override {
		viper.Set(over.name, over.value)
	}
	return nil
}

// *defaults
// "","false","no","stop" for string
// 0 for int
// 0.0 for float
// false for bool
