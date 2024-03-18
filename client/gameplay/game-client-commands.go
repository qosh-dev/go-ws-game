package gameplay

import (
	"fmt"
	"game-client/utils"
	"strconv"
	"sync"

	gosocketio "github.com/graarh/golang-socketio"
)

// ------------------------------------------------------------------------------------------

type GameClientCommands struct {
	events map[string]func(client *gosocketio.Client)
}

// ------------------------------------------------------------------------------------------
func (gA *GameClientCommands) Serve(client *gosocketio.Client, wg *sync.WaitGroup) {
	gA.events = make(map[string]func(client *gosocketio.Client))
	gA.events["SEND_FIREBALL"] = gA.sendFireBall
	gA.events["help"] = gA.help

	defer wg.Done()
	for {
		action := utils.ReadLine("Please type action(type (help) to get command list): ")
		fn, ok := gA.events[action]
		if !ok {
			fmt.Println("Invalid action, please try again")
			continue
		}
		fn(client)
	}
}

// ------------------------------------------------------------------------------------------

func (gA *GameClientCommands) sendFireBall(client *gosocketio.Client) {
	playerIdStr := utils.ReadLine("Type player id: ")
	playerId, err := strconv.Atoi(playerIdStr)

	if err != nil {
		fmt.Println("Invalid playerId, please try again")
		gA.sendFireBall(client)
		return
	}

	client.Emit("SEND_FIREBALL", playerId)
}

func (gA *GameClientCommands) help(client *gosocketio.Client) {
	commands := make([]string, 0, len(gA.events))
	for k := range gA.events {
		commands = append(commands, k)
	}

	fmt.Println("Available commands", len(gA.events))
	for index, v := range commands {
		fmt.Printf("%d) %s \n", index+1, v)
	}
}
