package db

import (
	"Library/types"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

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
	dirName string
}

func NewJsonDb(dirName string) *JsonDb {

	// Create the directory
	err := os.Mkdir(dirName, 0777)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		fmt.Printf("failed to create directory %s\n", dirName)
		return nil
	}

	// return ptr to new JsonDb
	return &JsonDb{
		dirName: dirName,
	}
}

func (db *JsonDb) ReadAll() ([]types.Book, error) {
	books := []types.Book{}

	fileNames, err := os.ReadDir(db.dirName)
	if err != nil {
		return nil, err
	}

	for _, fileName := range fileNames {
		data, err := os.ReadFile(filepath.Join(db.dirName, fileName.Name()))
		if err != nil {
			return nil, err
		}
		var book types.Book
		err = json.Unmarshal(data, &book)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (db *JsonDb) ReadOneById(id string) (*types.Book, error) {
	data, err := os.ReadFile(filepath.Join(db.dirName, id))
	if err != nil {
		return nil, err
	}

	var book types.Book
	err = json.Unmarshal(data, &book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// Write always creates a new file in the directory
func (db *JsonDb) Write(b types.Book) (string, error) {

	// Generate UUID for the new book
	recordId := uuid.NewString()
	b.Id = recordId

	// Marshall data from upstream
	data, err := json.Marshal(b)
	if err != nil {
		return "", err
	}

	// Create file with marshalled data
	err = os.WriteFile(filepath.Join(db.dirName, recordId), []byte(data), 0744)
	if err != nil {
		return "", err
	}

	return recordId, nil

}

func (db *JsonDb) Update(b types.Book) error {
	// Ensure file exists
	_, err := os.ReadFile(filepath.Join(db.dirName, b.Id))
	if err != nil {
		return nil
	}
	// Marshal data from Upstream
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	// Update file with marshalled data
	os.WriteFile(filepath.Join(db.dirName, b.Id), data, 0744)
	return nil
}

func (db *JsonDb) Delete(id string) error {
	return os.Remove(filepath.Join(db.dirName, id))
}
