package example

import (
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Hints
//
// 1. use `mapstructure` as if it is a yaml tag
// 2. viper can map time not only to string but also to time.Duration
//

// Constaraints on ENV , Flags , Config.file and Default values
//
// #ENV
// - Must be defined, otherwise program shouldn't start
// - Lifetime constants, shouldnt be overridden in runtime
// - Can Be *defaulted only in build\makefile, where they explicitly defined
//
// #Config.file
// - Must exist, have same structure as config.Config, otherwise program shouldn't start
// - May be overridden in runtime or exist in multiple variants across sessions
// - Cannot Be Defaulted
//
// #--Flags
// - May not be defined, program should start,
// 		then SHOULD be defaulted by:
// 		[false , 0 , -1 , "NO" , "stop"]
// 		and any other kind of negative value
// - Program-Lifetime constants, cannot be changed in any way in runtime
// - Can Be Defaulted in build\makefile and in code \
// 		Code variant defaulting is preferred

type Config struct {
	BasicField  string `mapstructure:"basic_field"`
	InnerStruct struct {
		Field time.Duration `mapstructure:"field"`
	} `mapstructure:"inner_struct"`
}

var C Config

// Initialaise config process
func InitConfig() {
	setEnv()
	setFlags()
	setConfig()
	pflag.Parse()
	bindFlags()
	fillGlobalConfig()
	setElse()
	override()
}

// Two main functions you should change in config code
// for ENV use register*(), any other variant do not

// Set ENV
// Immediately validate thorough utilitary register*()
//
//	or any other variant
func setEnv() {
	registerENV("CONFIG_FILE")
}

// Set flags and explicitly define defaults
// Defaults, as stated in constraints, should be *negative
func setFlags() {
	pflag.Bool("enable_debug", false, "Define if debug info is enabled")
}

// Callback on config change , aliases etc.
func setElse() {

	// Callback
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	// Aliases example
	viper.RegisterAlias("enable_debug", "debug")
}

// Do not use this, this violates constraints
// If there any way to not override - do not override (C) Me
func override() {
	// Override example
	viper.Set("enable_debug", true)
}

// *defaults
// "","false","no","stop" for string
// 0 for int
// 0.0 for float
// false for bool
