package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	fmt.Println(viper.GetString("CONFIG_FILE"))
	fmt.Println(C.Me)
}
