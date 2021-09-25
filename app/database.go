package app

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

func InitDB() *pg.DB {
	opts := &pg.Options{
		User: GetEnv("DB_USERNAME"),
		Password: GetEnv("DB_PASSWORD"),
		Addr: GetEnv("DB_HOST") + ":" + GetEnv("DB_PORT"),
		Database: GetEnv("DB_NAME"),
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed connect to database")
		os.Exit(0)
	}

	log.Printf("Connected to database")
	return db
}