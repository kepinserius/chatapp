package websocket

type BroadcastMessage struct {
	Message  []byte
	RoomID   uint
	Username string
}

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan *BroadcastMessage
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan *BroadcastMessage),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			// Send online users list
			h.sendOnlineUsers()
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				// Send updated online users list
				h.sendOnlineUsers()
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				if client.RoomID == message.RoomID {
					select {
					case client.Send <- message.Message:
					default:
						close(client.Send)
						delete(h.Clients, client)
					}
				}
			}
		}
	}
}

func (h *Hub) sendOnlineUsers() {
	// Implement logic to send online users list
}
