package player

import (
	"fmt"
	"game-client/utils"
)

var authToken *string

type PlayerService struct {
}

func (s *PlayerService) GetToken() *string {
	return authToken
}

func (s *PlayerService) InitializePlayer() *string {
	var action string

	for {
		action = utils.ReadLine("Please choose action(signup(s) or login(l)): ")
		if action == "s" || action == "l" {
			break
		}
		fmt.Println("\nPlease choose one of available actions(signup(s) or login(l))")
	}

	fmt.Println("Please enter your credentials ")

	login := utils.ReadLine("Enter your login: ")
	password := utils.ReadLine("Enter your password: ")

	var err error

	if action == "s" {
		err = signUpRequest(login, password)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Try again")
			return s.InitializePlayer()
		}
	}

	authToken, err = loginRequest(login, password)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Try again")
		return s.InitializePlayer()
	}

	return authToken
}
