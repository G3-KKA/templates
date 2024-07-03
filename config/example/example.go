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

type Config struct {
	BasicField  string `mapstructure:"basic_field"`
	Example     path   `mapstructure:"example"`
	InnerStruct struct {
		Field time.Duration `mapstructure:"field"`
	} `mapstructure:"inner_struct"`
}

var environment = [...]string{
	// Every path in service works around WORKSPACE
	// Removing this will break every env-based path in service
	"WORKSPACE",
	"CONFIG_FILE",
	"GOVERSION",
	"OS",
}

// Use pfalg.*
var flags = [...]flagSetter{
	func() { pflag.Bool("enable_debug", false, "Define if debug info is enabled") },
}

// Other viper options
var elses = [...]elseSetter{
	func() error {

		// Callback
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Println("Config file changed:", e.Name)
		})
		viper.WatchConfig()

		return nil
	},
	func() error {
		// Aliases example
		viper.RegisterAlias("enable_debug", "debug")
		return nil
	},
}

// Uses	viper.Set
var toOverride = [...]overrideContainer{
	{"enable_debug", true},
	{"dummy", "dummy"},
}
