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
	PlayerID string `json:"playerId"`
	Data     []byte
}

type AttackMessage struct {
	Message
	EnemyToAttackID string `json:"enemyToAttackId"`
}

type PositionMessage struct {
	Message
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type PlayerReceiveDamageMessage struct {
	Damage int
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
