package main

import (
    "fmt"
    "net/http"

    "github.com/google/uuid"
    "github.com/Syssos/goact/pkg/websocket"
)

func Validate(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/validate" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")


    fmt.Println("Hit")
    // Loop over header names
    // for name, values := range r.Header {
    //     // Loop over all values for the name.
    //     for _, value := range values {
    //         fmt.Println(name, value)
    //     }
    // }
    fmt.Println(r.Header.Get("Access-Control-Request-Headers"))
    
}

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
    fmt.Println("WebSocket Endpoint Hit")
    conn, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &websocket.Client{
        ID:       uuid.New(),
        UserName: "Computer",
        Conn:     conn,
        Pool:     pool,
    }

    pool.Register <- client
    client.Read()
}

func setupRoutes() {
    pool := websocket.NewPool()
    go pool.Start()


    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(pool, w, r)
    })
}

func main() {
    fmt.Println("Distributed Chat App v0.01")
    setupRoutes()
    http.HandleFunc("/validate", Validate)
    http.ListenAndServe(":8080", nil)
}