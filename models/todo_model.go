package models

type CreateTodoModel struct {
	UserId int    `json:"user_id" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type UpdateTodoModel struct {
	Title string `json:"title" binding:"required"`
}