package gameplay

import (
	"fmt"
	playerService "game-client/player"
	"game-client/utils"
	"log"
	"sync"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

type Gameplay struct {
	Player      playerService.PlayerService
	gameActions GameActions
}

func (inst *Gameplay) Serve() {

	inst.gameActions = GameActions{}
	inst.gameActions.Initialize()

	var wg sync.WaitGroup
	wg.Add(2)

	var params []string
	params = append(params, fmt.Sprintf("token=%s", *inst.Player.GetToken()))

	URL := utils.GetUrl("127.0.0.1", 8080, params, false)
	c, err := gosocketio.Dial(URL, transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer wg.Done()
		err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
			inst.gameActions.OnConnect(h)
		})
		err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
			inst.gameActions.OnDisconnect(h)
		})

		if err != nil {
			log.Fatal(err)
		}
	}()

	go inst.gameActions.clientCommands.Serve(c, &wg)
	go inst.gameActions.serverEvents.Serve(c, &wg)

	wg.Wait()
	log.Println("Thank you to play our game")
	log.Println("Exit")
}
