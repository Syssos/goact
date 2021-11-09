package routes

import (
  "fmt"
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

var ValidateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// ValidateUser is a function used to set a cookie to represent which data is signed in, each cookie has a refresh period and will
	// last under 10 minutes if not refreshed again.

	// validate route
	if r.URL.Path != "/validate" {
		http.Error(w, "404 not found.", http.StatusNotFound)
        return
	}
	
	var credentuals Credentuals
	
	// Uptain data from frontend
	err := json.NewDecoder(r.Body).Decode(&credentuals)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check password
	expectedPassword, ok := users[credentuals.Username]
	if !ok || expectedPassword != credentuals.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create claims used to track user
	expirationTime := time.Now().Add(time.Minute * 8)
	claims := &Claims{
		Username: credentuals.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create a signed token string securing the claims authenticity
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Status Internal Server Error:", err)
		return
	}

	// Sets The cookie for the user
	http.SetCookie(w,
		&http.Cookie{
			Name: "token",
			Value: tokenString,
			Expires: expirationTime,
		})
})

func ValidateCookieFMT(r *http.Request) string {
		cookie, err := r.Cookie("token")
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
    	return ""
    }
    if !tkn.Valid {
        return ""
    }
    return claims.Username
}