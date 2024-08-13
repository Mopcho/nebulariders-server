package gamecore

type UserData struct {
	ID       string
	Username string
}

type SocketMessage struct {
	Type string `json:"type"`
}

type Message struct {
	Type     string `json:"type"`
	PlayerID string `json:"player_id"`
	Data     []byte
}

type AttackMessage struct {
	Message
	EnemyToAttackID string `json:"enemy_to_attack_id"`
	AttackType      string `json:"attack_type"`
}

type PositionMessage struct {
	Message
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type PlayerReceiveDamageMessage struct {
	Type       string  `json:"type"`
	Damage     int     `json:"damage"`
	From       *Player `json:"from"`
	AttackType string  `json:"attack_type"`
}

func newPlayerReceiveDamageMessage(damage int, from *Player, attackType string) *PlayerReceiveDamageMessage {
	return &PlayerReceiveDamageMessage{
		Type:       "receive_damage",
		From:       from,
		Damage:     damage,
		AttackType: attackType,
	}
}

type PlayerDeathWSMessageStruct struct {
	Type string `json:"type"`
}

type PlayerPositionWSMessageStruct struct {
	Type string  `json:"type"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func NewServerPlayerDeathMsg() PlayerDeathWSMessageStruct {
	return PlayerDeathWSMessageStruct{Type: "death"}
}

func NewServerPlayerPositionMsg(position Position) PlayerPositionWSMessageStruct {
	return PlayerPositionWSMessageStruct{Type: "death", X: position.X, Y: position.Y}
}

type WorldStateWSMessage struct {
	Type       string     `json:"type"`
	WorldState WorldState `json:"data"`
}

func NewServerWorldStateMessage(state WorldState) WorldStateWSMessage {
	return WorldStateWSMessage{Type: "world_state", WorldState: state}
}

type WorldState struct {
	Players map[string]*Player `json:"players"`
	Me      Player             `json:"me"`
}
