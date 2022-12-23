package main

import (
	"Library/db"
	"Library/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	db := db.NewJsonDb("data")
	h := handlers.NewHttpHandler(db)

	// Handlers
	http.HandleFunc("/", h.GreetingsHandler)
	http.HandleFunc("/books/", h.BooksHandler)

	// Start a server on port
	// port := ":5050"
	port := os.Args[1]
	fmt.Printf("Starting server on localhost%s ...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
