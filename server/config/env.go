package config

// Specify  Env variables here

type Env struct {
	PORT                 string `validate:"required,numeric"`
	TOKEN                string `validate:"required,alphanum"`
	SECRET               string `validate:"required,alphanum"`
	DB_CONNECTION_STRING string `validate:"required"`
}
