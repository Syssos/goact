package routes

import (
    "os"
    "os/user"
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
    var filenameForT string
    dirname := GetHomeDir()

    if travis_check() {
        filenameForT = "/home/travis/gopath/src/github.com/Syssos/goact/backend/.env"
    } else {
        filenameForT = dirname + "/go/goact/backend/.env"
    }
    err := godotenv.Load(filenameForT)
    if err != nil {
        log.Println(err)
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

func GetHomeDir() string {
    dirname, direrr := os.UserHomeDir()
    check(direrr)

    return dirname
}

func check(e error) {
    log.Println(e)
}

func travis_check() bool {
    user, err := user.Current()
    if err != nil {
        return false
    }

    if user.Username == "travis" {
        return true
    }

    return false
  }