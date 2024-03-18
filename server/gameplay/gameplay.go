package gameplay

import (
	"fmt"
	"game-app/auth"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	socketio "github.com/googollee/go-socket.io"
)

// ----------------------------------------------------------------------------------------

type Gameplay struct {
	Engine      *gin.Engine
	gameActions GameActions
	roomName    string
}

var connections = make(map[int]socketio.Socket)

// ----------------------------------------------------------------------------------------

func (g *Gameplay) Serve() {

	server, err := socketio.NewServer(nil)
	g.Engine.GET("/socket.io/*any", gin.WrapH(server))
	g.Engine.POST("/socket.io/*any", gin.WrapH(server))

	if err != nil {
		panic("error establishing new socket.io server")
	}

	g.gameActions = GameActions{Server: server, Connections: &connections}
	g.gameActions.Serve()
	g.roomName = "arena"

	// -------------------------------------------------------

	server.On("connection", func(so socketio.Socket) {
		playerId := checkAutorization(so)
		if playerId == 0 {
			so.Emit("INVALID_TOKEN", 0)
			fmt.Println("INVALID TOKEN")
			return
		}
		so.Join(g.roomName)
		connections[playerId] = so
		g.gameActions.sendGameStatus(so, connections)
		g.gameActions.playerJoinGame(playerId, connections)
	})

	server.On("disconnection", func(so socketio.Socket) {
		playerId := checkAutorization(so)
		if playerId == 0 {
			so.Emit("INVALID_TOKEN", 0)
			fmt.Println("INVALID TOKEN")
			return
		}
		so.Leave(g.roomName)
		delete(connections, playerId)
		g.gameActions.playerLeaveGame(playerId, connections)
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	// -------------------------------------------------------

}

func checkAutorization(so socketio.Socket) int {
	playerToken := so.Request().URL.Query().Get("token")
	token, _ := auth.ExtractToken(playerToken)
	playerId := 0
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			playerId = 0
		} else {
			playerId = int(claims["sub"].(float64))
		}
	}
	return playerId
}
