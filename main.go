package main

import (
	"github.com/Amniversary/real-game-question/config"
	"github.com/Amniversary/real-game-question/service"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	service.Run(config.NewConfig())
}
