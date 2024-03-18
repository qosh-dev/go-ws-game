package database

import (
	"game-app/config"
	"game-app/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB

type Database struct {
	Config config.Config
}

// -------------------------------------------------------------------------------------------

func (d *Database) Initialize() {
	d.connect()
	d.syncModels()
}

func (d *Database) GetConnection() *gorm.DB {
	return Connection
}

// -------------------------------------------------------------------------------------------

func (d *Database) connect() {
	var err error
	conn, err := gorm.Open(postgres.Open(d.Config.Env.DB_CONNECTION_STRING), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database")
	}

	Connection = conn

}

func (d *Database) syncModels() {
	Connection.AutoMigrate(&models.Player{})
}
