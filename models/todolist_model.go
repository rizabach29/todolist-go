package models

type CreateTodolistModel struct {
	IdTodo   int    `json:"id_todo"`
	Task     string `json:"task"`
	IdStatus int    `json:"id_status"`
}

type UpdateTodolistModel struct {
	Task     string `json:"task"`
	IdStatus int    `json:"id_status"`
}