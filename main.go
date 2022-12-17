package main

import (
	"Library/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// Handlers
	http.HandleFunc("/", handlers.GreetingsHandler)
	http.HandleFunc("/books/", handlers.GetBooksHandler)

	// Start a server on port
	// port := ":5050"
	port := os.Args[1]
	fmt.Printf("Starting server on localhost%s ...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
