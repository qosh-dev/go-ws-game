package models

type Player struct {
	ID       uint   `gorm:"primaryKey"`
	Login    string `gorm:"unique"`
	Password string
	Health   int8 `gorm:"default:100"`
}
