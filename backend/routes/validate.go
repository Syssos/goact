package routes

import (
  "time"
  "net/http"
  "encoding/json"

  "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Credentuals struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Validate(w http.ResponseWriter, r *http.Request) {
	// This function is an API endpoint that generates a cookie to represent 
	// the signed-in user each cookie has a refresh period of 10 minutes. 
	// If not refreshed in that time the user will be foreced to login again.

	if r.URL.Path != "/validate" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	var credentuals Credentuals
	if err := json.NewDecoder(r.Body).Decode(&credentuals); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check password
	expectedPassword, ok := users[credentuals.Username]
	if !ok || expectedPassword != credentuals.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 10)
	
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
		return
	}

	// Sets The cookie for the user
	http.SetCookie(w,
		&http.Cookie{
			Name: "token",
			Value: tokenString,
			Expires: expirationTime,
		})

}