package base

import "time"

type Attachment struct {
	TableName struct{} `pg:"attachment"`
	Id   int `pg:"id,pk"`
	IdTodolist int `pg:"id_todolist"`
	Url string `pg:"url"`
	Caption string `pg:"caption"`
	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
}