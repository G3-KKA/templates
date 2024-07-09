package config

import (
	"time"

	_ "github.com/spf13/viper"
)

// Hints
//
// 1. use `mapstructure` in Config as if it is a yaml/json tag
// 2. viper can map time not only to string but also to time.Duration
// 3. use pflag.* inside []flagSetter

// Constaraints on ENV , flags , config.file and Default values
//
// # ENV
// - Must be defined, otherwise program shouldn't start
// - Lifetime constants, shouldnt be overridden in runtime
// - Cannot be defaulted
//
// # config.file
// - Must exist, have same structure as config.Config, otherwise program shouldn't start
// - May be overridden in runtime or exist in multiple variants across sessions
// - Cannot Be Defaulted
//
// # --flags
//   - May not be defined, program should start,
//   - Lifetime constants, shouldnt be overridden in runtime
//   - Can and should be defaulted by:
//	   [false , 0 , -1 , "NO" , "stop"]
//   	and any other kind of negative value

// Use this string *alias to be able to decode env in config.file
// See config.utilitary#envInConfigValuesHook for details
// Brief example of usage:
// WORKSPACE = ~/user/goapp
// ${WORKSPACE}/file/path => ~/user/goapp/file/path
type path string

// Represents config file, must be changed manually
// Only public fields
type Config struct {
	Logger struct {
		SyncTimeout time.Duration `mapstructure:"SyncTimeout"`
		Cores       []struct {
			Name           string `mapstructure:"Name"`         // name of the core, for idientification reasons
			EncoderLevel   string `mapstructure:"EncoderLevel"` // production or development
			Path           path   `mapstructure:"Path"`         // everything that getLogFile can handle
			Level          int    `mapstructure:"Level"`        // level of the core , might be negative
			MustCreateCore bool   `mapstructure:"MustCreateCore"`
		} `mapstructure:"Cores"`
	} `mapstructure:"Logger"`
}

var environment = [...]string{
	// Every path in service works around WORKSPACE
	// Removing this will break every env-based path in service
	"WORKSPACE",
	"CONFIG_FILE",
}

// Use pfalg.*
var flags = [...]flagSetter{}

// Other viper options
var elses = [...]elseSetter{}

// Uses	viper.Set
var override = [...]overrideContainer{}
