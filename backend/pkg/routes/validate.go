package routes

import (
  "fmt"
  "time"
  "net/http"
  "encoding/json"

  "github.com/dgrijalva/jwt-go"
)

var ValidateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/validate" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
  }
	var credentuals Credentuals
	err := json.NewDecoder(r.Body).Decode(&credentuals)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credentuals.Username]

	if !ok || expectedPassword != credentuals.Password {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Status Unauthorized")
		return
	}

	expirationTime := time.Now().Add(time.Minute * 8)

	claims := &Claims{
		Username: credentuals.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

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