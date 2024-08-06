package networkcore

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Mopcho/nebulariders-server/common/mopHttp"
	gamecore "github.com/Mopcho/nebulariders-server/game-server/internal/game"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func validateToken(token string) (gamecore.UserData, error) {
	requestUrl := os.Getenv("AUTH_SERVER_ADDRESS") + "/api/auth/verifyToken" + "?token=" + token
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer([]byte{}))

	if err != nil {
		return gamecore.UserData{}, err
	}

	defer resp.Body.Close()
	v := mopHttp.ApiResponse{Error: &mopHttp.ApiError{}}
	err = mopHttp.GetJsonBody(resp.Body, &v)
	if err != nil {
		return gamecore.UserData{}, errors.New("failed reading body")
	}

	if v.Error == (&mopHttp.ApiError{}) {
		return gamecore.UserData{}, errors.New(v.Error.Message)
	}

	userData, err := getUserDataFromToken(token)

	if err != nil {
		return gamecore.UserData{}, err
	}

	return userData, nil
}

func getUserDataFromToken(tokenString string) (gamecore.UserData, error) {
	userData := gamecore.UserData{}

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		log.Fatalf("Error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if sub, ok := claims["sub"].(string); ok {
			userData.ID = sub
			fmt.Printf("Sub: %s\n", sub)
		} else {
			fmt.Println("sub claim not found or not a string")
			return gamecore.UserData{}, errors.New("sub claim not present on token or not a string")
		}

		if username, ok := claims["username"].(string); ok {
			userData.Username = username
		} else {
			fmt.Println("username claim not found or not a string")
		}
	} else {
		log.Fatalf("Error asserting claims as MapClaims")
		return gamecore.UserData{}, errors.New("error asserting claims as MapClaims")
	}

	return userData, nil
}

func HandleWebSocket(game *gamecore.Game) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate token
		token := r.URL.Query().Get("token")
		userData, err := validateToken(token)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(401)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade:", err)
			return
		}

		player := gamecore.NewPlayer(userData.ID, userData.Username, game, conn)
		game.Players[userData.ID] = player
		player.ReadPump()
		player.SendPump()
	}
}