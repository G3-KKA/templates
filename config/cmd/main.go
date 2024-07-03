package main

import (
	"fmt"
	config "yet-again-templates/config/internal/config"
)

func main() {

	config.InitConfig()
	c := config.Get()
	fmt.Printf("|%s|\n", c.Example)
}
