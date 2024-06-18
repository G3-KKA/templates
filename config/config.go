package config

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

/*
	hints

1. use `mapstructure` as if it is a struct tag
2. viper can map time not only to string but also to time.Duration
*/
/*
Core idea about ENV , Config , Flags and Default values

#ENV
.Must be defined, otherwise program shouldn't start
.Program-Lifetime constants, cannot be changed in any way in runtime
.Can Be Defaulted only in build\makefile, where they explicitly defined

#Config (( or Session.Config , if can ))
.Must be defined, otherwise program shouldn't start
.May be overridden in runtime or exist in multiple variants across sessions
.Can Be Defaulted in build\makefile and in code \
	Build/Makefile variant defaulting is preferred

#Flags
.May not be defined , program should start ,
	then SHOULD be defaulted by:
	[false , 0 , -1 , "NO" , "stop"]
	and any other kind of negative value
.Program-Lifetime constants, cannot be changed in any way in runtime






*/
/* PATH_TO_CONFIG should be defined ONLY in ENV */
type Config struct {
	dynamicConfig
	f falgs
	e env
	// __TODO: Add more config types
	// __TODO: Add more config types

}
type falgs struct {
}

type env struct {
	CONFIG_PATH    string
	StorageConnect string
	StoragePath    string
	HTTPServer     struct {
		Address string
		Port    string
	}
}
type dynamicConfig struct {
	HTTPDynamic struct {
		Timeout     time.Duration `mapstructure:"timeout"`
		IdleTimeout string        `mapstructure:"idle_timeout"`
	} `mapstructure:"http_dynamic"`
}

// go run *.go --flagname 444 // flagname=444
// go run *.go  //  flagname=1234
// __TODO: Logic for config initialiser
func InitConfig() {

	setEnv()
	serFlags()
	pflag.Parse()
	bindFlags()
	/*
		setConfigFile()
		setElse()
		setDefaults() */
	readConfig()
	validate()
}

func setEnv() {

}

/*  паника здесь приемлима  */
func bindFlags() {
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Printf("Cannot read config file from path(s), error: %v\n", err)
		panic(err)
	}
}
func readConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Cannot read config file from path(s), error: %v\n", err)
		panic(err)
	}
}
func setConfigOptions() {
	/* === Config file === */
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
}
func setConfigDefaults() {
	/*  === Default values === */
	viper.SetDefault("workdir", "./")

}

func setConfigEnvAndCommandLine() {
	/* === Environment variables === */
	viper.MustBindEnv("GOVERSION", "GOVERSION")
	viper.BindEnv("ZZGOSRC", "ZZGOSRC", "MYGOSRC", "ANYOTHERALIAS")
	/* === Command line arguments === */
	pflag.Int("flagname", 1234, "flagname")

}
func setConfigElse() {
	/* === Watch config file changes === */
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
	viper.Get("workdir")

}
func setConfigValidate() {

}
