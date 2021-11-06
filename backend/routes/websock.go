package routes

import (
    "fmt"
    "net/http"
  
    "github.com/google/uuid"
    "github.com/dgrijalva/jwt-go"
    "github.com/Syssos/goact/models/chatroom"
)

func serveWs(room *chatroom.Room, w http.ResponseWriter, r *http.Request) {
    
    // Checking if token exists from Login page
    cookie, err := r.Cookie("token")
    if err != nil {
    	if err == http.ErrNoCookie {
    		w.WriteHeader(http.StatusUnauthorized)
    		return
    	}
    	w.WriteHeader(http.StatusBadRequest)
    	return
    }

    // Getting cookie data from "token" cookie
    // Then validating if the cookie, is a valid jwt token
    tokenStr := cookie.Value
    claims := &Claims{}
    tkn, err := jwt.ParseWithClaims(tokenStr, claims,
    	func(t *jwt.Token) (interface{}, error){
    		return jwtKey, nil
    	})
    
    // check if theres an error validating jwt token.
    if err != nil {
    	if err == jwt.ErrSignatureInvalid {
    		w.WriteHeader(http.StatusUnauthorized)
    		return
    	}
    	w.WriteHeader(http.StatusBadRequest) // triggered if token="IncorrectValueForSure", as well as if token is expired
    	return
    }

    if !tkn.Valid {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    conn, err := Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &chatroom.Client{
        ID:       uuid.New(),
        UserName: claims.Username,
        Conn:     conn,
        Room:     room,
    }

    room.Register <- client
    client.Read()
}