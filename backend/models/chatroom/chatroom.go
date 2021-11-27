package chatroom

import (
	"fmt"
)

type Room struct {
    Register   chan *Client
    Unregister chan *Client
    Clients    map[*Client]bool
    Broadcast  chan Message
}

func NewRoom() *Room {
    return &Room{
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Clients:    make(map[*Client]bool),
        Broadcast:  make(chan Message),
    }
}

func (room *Room) Start() {
    for {
        select {
        case client := <-room.Register:
            // When User Connects, a message is sent out from admin machine
            room.Clients[client] = true
            username := client.UserName
            fmt.Println("Size of Connection room: ", len(room.Clients))
            for client, _ := range room.Clients {
                client.Conn.WriteJSON(Message{Type: 1, User: username, Body: "User Joined"})
            }
            break
        case client := <-room.Unregister:
            username := client.UserName
            delete(room.Clients, client)
            fmt.Println("Size of Connection room: ", len(room.Clients))
            for client, _ := range room.Clients {
                client.Conn.WriteJSON(Message{Type: 1, User: username, Body: "User Disconnected"})
            }
            break
        case message := <-room.Broadcast:
            for client, _ := range room.Clients {
                if err := client.Conn.WriteJSON(message); err != nil {
                    return
                }
            }
        }
    }
}