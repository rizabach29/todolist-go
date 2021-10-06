package app

import (
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/rizabach29/todolist-go/models/base"
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


func CreateSchema(db *pg.DB) {
	models := []interface{}{
		(*base.User)(nil),
		(*base.Role)(nil),
		(*base.Todo)(nil),
		(*base.Status)(nil),
		(*base.TodoList)(nil),
		(*base.Attachment)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
				IfNotExists: true,
		})
		if err != nil {
				panic(err)
		} else {
			log.Printf("Success Create table %T", model)
		}
	}
}