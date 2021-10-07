package main

import (
	"fmt"
	"log"

	"github.com/logica0419/scheduled-messenger-bot/config"
	"github.com/logica0419/scheduled-messenger-bot/router"
)

func main() {
	fmt.Print("Initializing Scheduled Messenger Bot...")

	if err := config.GetConfig(); err != nil {
		log.Panic(err)
	}

	r := router.Setup()

	r.Start()
}
