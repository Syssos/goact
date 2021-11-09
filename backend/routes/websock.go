package routes

import (
    "fmt"
    "net/http"
  
    "github.com/google/uuid"
    // "github.com/dgrijalva/jwt-go"
    "github.com/Syssos/goact/models/chatroom"
)

func serveWs(room *chatroom.Room, w http.ResponseWriter, r *http.Request) {
    
    tokenStr := ValidateCookieFMT(r)
    if tokenStr != "" {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    ValidatedUser := ValidateCookieJWT(tokenStr)
    if ValidatedUser != "" {        
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