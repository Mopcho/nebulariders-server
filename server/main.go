package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Auth struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type GameActor struct {
	Players map[string]*PlayerActor `json:"players"`
	Auth    map[string]Auth         `json:"auth"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(gameActor *GameActor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade:", err)
			return
		}
		player := &PlayerActor{Conn: conn, GameActor: gameActor, Health: 100, BasicAttack: 10}
		player.ReadPump()
	}
}

func main() {
	authMap := make(map[string]Auth)
	playersMap := make(map[string]*PlayerActor)
	gameActor := GameActor{Auth: authMap, Players: playersMap}
	http.HandleFunc("/ws", handleWebSocket(&gameActor))
	log.Fatal(http.ListenAndServe(":8080", nil))
}