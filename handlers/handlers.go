package handlers

import (
	"Library/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func GetBooksHandler(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/books/")
	fmt.Printf("Received request for id: %s\n", id)
	if id == "1" {
		w.Header().Set("Content-type", "application/json")
		b := types.Book{Id: "1", Name: "Go for Beginners"}
		time.Sleep(25 * time.Second)
		err := json.NewEncoder(w).Encode(b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Unmarshalling failed", err)
			return
		}
	} else {
		// Not found
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Id"))
	}
}

func GreetingsHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}
