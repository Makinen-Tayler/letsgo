package main

import (
	"fmt"
    "github.com/Makinen-Tayler/letsgo/handlers"
	"net/http"
    "rsc.io/quote"
)


func main() {
    fmt.Println(quote.Go())
	fmt.Println("Starting server on :8080")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.HomeHandler)
	http.ListenAndServe(":8080", nil)
}