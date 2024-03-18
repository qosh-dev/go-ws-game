package auth

import (
	"game-app/auth/dto"
	"game-app/database"
	"game-app/database/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	Engine *gin.Engine
}

// ----------------------------------------------------------------------------------------

func (c *AuthController) Serve() {
	v1 := c.Engine.Group("/v1/auth")
	{
		v1.POST("/signup", c.signUp)
		v1.POST("/login", c.login)
	}
}

// ----------------------------------------------------------------------------------------

func (contr *AuthController) signUp(c *gin.Context) {

	var body dto.SignUpDTO

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}

	player := models.Player{Login: body.Login, Password: string(hash)}

	result := database.Connection.Create(&player)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create player.",
		})
	}

	c.JSON(http.StatusCreated, gin.H{})
}

// ------------------------------------------

func (contr *AuthController) login(c *gin.Context) {
	var body dto.LoginDTO

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var player models.Player

	database.Connection.First(&player, "login = ?", body.Login)

	if player.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid login or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid login or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": player.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}
