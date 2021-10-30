package routes

import (
  "fmt"
  "time"
  "net/http"
  
  "github.com/dgrijalva/jwt-go"
)
var Refresh = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    
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

    if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 60 * time.Second {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    expirationTime := time.Now().Add(time.Minute * 8)

    claims.ExpiresAt = expirationTime.Unix()

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Println("Status Internal Server Error:", err)
        return
    }

    http.SetCookie(w,
        &http.Cookie{
            Name: "token",
            Value: tokenString,
            Expires: expirationTime,
        })
})