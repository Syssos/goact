package chatroom

import (
	"fmt"
    "log"
	
    "github.com/google/uuid"
    "github.com/gorilla/websocket"
)

type Client struct {
    ID        uuid.UUID
    UserName  string
    Conn      *websocket.Conn
    Room      *Room
}

type Message struct {
    Type int       `json:"type"`
    User string    `json:"user"`
    Body string    `json:"body"`
}

func (c *Client) Read() {
    defer func() {
        c.Room.Unregister <- c
        c.Conn.Close()
    }()

    for {
        messageType, p, err := c.Conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        message := Message{Type: messageType, User: c.UserName,Body: string(p)}
        c.Room.Broadcast <- message
        fmt.Printf("Message Received: %v %v\n", message.User, message.Body)
    }
}