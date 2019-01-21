package main

import (
	"awesomeProject/autrade"
	"awesomeProject/conf"
	"fmt"
	"os"
)

func main() {
	//config := conf.Default
	err := autrade.Run(conf.MainConfig{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
