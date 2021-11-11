package routes

import (
    "log"
    "fmt"
    "net/http"
  
    "github.com/google/uuid"
    "github.com/dgrijalva/jwt-go"
    "github.com/gorilla/websocket"
    "github.com/Syssos/goact/models/chatroom"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

func serveWs(room *chatroom.Room, w http.ResponseWriter, r *http.Request) {
    tokenStr := ValidateCookieFMT(r)
    if tokenStr != "" {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    ValidatedUser := ValidateCookieJWT(tokenStr)
    fmt.Println(ValidatedUser)
    if ValidatedUser == "" {        
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    conn, err := Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &chatroom.Client{
        ID:       uuid.New(),
        UserName: ValidatedUser,
        Conn:     conn,
        Room:     room,
    }

    room.Register <- client
    client.Read()
}

func ValidateCookieFMT(r *http.Request) string {
    cookie, err := r.Cookie("Token")
    if err != nil {
        return ""
    }

    return cookie.Value
}

func ValidateCookieJWT(tokenStr string) string {
    claims := &Claims{}
    tkn, err := jwt.ParseWithClaims(tokenStr, claims,
        func(t *jwt.Token) (interface{}, error){
            return jwtKey, nil
        })

    if err != nil {
        fmt.Println(err)
        return ""
    }
    if !tkn.Valid {
        return ""
    }
    return claims.Username
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    return conn, nil
}