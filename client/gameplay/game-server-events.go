package gameplay

import (
	"fmt"
	"os"
	"sync"

	gosocketio "github.com/graarh/golang-socketio"
)

// ------------------------------------------------------------------------------------------

type GameServerEvents struct {
}

// ------------------------------------------------------------------------------------------

func (gA *GameServerEvents) Serve(client *gosocketio.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		client.On("GAME_STATUS", logMessage(gameStatus))
		client.On("PLAYER_JOIN", logMessage(playerJoin))
		client.On("PLAYER_LEAVE", logMessage(playerLeave))
		client.On("PLAYER_ALREADY_LOSE", logMessage(playerAlreadyLose))
		client.On("PLAYER_DIED", logMessage(playerDied))
		client.On("PLAYER_ATTACKED", logMessage(playerAttacked))
		client.On("INVALID_TOKEN", logMessage(invalidToken))
	}
}

// ------------------------------------------------------------------------------------------

func gameStatus(h *gosocketio.Channel, data []GameStatusPayload) {
	fmt.Println("Current game status")
	for index, v := range data {
		fmt.Printf(">> %d) Player: '%s', id: %d, health: '%d'\n", index+1, v.Login, v.Id, v.Health)
	}
}

func playerJoin(h *gosocketio.Channel, playerName string) {
	fmt.Printf(">> New player '%s' join game\n", playerName)
}

func playerLeave(h *gosocketio.Channel, playerName string) {
	fmt.Printf(">> Player '%s' leave game\n", playerName)
}

func playerDied(h *gosocketio.Channel, health uint) {
	fmt.Printf(">> Unfortunately your magician died")
	h.Close()
	os.Exit(1)
}

func playerAlreadyLose(h *gosocketio.Channel, health uint) {
	fmt.Printf(">> Unfortunately your magician died, you cannot attack other players")
}

func playerAttacked(h *gosocketio.Channel, newHealth uint) {
	fmt.Printf(">> You have been attacked, your health: %d\n", newHealth)
}

func invalidToken(h *gosocketio.Channel, newHealth uint) {
	fmt.Printf(">> Something wrong with your token ")
}

// ----------------------------------------------------------------------------------------------------

func logMessage[T interface{}](fn func(h *gosocketio.Channel, data T)) func(h *gosocketio.Channel, data T) {
	return func(h *gosocketio.Channel, data T) {
		fmt.Println()
		fmt.Println()

		fn(h, data)

		fmt.Println()
		fmt.Println()
		fmt.Print("Please type action(type (help) to get command list): ")
	}
}
