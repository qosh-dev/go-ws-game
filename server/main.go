package main

import (
	"game-app/config"
	"game-app/database"

	"game-app/auth"
	"game-app/gameplay"

	"github.com/gin-gonic/gin"
)

func main() {
	Gin := gin.Default()
	config := config.Config{}
	config.Initialize()
	var db = database.Database{Config: config}
	var gatewayConnection = gameplay.Gameplay{Engine: Gin}
	var authController = auth.AuthController{Engine: Gin}

	db.Initialize()

	authController.Serve()
	gatewayConnection.Serve()
	Gin.Run(":" + config.Env.PORT)
}
