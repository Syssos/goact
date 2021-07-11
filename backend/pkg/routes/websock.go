package routes

import (
  "fmt"
  "net/http"
  
  "github.com/google/uuid"
  "github.com/dgrijalva/jwt-go"
  "github.com/Syssos/goact/pkg/websocket"	
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
    
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

    	w.WriteHeader(http.StatusBadRequest)
    	return
    }

    if !tkn.Valid {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    conn, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &websocket.Client{
        ID:       uuid.New(),
        UserName: claims.Username,
        Conn:     conn,
        Pool:     pool,
    }

    pool.Register <- client
    client.Read()
}