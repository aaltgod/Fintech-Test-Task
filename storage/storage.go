package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	DSN = "host=localhost port=5432 user=postgres dbname=storage sslmode=disable"
)

type URL struct {
	Name string `json:"url"`
	ID string `json:"-,omitempty"`
	Long, Short string `json:"-,omitempty"`
}

type Storage interface {
	Create(ID, longURL, shortURL string) error
	GetShort(shortURL string) string
	GetLong(longURL string) string
}

type URLStorage struct {}

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
		`CREATE TABLE url(id VARCHAR(10), long VARCHAR(75), short VARCHAR(30));`,
	}
	for _, q := range qs {
		_, err := db.Exec(q)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (us *URLStorage) Create(ID, longURL, shortURL string) error {
	db := CreateConnection()
	defer db.Close()

	_, err := db.Exec(
		"INSERT INTO url (id, long, short) VALUES ($1, $2, $3)", ID, longURL, shortURL,
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

	return url.Short
}

func (us *URLStorage) GetLong(shortURL string) string {
	db := CreateConnection()
	defer db.Close()

	var url URL

	row := db.QueryRow("SELECT long FROM url WHERE short=$1", shortURL)
	row.Scan(&url.Long)

	return url.Long
}

