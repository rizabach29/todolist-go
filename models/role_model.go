package models

type CreateRoleModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateRoleModel struct {
	CreateTodoModel
}