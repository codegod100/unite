package main

import (
	"log"

	"github.com/fuck-capitalism/unite/common"

	"github.com/fuck-capitalism/unite/config"
	"github.com/fuck-capitalism/unite/storage"
)

func main() {
	common.Log("Hello, World!")
	err := storage.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	common.Log(config.Config)
	println("Reusing connection success")
	people, err := storage.ReadPeople()
	if err != nil {
		log.Fatal(err)
	}
	common.Log(people)
}
