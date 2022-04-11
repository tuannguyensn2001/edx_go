package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type MessageWs struct {
	Channel     string      `json:"channel"`
	Detail      string      `json:"detail"`
	Data        interface{} `json:"data"`
	IsJoinRoom  bool        `json:"isJoinRoom"`
	RoomId      string      `json:"roomId"`
	IsLeaveRoom bool        `json:"isLeaveRoom"`
	UserId      int         `json:"userId"`
}

type Ws struct {
	hub  *Hub
	conn *websocket.Conn
	send chan MessageWs
}

var upgrader = websocket.Upgrader{}

func (ws *Ws) read() {
	defer func() {
		ws.hub.unregister <- ws
		ws.conn.Close()
	}()

	for {
		var message MessageWs
		err := ws.conn.ReadJSON(&message)

		if err != nil {
			log.Println("err ws read json", err)
			break
		}

		if message.IsJoinRoom {
			ws.hub.JoinRoom(message.RoomId, ws)
		} else if message.IsLeaveRoom {
			ws.hub.LeaveRoom(message.RoomId, ws)
		}

		if message.Channel == "comment" {
			//commenttransportsocket.CreateComment(ws)(context.Background(), message)
		}

	}
}

func (ws *Ws) write() {
	for {
		select {
		case message, ok := <-ws.send:
			if !ok {
				ws.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := ws.conn.WriteJSON(message)
			if err != nil {
				return
			}

		}
	}
}

func WsHandler(w http.ResponseWriter, r *http.Request, hub *Hub) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		log.Println(r)
		return true
	}

	c, err := upgrader.Upgrade(w, r, nil)
	
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	ws := &Ws{hub: hub, conn: c, send: make(chan MessageWs)}

	ws.hub.register <- ws

	go ws.read()
	go ws.write()

}
