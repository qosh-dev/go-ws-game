package auth

import (
	"fmt"
	"game-app/database"
	"game-app/database/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticated(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, _ := ExtractToken(tokenString)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var playerId = claims["sub"]
		var player models.Player
		database.Connection.First(&player, playerId)

		if player.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("player", player)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func ExtractToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
}
