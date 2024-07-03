package config

import "sync"

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
			override,
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
func setFlags() error { return nil }

// Callback on config change , aliases etc.
func setElse() error { return nil }

// Do not use, this violates constraints
// If there any way to not override - do not override (C) Me
func override() error { return nil }

// *defaults
// "","false","no","stop" for string
// 0 for int
// 0.0 for float
// false for bool
