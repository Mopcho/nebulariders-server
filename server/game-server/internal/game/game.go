package gamecore

import (
	"encoding/json"
	"fmt"
)

type Game struct {
	Players map[string]*Player
}

func NewGame() *Game {
	return &Game{
		Players: make(map[string]*Player),
	}
}

func (s *Game) receive(msg Message) {
	switch msg.Type {
	case "attack":
		attackMessage := AttackMessage{}
		err := json.Unmarshal(msg.Data, &attackMessage)
		if err != nil {
			fmt.Println(err)
			break
		}
		attackedFrom, ok := s.Players[attackMessage.EnemyToAttackID]
		if !ok {
			fmt.Println("No player with this id")
			break
		}
		s.Players[attackMessage.EnemyToAttackID].receive(PlayerReceiveDamageMessage{Damage: 10, From: attackedFrom, AttackType: attackMessage.AttackType})
	case "position":
		positionMessage := PositionMessage{}
		err := json.Unmarshal(msg.Data, &positionMessage)
		if err != nil {
			fmt.Println(err)
			break
		}
		s.Players[msg.PlayerID].receive(positionMessage)
	}
}
