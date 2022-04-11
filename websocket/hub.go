package websocket

import "log"

type Hub struct {
	clients    map[*Ws]bool
	broadcast  chan MessageWs
	register   chan *Ws
	unregister chan *Ws
	room       map[string][]*Ws
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan MessageWs),
		clients:    make(map[*Ws]bool),
		unregister: make(chan *Ws),
		register:   make(chan *Ws),
		room:       make(map[string][]*Ws),
	}
}

func (h *Hub) JoinRoom(roomId string, ws *Ws) {
	if _, ok := h.room[roomId]; !ok {
		h.room[roomId] = make([]*Ws, 0)
	}
	h.room[roomId] = append(h.room[roomId], ws)
	log.Println("join room", roomId)
}

func (h *Hub) LeaveRoom(roomId string, ws *Ws) {
	if val, ok := h.room[roomId]; ok {
		for index, conn := range val {
			if conn == ws {
				h.room[roomId] = append(val[:index], val[index+1:]...)
				log.Println("leave room", roomId)
				break
			}
		}
	}
}

func (h *Hub) EmitToRoom(roomId string, message MessageWs) {
	conns, ok := h.room[roomId]
	if !ok {
		return
	}

	for _, conn := range conns {
		conn.send <- message
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
