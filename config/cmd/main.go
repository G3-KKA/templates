package main

import (
	"fmt"
	config "yet-again-templates/config/internal/config"

	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	c := config.Get()
	fmt.Printf("|%s|\n", c.Dummy)
	fmt.Println(viper.GetString("WORKSPACE"))
	fmt.Println(viper.GetString("CONFIG_FILE"))
}
