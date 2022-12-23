package db

import (
	"Library/types"
	"encoding/json"
	"errors"
	"os"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("record not found")
)

type Db interface {
	ReadAll() ([]types.Book, error)
	ReadOneById(string) (*types.Book, error)
	Write(types.Book) (string, error)
	Update(types.Book) error
	Delete(string) error
}

type JsonDb struct {
	FileName string
}

func NewJsonDb(filename string) *JsonDb {
	return &JsonDb{
		FileName: filename,
	}
}

func (db *JsonDb) ReadAll() ([]types.Book, error) {
	data, err := os.ReadFile(db.FileName)
	if err != nil {
		return nil, err
	}
	var books []types.Book
	err = json.Unmarshal(data, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (db *JsonDb) ReadOneById(id string) (*types.Book, error) {
	data, err := os.ReadFile(db.FileName)
	if err != nil {
		return nil, err
	}
	var books []types.Book
	err = json.Unmarshal(data, &books)
	if err != nil {
		return nil, err
	}
	for _, book := range books {
		if book.Id == id {
			return &book, nil
		}
	}
	return nil, ErrNotFound
}

// if book already exists is not checked, new record is created every time
func (db *JsonDb) Write(b types.Book) (string, error) {
	//TODO:
	data, err := os.ReadFile(db.FileName)
	if err != nil {
		return "", err
	}
	var books []types.Book
	err = json.Unmarshal(data, &books)
	if err != nil {
		return "", err
	}
	_ = uuid.New()
	return "", nil
}

func (db *JsonDb) Update(b types.Book) error {
	//TODO:
	return nil
}

func (db *JsonDb) Delete(id string) error {
	//TODO:
	return nil
}
