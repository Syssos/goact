package websocket

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
    Pool      *Pool
}

type Message struct {
    Type int       `json:"type"`
    User uuid.UUID `json:"user"`
    Body string    `json:"body"`
}

func (c *Client) Read() {
    defer func() {
        c.Pool.Unregister <- c
        c.Conn.Close()
    }()

    for {
        messageType, p, err := c.Conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        message := Message{Type: messageType, User: c.ID,Body: string(p)}
        c.Pool.Broadcast <- message
        fmt.Printf("Message Received: %v %v\n", message.User, message.Body)
    }
}