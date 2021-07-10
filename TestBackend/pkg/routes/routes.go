package routes

import (
  "fmt"
  "time"
  "net/http"
  "encoding/json"
  
  "github.com/google/uuid"
  "github.com/gorilla/mux"
  "github.com/dgrijalva/jwt-go"
  "github.com/Syssos/goact/pkg/products"
  "github.com/Syssos/goact/pkg/websocket"
)

var TempPool = websocket.NewPool()

var jwtKey = "ThisisSuperSecret"

var users = map[string]string{
	"user1":"secret1",
	"user2":"secret1",
}

type Credentuals struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Not Implemented"))
})

var WebSock = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    serveWs(TempPool, w, r)
})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  	w.Write([]byte("API is up and running"))
})

var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	// Here we are converting the slice of products to JSON
	payload, _ := json.Marshal(products.Plist)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  var product products.Product
  vars := mux.Vars(r)
  slug := vars["slug"]

  for _, p := range products.Plist {
      if p.Slug == slug {
          product = p
      }
  }

  w.Header().Set("Content-Type", "application/json")
  if product.Slug != "" {
    payload, _ := json.Marshal(product)
    w.Write([]byte(payload))
  } else {
    http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
  }
})

var ValidateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/validate" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
	var credentuals Credentuals
	fmt.Println("Body:", r.Body)
	err := json.NewDecoder(r.Body).Decode(&credentuals)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credentuals.Username]

	if !ok || expectedPassword != credentuals.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

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

	http.SetCookie(w,
		&http.Cookie{
			Name: "token",
			Value: tokenString,
			Expires: expirationTime,
		})
})

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
    fmt.Println("WebSocket Endpoint Hit")
    cookie, err := r.Cookie("token")
    if err != nil {
    	fmt.Println(err)
    }
    fmt.Println("Cookie:",cookie)
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