package main

import (
	"log"
	"net/http"

	gamecore "github.com/Mopcho/nebulariders-server/game-server/internal/game"
	networkcore "github.com/Mopcho/nebulariders-server/game-server/internal/network"
	"github.com/joho/godotenv"
)



func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	game := gamecore.NewGame()
	http.HandleFunc("/ws", networkcore.HandleWebSocket(game))
	log.Println("Starting game server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}