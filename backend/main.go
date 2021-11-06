package main

import (
    "fmt"
    "net/http"

    "github.com/rs/cors"
    "github.com/gorilla/mux"
    "github.com/Syssos/goact/routes"
)

func main() {
    go routes.TempPool.Start()
    r := CreateRouter()
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"},
        AllowCredentials: true,
    })
    handler := c.Handler(r)

    fmt.Println("Starting server on localhost:8080")
    http.ListenAndServe(":8080", handler)
}


func CreateRouter() *mux.Router {
    r := mux.NewRouter()
    
    // Static assest route handling for images, css, and static js from <url.com>/static/{file}
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

    r.Handle("/", http.FileServer(http.Dir("./views/")))
    r.Handle("/refresh", routes.Refresh).Methods("GET")
    r.Handle("/validate", routes.ValidateUser).Methods("POST")
    r.Handle("/ws", routes.WebSock).Methods("GET")

    return r
}