package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    username := os.Getenv("USERNAME")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello %s, you've requested: %s\n", username, r.URL.Path)
    })

    http.ListenAndServe(":" + port, nil)
}