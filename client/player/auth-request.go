package player

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func signUpRequest(login string, password string) error {
	requestURL := "http://localhost:8080/v1/auth/signup"

	jsonBody, err := json.Marshal(map[string]string{
		"login":    login,
		"password": password,
	})

	if err != nil {
		return errors.New("invalid payload")
	}

	res, err := http.Post(requestURL, "application/json", bytes.NewReader(jsonBody))

	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New("player with specified credentials already exist")
	}

	return nil
}

func loginRequest(login string, password string) (*string, error) {
	requestURL := "http://localhost:8080/v1/auth/login"

	jsonBody, err := json.Marshal(map[string]string{
		"login":    login,
		"password": password,
	})

	if err != nil {
		return nil, errors.New("invalid payload")
	}

	res, err := http.Post(requestURL, "application/json", bytes.NewReader(jsonBody))

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("account not found")
	}

	var authorizationCookie *string
	for _, cookie := range res.Cookies() {
		if cookie.Name == "Authorization" {
			authorizationCookie = &cookie.Value
			break
		}
	}

	if authorizationCookie == nil {
		return nil, errors.New("authorization cookie not found in response")
	}

	return authorizationCookie, nil
}
