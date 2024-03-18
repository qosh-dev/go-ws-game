package gameplay

import (
	"log"

	gosocketio "github.com/graarh/golang-socketio"
)

// ------------------------------------------------------------------------------------------

type GameActions struct {
	serverEvents   GameServerEvents
	clientCommands GameClientCommands
}

// ------------------------------------------------------------------------------------------

func (gA *GameActions) Initialize() {
	gA.clientCommands = GameClientCommands{}
	gA.serverEvents = GameServerEvents{}
}

func (gA *GameActions) OnConnect(channel *gosocketio.Channel) {
	// log.Println("\nOnConnect")
}

func (gA *GameActions) OnDisconnect(channel *gosocketio.Channel) {
	log.Fatal("\nDisconnect")
}

// ------------------------------------------------------------------------------------------
