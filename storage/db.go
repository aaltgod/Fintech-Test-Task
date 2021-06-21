package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var (
	DSN = "host=localhost port=5432 user=postgres password=pwd dbname=storage sslmode=disable"
)

type URL struct {
	Name string `json:"url"`
	ID int `json:"-,omitempty"`
	Long, Short string `json:"-,omitempty"`
}

type Storage interface {
	Create(longURL, shortURL string) error
	GetShort(shortURL string) string
	GetLong(longURL string) string
}

type URLStorage struct {
	sync.Mutex
}

func NewURLStorage() *URLStorage {
	return &URLStorage{}
}

func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", DSN)
	if  err != nil {
		log.Fatalf("[Open] %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("[Ping] %s", err)
	}

	return db
}

func PrepareStorage(db *sql.DB) {
	qs := []string{
		`DROP TABLE IF EXISTS url;`,
		`CREATE TABLE url(long varchar(75), short varchar(15));`,
	}
	for _, q := range qs {
		_, err := db.Exec(q)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (us *URLStorage) Create(longURL, shortURL string) error {
	db := CreateConnection()
	defer db.Close()

	_, err := db.Exec(
		"INSERT INTO url (long, short) VALUES ($1, $2)", longURL, shortURL,
		)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (us *URLStorage) GetShort(longURL string) string {
	db := CreateConnection()
	defer db.Close()

	var url URL

	row := db.QueryRow("SELECT short FROM url WHERE long=$1", longURL)
	row.Scan(&url.Short)

	log.Println("Short", row)

	return url.Short
}

func (us *URLStorage) GetLong(shortURL string) string {
	db := CreateConnection()
	defer db.Close()

	var url URL

	row := db.QueryRow("SELECT long FROM url WHERE short=$1", shortURL)
	row.Scan(&url.Long)

	log.Println("Long", row)

	return url.Long
}

