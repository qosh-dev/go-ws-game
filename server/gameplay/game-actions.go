package gameplay

import (
	"game-app/database"
	"game-app/database/models"

	socketio "github.com/googollee/go-socket.io"
	"golang.org/x/exp/maps"
)

// -------------------------------------------------------

type GameActions struct {
	Server      *socketio.Server
	Connections *map[int]socketio.Socket
}

// -------------------------------------------------------

func (gA *GameActions) Serve() {
	gA.Server.On("SEND_FIREBALL", gA.playerSendFireball)
}

// -------------------------------------------------------

func (gA *GameActions) playerSendFireball(so socketio.Socket, attacketPlayerId int) {
	currentPlayerId := checkAutorization(so)
	if currentPlayerId == 0 {
		so.Emit("INVALID_TOKEN", 0)
		return
	}

	var currentPlayer models.Player
	database.Connection.First(&currentPlayer, "id = ?", currentPlayerId)

	if currentPlayer.Health <= 0 {
		so.Emit("PLAYER_ALREADY_LOSE", 0)
		return
	}

	var attackedPlayer models.Player
	database.Connection.First(&attackedPlayer, "id = ?", attacketPlayerId)

	attackedPlayerConn, ok := (*gA.Connections)[attacketPlayerId]

	if attackedPlayerConn == nil || !ok {
		return
	}

	if attackedPlayer.Health-10 <= 0 {
		if attackedPlayer.Health > 0 {
			attackedPlayer.Health -= 10
			database.Connection.Save(&attackedPlayer)
		}
		attackedPlayerConn.Emit("PLAYER_DIED", 0)
		return
	}

	attackedPlayer.Health -= 10
	database.Connection.Save(&attackedPlayer)

	attackedPlayerConn.Emit("PLAYER_ATTACKED", int(attackedPlayer.Health))
}

func (gA *GameActions) playerJoinGame(playerId int, allConnections map[int]socketio.Socket) {
	var login string
	database.Connection.Table("players").Select("login").Where("id = ?", playerId).Scan(&login) // Retrieve only login

	for key, connection := range allConnections {
		if key == playerId {
			continue
		}
		connection.Emit("PLAYER_JOIN", login)
	}
}

func (gA *GameActions) playerLeaveGame(playerId int, allConnections map[int]socketio.Socket) {
	var login string
	database.Connection.Table("players").Select("login").Where("id = ?", playerId).Scan(&login) // Retrieve only login

	for key, connection := range allConnections {
		if key == playerId {
			continue
		}
		connection.Emit("PLAYER_LEAVE", login)
	}
}

func (gA *GameActions) sendGameStatus(connection socketio.Socket, allConnections map[int]socketio.Socket) {
	activePlayerIds := maps.Keys(allConnections)

	var players []models.Player
	database.Connection.Find(&players, "id In(?)", activePlayerIds)

	var array = []GameStatusPayload{}

	for _, v := range players {
		record := GameStatusPayload{
			Id:     v.ID,
			Login:  v.Login,
			Health: v.Health,
		}
		array = append(array, record)
	}

	connection.Emit("GAME_STATUS", array)
}
