package postgresql

import (
	"database/sql"
	"log"
	"sync"
)

type URL struct {
	Name string `json:"url"`
}

type Storage interface {
	Create(longURL, shortURL string) error
	Get(shortURL string) error
}

type URLStorage struct {
	sync.Mutex
}

func NewURLStorage() *URLStorage {
	return &URLStorage{}
}

func CreateConnection() *sql.DB {

	db, err := sql.Open("postgres", "URL")
	if  err != nil {
		log.Fatalf("%s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("%s", err)
	}

	return db
}

func (us *URLStorage) Create(longURL, shortURL string) error {
	return nil
}

func (us *URLStorage) Get(longURL string) error {
	return nil
}

