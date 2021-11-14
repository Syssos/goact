package routes

import (
    "os"
    "log"
    "os/user"
    "net/http"

    "github.com/joho/godotenv"
    "github.com/Syssos/goact/models/chatroom"
)

var jwtKey = []byte{}
var TempPool = chatroom.NewRoom()
var users = map[string]string{}

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

var ValidateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    Validate(w, r)
})

var Refresh = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    RefreshToken(w, r)
})

var WebSock = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    serveWs(TempPool, w, r)
})

func GetHomeDir() string {
    dirname, direrr := os.UserHomeDir()
    check(direrr)

    return dirname
}

func check(e error) {
    if e != nil {
        log.Println(e)
    }
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