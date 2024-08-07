package gamecore

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type Player struct {
	ID         string          `json:"id"`
	Username   string          `json:"username"`
	BaseAttack int             `json:"baseAttack"`
	Health     int             `json:"health"`
	Conn       *websocket.Conn `json:"-"`
	Game       *Game           `json:"-"`
	X          float64         `json:"x"`
	Y          float64         `json:"y"`
}

func NewPlayer(id string, username string, game *Game, conn *websocket.Conn) *Player {
	return &Player{
		ID:         id,
		Username:   username,
		Conn:       conn,
		Game:       game,
		X:          1.00,
		Y:          1.00,
		BaseAttack: 10,
		Health:     100,
	}
}

func (s *Player) ReadPump() {
	go func() {
		ticker := time.NewTicker(time.Microsecond * 50)
		for range ticker.C {
			_, bytes, err := s.Conn.ReadMessage()

			if err != nil {
				fmt.Println("Error reading message bytes, not proccessing it")
				_ = s.Conn.Close()
				return
			}

			socketMsg := SocketMessage{}
			err = json.Unmarshal(bytes, &socketMsg)

			if err != nil {
				fmt.Println("Can't parse message, not processing it. Maybe its missing \"type\"")
				continue
			}

			message := Message{Type: socketMsg.Type, PlayerID: s.ID, Data: bytes}
			s.Game.receive(message)
		}
	}()
}

func (s *Player) receive(msg interface{}) {
	switch m := msg.(type) {
	case PlayerReceiveDamageMessage:
		s.Health -= m.Damage
		if s.Health <= 0 {
			_ = s.Conn.WriteJSON(NewServerPlayerDeathMsg())
		}
	case PositionMessage:
		s.X = m.X
		s.Y = m.Y
	default:
		fmt.Println("Unknown message type received")
	}
}

func (s *Player) SendPump() {
	go func() {
		ticker := time.NewTicker(time.Millisecond * 500)
		for range ticker.C {
			_ = s.Conn.WriteJSON(WorldState{Players: s.Game.Players})
		}
	}()
}
