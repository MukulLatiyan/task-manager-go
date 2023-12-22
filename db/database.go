package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"log"
)

var DB *pg.DB

func Connect() {
	DB = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "subhanshu",
		Database: "postgres",
		Addr:     "localhost:5432",
	})

	if DB == nil {
		log.Fatal("Failed to connect to the database")
	}

	fmt.Println("Connected to the database")
}
