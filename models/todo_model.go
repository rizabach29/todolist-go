package models

import "time"

type Todo struct {
	Id   int
	UserId int
	Title string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RequestTodoModel struct {
	UserId int `json:"user_id" binding:"required"`
	Title string `json:"title" binding:"required"`
}