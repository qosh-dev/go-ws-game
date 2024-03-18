package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Env Env
}

// -------------------------------------------------------------------------------------------

func (c *Config) Initialize() {
	c.extractEnv()
	c.validateEnv()
}

// -------------------------------------------------------------------------------------------

func (c *Config) validateEnv() {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(c.Env)
	if err != nil {
		panic(err)
	}
}

func (c *Config) extractEnv() {
	_, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	errEnv := godotenv.Load(".env")

	if errEnv != nil {
		panic(errEnv)

	}

	enumMap, err := godotenv.Read()

	if err != nil {
		panic(err)
	}
	v := reflect.ValueOf(&c.Env).Elem()

	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		if value, ok := enumMap[fieldName]; ok {
			v.FieldByName(fieldName).Set(reflect.ValueOf(value))
		}
	}
}
