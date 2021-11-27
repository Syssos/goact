package routes

import (
    "fmt"
    "errors"
    "net/http"
  
    "github.com/google/uuid"
    "github.com/dgrijalva/jwt-go"
    "github.com/gorilla/websocket"
    "github.com/Syssos/goact/models/chatroom"
)

// Check origin is insecure due to bypassing CORS checks. This is overwritten for the sake of easily testing.
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

func serveWs(room *chatroom.Room, w http.ResponseWriter, r *http.Request) {
    tokenStr, err := ValidateCookieFMT(r)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    ValidatedUser, err := ValidateCookieJWT(tokenStr)
    if err != nil {
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

func ValidateCookieFMT(r *http.Request) (string, error) {
    cookie, err := r.Cookie("token")
    if err != nil {
        return "", errors.New("Cookie token fmt error")
    }

    return cookie.Value, nil
}

func ValidateCookieJWT(tokenStr string) (string, error) {
    claims := &Claims{}
    tkn, err := jwt.ParseWithClaims(tokenStr, claims,
        func(t *jwt.Token) (interface{}, error){
            return jwtKey, nil
        })

    if err != nil {
        return "", err
    }

    if !tkn.Valid {
        return "", errors.New("Cookie not valid")
    }

    return claims.Username, nil
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return nil, err
    }

    return conn, nil
}