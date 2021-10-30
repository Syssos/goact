package routes

import (
  "os"
  "log"
  "net/http"

  "github.com/joho/godotenv"
  "github.com/gorilla/websocket"
  "github.com/Syssos/goact/models/chatroom"
)

var jwtKey = []byte{}
var TempPool = chatroom.NewRoom()
var users = map[string]string{}
var upgrader = websocket.Upgrader{
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool { return true },
}

func init() {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  jwtKey= []byte(os.Getenv("APP_JWT_TOKEN"))

  users = map[string]string{
    os.Getenv("TEST_USER1"): os.Getenv("TEST_PW1"),
    os.Getenv("TEST_USER2"): os.Getenv("TEST_PW2"),
  }
}

var WebSock = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  serveWs(TempPool, w, r)
})

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
      log.Println(err)
      return nil, err
  }

  return conn, nil
}