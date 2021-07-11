package routes

import (
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"
  "github.com/dgrijalva/jwt-go"
  "github.com/Syssos/goact/pkg/products"
  "github.com/Syssos/goact/pkg/websocket"
)

var TempPool = websocket.NewPool()

var jwtKey = []byte("ThisisSuperSecret")

var users = map[string]string{
	"user1":"secret1",
	"user2":"secret1",
	"test":"test",
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
