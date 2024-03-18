package main

import (
	"fmt"
	"game-client/gameplay"
	playerService "game-client/player"
)

func main() {

	fmt.Println("Welcome to game")
	fmt.Println("To enter the game first need to log in")

	player := playerService.PlayerService{}
	player.InitializePlayer()

	connection := gameplay.Gameplay{Player: player}

	connection.Serve()

}
