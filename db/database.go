package db

import "github.com/go-pg/pg/v10"

var DB *pg.DB

func Connect() {
	DB = pg.Connect(&pg.Options{
		User:     "super_user",
		Password: "admin_123",
		Database: "task_manager_db",
		Addr:     "localhost:5432",
	})
}
