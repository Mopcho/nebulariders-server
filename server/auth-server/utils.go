package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func sendResponse(w http.ResponseWriter, data ApiResponse) {
	var bytes, err = json.Marshal(data)

	if err != nil {
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Write(bytes)
}

func getJsonBody(w http.ResponseWriter, body io.ReadCloser, v any) error {
	var err = json.NewDecoder(body).Decode(&v)
	if err != nil {
		return errors.New("failed decoding request body")
	}

	return nil
}

func createToken(sub string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": sub,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}