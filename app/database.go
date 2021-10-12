package app

import (
	"log"
	"os"
	"sync"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/rizabach29/todolist-go/models/base"
)


func CreateSchema() {
	models := []interface{}{
		(*base.User)(nil),
		(*base.Role)(nil),
		(*base.Todo)(nil),
		(*base.Status)(nil),
		(*base.TodoList)(nil),
		(*base.Attachment)(nil),
	}

	for _, model := range models {
		err := GetDatabase().Model(model).CreateTable(&orm.CreateTableOptions{
				IfNotExists: true,
		})
		if err != nil {
			panic(err)
		} else {
			log.Printf("Success Create table %T", model)
		}
	}
}


// Singleton DB
var lock = &sync.Mutex{}

var database *pg.DB

func GetDatabase() *pg.DB {
	if database == nil {
		lock.Lock()
		defer lock.Unlock()

		if database == nil {
			opts := &pg.Options{
				User: GetEnv("DB_USERNAME"),
				Password: GetEnv("DB_PASSWORD"),
				Addr: GetEnv("DB_HOST") + ":" + GetEnv("DB_PORT"),
				Database: GetEnv("DB_NAME"),
			}
		
			database = pg.Connect(opts)
			if database == nil {
				log.Printf("Failed connect to database")
				os.Exit(0)
			}
		
			log.Printf("Connected to database")
		}
	}

	return database
}