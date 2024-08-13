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
	BaseAttack int             `json:"base_attack"`
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
		ticker := time.NewTicker(time.Millisecond * 50)
		for range ticker.C {
			_, bytes, err := s.Conn.ReadMessage()

			if err != nil {
				fmt.Println(err)
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
		err := s.Conn.WriteJSON(newPlayerReceiveDamageMessage(m.Damage, m.From, m.AttackType))
		if err != nil {
			return
		}
		s.Health -= m.Damage
		if s.Health <= 0 {
			_ = s.Conn.WriteJSON(NewServerPlayerDeathMsg())
			delete(s.Game.Players, s.ID)
			err := s.Conn.Close()
			if err != nil {
				return
			}
			return
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
		ticker := time.NewTicker(time.Millisecond * 50)
		for range ticker.C {
			playersWithoutMe := filterPlayers(s.Game.Players, s.ID)
			_ = s.Conn.WriteJSON(NewServerWorldStateMessage(WorldState{Players: playersWithoutMe, Me: *s}))
		}
	}()
}

func filterPlayers(players map[string]*Player, playerId string) map[string]*Player {
	newMap := make(map[string]*Player)

	for key, value := range players {
		if key != playerId {
			newMap[key] = value
		}
	}

	return newMap
}
