package networkcore

// import (
// 	"net/http"

// 	"github.com/gorilla/websocket"
// )

// type NetworkActor struct {
// 	conn *websocket.Conn
// }

// func newNetworkActor(route string) *NetworkActor {
// 	http.HandleFunc(route, HandleWebSocket(game))
// 	return &NetworkActor{
// 		conn: conn,
// 	}
// }

// func (s *NetworkActor) Receive(msg interface{}) {
// 	s.conn.WriteJSON(msg)
// }