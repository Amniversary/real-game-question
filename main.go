package main

import (
	"github.com/Amniversary/real-game-question/config"
	"github.com/Amniversary/real-game-question/service"
)

func main() {
	service.Run(config.NewConfig())
}
