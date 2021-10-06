package base

import "time"

type TodoList struct {
	TableName struct{} `pg:"todolist"`
	Id   int `pg:"id,pk"`
	IdTodo int `pg:"id_todo"`
	Task string `pg:"task"`
	IdStatus int `pg:"id_status"`
	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
}