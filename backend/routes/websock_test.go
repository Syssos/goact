package routes

import (
    "fmt"
    "strings"
    "testing"
    "net/http"
    "encoding/json"
    "net/http/httptest"

    "github.com/gorilla/websocket"
    "github.com/Syssos/goact/models/chatroom"
)

func TestValidateFMTCheck(t *testing.T) {
    t.Run("Testing Incorrect Cookie name", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/ws", nil)
        tkStr := "IncorrectValueForSure"

        cookie := http.Cookie{Name: "noket", Value: tkStr}
        request.AddCookie(&cookie)

        res, err  := ValidateCookieFMT(request)
        if err == nil {
            t.Error("error not detected for bad token name")
        }

        if res != "" {
            t.Error("Accepted incorrectly named cookie")
        }
    })
    t.Run("Testing Correct Cookie name", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/ws", nil)
        tkStr := GenerateExpired()
        
        cookie := http.Cookie{Name: "token", Value: tkStr}
        request.AddCookie(&cookie)

        res, err  := ValidateCookieFMT(request)
        if err != nil {
            t.Error("Error Validating Cookie")
        }
        if res != tkStr {
            t.Error("Couldn't validate correct string")
        }
    })
}

func TestValidateCookieJWT(t *testing.T) {
    t.Run("Testing Cookie validation with bad token", func(t *testing.T) {
        res, err := ValidateCookieJWT("TheHamburgler")
        if err == nil {
            t.Error("Error not detected while validating JWT String")
        }
        if res != "" {
            t.Error("Validation string not blank with incorrect JWT format")
        }
    })

    t.Run("Testing Cookie validation with expired token", func(t *testing.T) {
        res, err  := ValidateCookieJWT(GenerateExpired())
        if err == nil {
            t.Error("Error not detected while validating JWT String")
        }
        if res != "" {
            t.Error("Could not detect bad token string")
        }
    })

    t.Run("Testing Cookie validation with valid token", func(t *testing.T) {
        res, err  := ValidateCookieJWT(GenerateTokenString())
        if err != nil {
            t.Error("Error Validating JWT String")
        }

        if res != "TestUser1" {
            t.Error("Could not validate good token string, or obtain username")
        }
    })
}

func TestUpgradeConn(t *testing.T) {
    t.Run("Testing if connection is upgraded to TCP", func (t *testing.T) {
        s := httptest.NewServer(http.HandlerFunc(echo))
        defer s.Close()

        // Convert http://127.0.0.1 to ws://127.0.0.
        u := "ws" + strings.TrimPrefix(s.URL, "http")

        // Connect to the server
        ws, _, err := websocket.DefaultDialer.Dial(u, nil)
        if err != nil {
            t.Fatalf("%v", err)
        }
        defer ws.Close()

        // Send message to server, read response and check to see if it's what we expect.
        for i := 0; i < 10; i++ {
            if err := ws.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
                t.Fatalf("%v", err)
            }
            _, p, err := ws.ReadMessage()
            if err != nil {
                t.Fatalf("%v", err)
            }
            if string(p) != "hello" {
                t.Fatalf("bad message")
            }
        }
    })
}

func TestChatroomPoolSize(t *testing.T) {
    t.Run("Checking chatroom pool size", func(t *testing.T) {
        newroom := chatroom.NewRoom()
        go newroom.Start()
        
        s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
            cookie := http.Cookie{Name: "token", Value: GenerateTokenString()}
            r.AddCookie(&cookie)

            serveWs(newroom, w, r)
        }))

        defer s.Close()

        // Convert http://127.0.0.1 to ws://127.0.0.
        u := "ws" + strings.TrimPrefix(s.URL, "http")

        // Connect to the server
        ws, _, err := websocket.DefaultDialer.Dial(u, nil)
        if err != nil {
            t.Fatalf("%v", err)
        }
        defer ws.Close()

        if len(newroom.Clients) == 0 {
            t.Error("The client was not added to chat pool")
        }
    })
}

func TestChatroomMSGBroadcast(t *testing.T) {
    t.Run("Checking msg broadcasting", func (t *testing.T){
        newroom := chatroom.NewRoom()
        go newroom.Start()
        
        s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
            cookie := http.Cookie{Name: "token", Value: GenerateTokenString()}
            r.AddCookie(&cookie)

            serveWs(newroom, w, r)
        }))

        // Convert http://127.0.0.1 to ws://127.0.0.
        u := "ws" + strings.TrimPrefix(s.URL, "http")

        // User1
        ws, _, err := websocket.DefaultDialer.Dial(u, nil)
        if err != nil {
            t.Fatalf("%v", err)
        }

        // User2
        ws2, _, err := websocket.DefaultDialer.Dial(u, nil)
        if err != nil {
            t.Fatalf("%v", err)
        }

        if len(newroom.Clients) != 2 {
            t.Error("The clients were not added to chat pool")
        }

        responseList := []string{}
        // Send message to server, read response and check to see if it's what we expect.
        for i := 0; i < 10; i++ {
            if err := ws.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
                t.Fatalf("%v", err)
            }
            
            _, p, err := ws2.ReadMessage()
            if err != nil {
                t.Fatalf("%v", err)
            }
            
            var msg chatroom.Message
            if err := json.Unmarshal(p, &msg); err != nil {
                t.Errorf("Couldn't serialize mseeage data: %v", err)
            }

            // Account for message from user joining
            joinStr := "User Joined"
            helloStr := "hello"
            
            if msg.Body != helloStr {
                if msg.Body != joinStr {
                    str := fmt.Sprintf("MSG Body: |%v| joinStr: |%v| helloStr: |%v|\n", msg.Body, joinStr, helloStr)
                    t.Errorf("bad message: %v", str)
                }
            }
            responseList = append(responseList, msg.Body)
        }
        if len(responseList) == 9 {
            t.Error("Incorrect amount of messages sent")
        }
        ws.Close()
        ws2.Close()
        s.Close()
    })
}

func echo(w http.ResponseWriter, r *http.Request) {
    c, err := Upgrade(w, r)
    if err != nil {
        return
    }

    defer c.Close()
    for {
        mt, message, err := c.ReadMessage()
        if err != nil {
            break
        }
        err = c.WriteMessage(mt, message)
        if err != nil {
            break
        }
    }
}