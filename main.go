package main

import (
	"fmt"
	"os"

	"github.com/apodhrad/ubi-init/log"
)

func main() {
	fmt.Printf(FIGLET)
	dir, err := os.Getwd()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	err = copyTemplate("micro", dir)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	log.Info("UBI was successfully initialized")
}
