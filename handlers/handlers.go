package handlers

import (
	"Library/db"
	"Library/types"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strings"
)

type HttpHandler struct {
	db db.Db
}

func NewHttpHandler(db db.Db) *HttpHandler {
	return &HttpHandler{
		db: db,
	}
}

func (s *HttpHandler) GreetingsHandler(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/")
	if id != "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	io.WriteString(w, "Hello, world!\n")
}

func (s *HttpHandler) BooksHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		s.GetBooksHandler(w, req)
		return
	case http.MethodPost:
		s.PostBookHandler(w, req)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (s *HttpHandler) GetBooksHandler(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/books/")
	if id == "" {
		s.GetAllBooks(w, req)
	} else {
		s.GetOneBookById(w, req)
	}
}

func (s *HttpHandler) GetAllBooks(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Received GET all books request")
	books, err := s.db.ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("error occured while doing ReadAll op", err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		fmt.Println("error occured while responding to ReadAll Op", err)
		return
	}
}

func (s *HttpHandler) GetOneBookById(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/books/")
	fmt.Printf("Received GET book request for id: %s\n", id)
	book, err := s.db.ReadOneById(id)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("error occured while doing ReadbyId op", err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		fmt.Println("error occured while responding to ReadById Op", err)
		return
	}
}

func (s *HttpHandler) PostBookHandler(w http.ResponseWriter, req *http.Request) {

	var b types.Book
	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("error decoding POST request body", err)
		return
	}
	id, err := s.db.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("error persisting book", err)
		return
	}
	b.Id = id
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		w.WriteHeader(http.StatusCreated)
		fmt.Println("error encoding POST response", err)
		return
	}
}
