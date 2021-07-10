package main

import (
    "fmt"
    "time"
    "net/http"
    "crypto/sha1"

    "github.com/google/uuid"
    "github.com/Syssos/goact/pkg/websocket"
)

// User
var UserList = make(map[string]string)
var UserCookie = make(map[string]string)

func Validate(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/validate" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
    w.Header().Set("Access-Control-Allow-Origin", "*")
    if (r.Method == "OPTIONS") {
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        w.Header().Set("Access-Control-Allow-Headers", "User-Name, User-Secret") // You can add more headers here if needed
    } else {
        username := r.Header["User-Name"][0]
        password := r.Header["User-Secret"][0]
        ValidUser := checkUser(username, password)
        if ValidUser {

            h := sha1.New()
            h.Write([]byte(username))
            cookVal := fmt.Sprintf("%x", h.Sum(nil))
            UserCookie[username] = cookVal
            
            expiration := time.Now().Add(365 * 24 * time.Hour)
            cookie := http.Cookie{Name: username, Value: cookVal, Expires: expiration}
            http.SetCookie(w, &cookie)
            
        } else {
            w.Header().Set("User", "Un-Validated")
            w.WriteHeader(401)
        }
    }    
}

func checkUser(u string, p string) bool {
    UserValid := false

    for user, pass := range UserList {
        if user == u {
            fmt.Println("User Pass", pass)
            UserValid = true
            return UserValid
        }
    }

    return UserValid
}

// tut
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
    UserList["Test"] = "test"
    http.HandleFunc("/validate", Validate)
    http.ListenAndServe(":8080", nil)
}