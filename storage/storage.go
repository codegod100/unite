package storage

import (
	"log"

	"github.com/fuck-capitalism/unite/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func init() {

	db, err := sqlx.Open("sqlite3", config.Config.DatabaseFile)
	DB = db
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	println("Data connection secured")
}

type Record struct {
	ID     string
	person Person
}

type Person struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func ReadPeople() ([]Person, error) {

	// var people []Person
	people := []Person{}
	err := DB.Select(&people, "SELECT * FROM people")
	if err != nil {
		return nil, err
	}
	println("got here")
	return people, nil
}

func (Record) Read(b []byte) (n int, err error) {
	return 0, nil

}
