package main

import (
    "fmt"
    "net/http"

    "github.com/rs/cors"
    "github.com/gorilla/mux"
    "github.com/Syssos/goact/pkg/routes"
)

func main() {
    r := mux.NewRouter()
    go routes.TempPool.Start()

    // We will setup our server so we can serve static assest like images, css from the /static/{file} route
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

    r.Handle("/", http.FileServer(http.Dir("./views/")))

    r.Handle("/status", routes.StatusHandler).Methods("GET")
    r.Handle("/products", routes.ProductsHandler).Methods("GET")
    r.Handle("/refresh", routes.Refresh).Methods("GET")
    r.Handle("/validate", routes.ValidateUser).Methods("POST")
    r.Handle("/ws", routes.WebSock).Methods("GET")
    r.Handle("/products/{slug}/feedback", routes.AddFeedbackHandler).Methods("POST")

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"},
        AllowCredentials: true,
    })

    handler := c.Handler(r)

    fmt.Println("Starting server on localhost:8080")
    http.ListenAndServe(":8080", handler)
}
