package base

import "time"

type Todo struct {
	TableName struct{} `pg:"todos"`
	Id   int `pg:"id,pk"`
	UserId int `pg:"user_id"`
	Title string `pg:"title"`
	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
}