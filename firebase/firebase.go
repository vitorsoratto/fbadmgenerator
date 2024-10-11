package firebase

import (
	"context"
	"errors"
	"os"

	"golang.org/x/oauth2/google"
)

const firebaseScope = "https://www.googleapis.com/auth/firebase.messaging"

var JSONCredentials string

func NewTokenProvider() (*string, error) {
	jsonKey, err := os.ReadFile(JSONCredentials)
	if err != nil {
		return nil, errors.New("fcm: failed to read credentials file at: " + JSONCredentials)
	}
	cfg, err := google.JWTConfigFromJSON(jsonKey, firebaseScope)
	if err != nil {
		return nil, errors.New("fcm: failed to get JWT config for the firebase.messaging scope")
	}
	ts := cfg.TokenSource(context.Background())

	t, err := ts.Token()
	if err != nil {
		return nil, errors.New("fcm: failed to generate Bearer token")
	}
	token := t.AccessToken

	err = writeToken(token)
	if err != nil {
		return nil, errors.New("fcm: failed to write token to file")
	}

	return &token, nil
}

func writeToken(token string) error {
	file, err := os.Create("token.txt")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(token)

	return nil
}
