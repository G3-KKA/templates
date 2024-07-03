package config

import (
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Phase of config initialisation
type initPhase func() error

type overrideContainer struct {
	name  string
	value any
}

type flagSetter func()
type elseSetter func() error

// Executes every phase, panics on first error
// Program shouldn't start if any phase of configuration fails
func execute(pipeline []initPhase) {
	for _, phase := range pipeline {
		err := phase()
		if err != nil {
			log.Fatalf("Cannot init condfig %e\n", err)
		}
	}
}

// The difference between register() and bind() is
// that register() extends bind() logic
// in this case - validates that env

// Adds validation to env binding
func registerENV(input ...string) error {
	viper.BindEnv(input...)
	for _, env := range input {
		// Type-free validation
		// Not defined integer or bool would be "" as well
		envalue := viper.GetString(env)
		if envalue == "" {
			return fmt.Errorf("%s is not defined", env)
		}
	}
	return nil
}

// Wraps viper.BindPFlags(pflag.CommandLine)
func bindFlags() error {
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return fmt.Errorf("cannot bind flags: %v", err)
	}
	return nil
}

// Wraps viper.BindPFlags(pflag.CommandLine) into panic + os.Exit(1)
func fillGlobalConfig() error {
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("cannot read config file from CONFIG_FILE: %v", err)
	}
	err = viper.Unmarshal(&c, envInConfigValuesHook())
	if err != nil {
		return fmt.Errorf("cannot unmarshal config.file into config.C: %v", err)
	}
	return nil
}

// Parse config file path for  ext
func extFromPath(path string) string {
	dotIndex := strings.LastIndexByte(path, '.')
	if dotIndex == -1 {
		return ""
	}
	return path[dotIndex+1:]
}

// Parse config file path for name
func nameFromPath(path string) string {
	dotIndex := strings.LastIndexByte(path, '.')
	if dotIndex == -1 {
		return ""
	}
	slashIndex := strings.LastIndexByte(path[:dotIndex], '/')
	return path[slashIndex+1 : dotIndex]
}

// Set config file name and extention
// Change only if something breaks
// For ./relative/path/to/config  and //full/path/to/config
// For config    .yaml .json .toml
// Works just fine
func handleConfigFile() error {
	configFileEnv := viper.GetString("CONFIG_FILE")

	// Getting parts of path
	name := nameFromPath(configFileEnv)
	ext := extFromPath(configFileEnv)
	dir := filepath.Dir(configFileEnv)

	// Setting Config
	viper.AddConfigPath(dir)
	viper.SetConfigName(name)
	viper.SetConfigType(ext)
	return nil
}

// Parse ${ENV}/dir/file kind of path,
// Only works if variable to decode is config.path type
func envInConfigValuesHook() viper.DecoderConfigOption {
	hook := mapstructure.DecodeHookFuncType(func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		// Ignoring other data
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(path("")) {
			return data, nil
		}

		// #config.yaml
		// key : ${WORKSPACE}/file/path
		dataString := data.(string)
		dollarIndex := strings.Index(dataString, "$")
		if dollarIndex == -1 {
			errmsg := `'$' not found in variable 
			so that specificaly declared in config struct as config.path type
			which should only be used if variable in config.file 
			uses ${WORKSPACE}/file/path form of path declaration`
			log.Println(errmsg)
			return data, fmt.Errorf(errmsg)
		}

		// Getting indexes of '{' and '}'
		openBracket := strings.Index(dataString[dollarIndex:], "{")
		closeBracket := strings.Index(dataString[openBracket:], "}")

		// After this operation envName == WORKSPACE
		envName := dataString[openBracket+1 : closeBracket+1]

		// ~/user/goapps/thisapp/internal + /file/path
		// ENV we trying to get should not contain '/'
		// and actual data we want to get should start with'/'
		// This is the most common case for any path operation
		// $(pwd)/file/path or WORKSPACE = ..
		path := viper.GetString(envName) + dataString[closeBracket+2:]
		return path, nil
	})
	return viper.DecodeHook(hook)

}
