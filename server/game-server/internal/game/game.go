package gamecore

import (
	"encoding/json"
)

type Game struct {
	Players map[string]*Player
	Channel chan Message
}

func NewGame() *Game {
	return &Game{
		Players: make(map[string]*Player),
		Channel: make(chan Message),
	}
}

func (s *Game) receive(msg Message) {
	switch msg.Type {
	case "attack":
		attackMessage := AttackMessage{}
		err := json.Unmarshal(msg.Data, &attackMessage)
		if err != nil {
			panic(err) // TODO: Handle properly
		}
		if _, ok := s.Players[attackMessage.EnemyToAttackID]; !ok {
			break // TODO: Handle properly
		}
		s.Players[attackMessage.EnemyToAttackID].receive(PlayerReceiveDamageMessage{Damage: 10})
	case "position":
		positionMessage := PositionMessage{}
		err := json.Unmarshal(msg.Data, &positionMessage)
		if err != nil {
			panic(err) // TODO: Handle properly
		}
		s.Players[msg.PlayerID].receive(positionMessage)
	}
}