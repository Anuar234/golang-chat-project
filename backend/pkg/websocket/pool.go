package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Printf("Size of connection pool: %d\n", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println("Client:", client)
				if err := client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."}); err != nil {
					fmt.Println("Error sending join message:", err)
				}
			}
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Printf("Size of connection pool: %d\n", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println("Client:", client)
				if err := client.Conn.WriteJSON(Message{Type: 1, Body: "User disconnected..."}); err != nil {
					fmt.Println("Error sending disconnect message:", err)
				}
			}
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println("Error broadcasting message:", err)
					// Не возвращаемся, продолжаем отправку другим клиентам
				}
			}
		}
	}
}