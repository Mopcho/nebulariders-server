package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type PlayerActor struct {
	GameActor   *GameActor    `json:"-"`
	Conn        *websocket.Conn `json:"-"`
	Health      int           `json:"health"`
	BasicAttack int           `json:"basicAttack"`
	ID          uuid.UUID     `json:"id"`
	LoggedIn    bool          `json:"loggedIn"`
}

type Message struct {
	Type        string `json:"type"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	BasicAttack string `json:"basicAttack"`
	PlayerId    string `json:"playerId"`
}

type AuthResponse struct {
	OK bool `json:"ok"`
}

type PlayerKilledResponse struct {
	Type     string `json:"type"`
	PlayerId string `json:"playerId"`
}

type PlayerDeathResponse struct {
	Type     string `json:"type"`
}

type StartGameResponse struct {
	Type    string                  `json:"type"`
	Players map[string]*PlayerActor `json:"players"` 
}

func getBytes(obj interface{}, conn *websocket.Conn) []byte {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Printf("Error marshaling object: %v", err)
		conn.Close()
		return nil
	}
	return bytes
}

func findUserByUsername(m map[string]Auth, username string) (Auth, bool) {
	for _, v := range m {
		if v.Username == username {
			return v, true
		}
	}
	return Auth{}, false
}

func findOtherPlayers(players map[string]*PlayerActor, currentPlayerID string) map[string]*PlayerActor {
	filteredPlayers := make(map[string]*PlayerActor)
	for id, player := range players {
		if id != currentPlayerID {
			filteredPlayers[id] = player
		}
	}
	return filteredPlayers
}

func (s *PlayerActor) Receive(msg interface{}) {
	switch m := msg.(type) {
	case Message:
		switch m.Type {
		case "login":
			user, ok := findUserByUsername(s.GameActor.Auth, m.Username)
			if !ok || user.Password != m.Password {
				bytes := getBytes(AuthResponse{OK: false}, s.Conn)
				s.Conn.WriteMessage(websocket.TextMessage, bytes)
				return
			}

			s.LoggedIn = true
			s.ID = uuid.New()
			s.GameActor.Players[s.ID.String()] = s
			bytes := getBytes(AuthResponse{OK: true}, s.Conn)
			s.Conn.WriteMessage(websocket.TextMessage, bytes)
			go s.sendPump()
		case "register":
			id := uuid.New()
			s.GameActor.Auth[id.String()] = Auth{
				Username: m.Username,
				Password: m.Password,
			}
			bytes := getBytes(AuthResponse{OK: true}, s.Conn)
			s.Conn.WriteMessage(websocket.TextMessage, bytes)
		case "start":
			bytes := getBytes(StartGameResponse{Players: s.GameActor.Players, Type: "start"}, s.Conn)
			s.Conn.WriteMessage(websocket.TextMessage, bytes)
		case "basic_attack":
			enemyPlayer, ok := s.GameActor.Players[m.PlayerId]
			if !ok {
				log.Printf("Player %s not found", m.PlayerId)
				return
			}

			enemyPlayer.Health -= s.BasicAttack

			if enemyPlayer.Health <= 0 {
				bytes := getBytes(PlayerKilledResponse{PlayerId: enemyPlayer.ID.String(), Type: "player_killed"}, s.Conn)
				bytesPlayerDead := getBytes(PlayerDeathResponse{Type: "dead"}, s.Conn)
				enemyPlayer.Conn.WriteMessage(websocket.TextMessage, bytesPlayerDead)
				s.Conn.WriteMessage(websocket.TextMessage, bytes)
			}
		default:
			fmt.Println("Received unknown message type:", m.Type)
		}

	default:
		fmt.Println("Received unknown type")
	}
}

func (s *PlayerActor) sendPump() {
	ticker := time.NewTicker(time.Second * 1)
	defer func() {
		ticker.Stop()
		s.Conn.Close()
	}()

	for range ticker.C {
		var otherPlayers = findOtherPlayers(s.GameActor.Players, s.ID.String())
		var bytes = getBytes(otherPlayers, s.Conn)
		s.Conn.WriteMessage(websocket.TextMessage, bytes)
	}
}

func (s *PlayerActor) ReadPump() {
	defer func() {
		s.Conn.Close()
	}()

	for {
		var msg Message
		err := s.Conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Websocket error:", err)
			return
		}
		log.Println("Received message:", msg)
		s.Receive(msg)
	}
}
