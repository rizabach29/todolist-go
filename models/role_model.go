package models

import "time"

type Role struct {
	TableName struct{} `pg:"roles"`
	Id          int `pg:"id,pk"`
	Name        string `pg:"name"`
	Description string `pg:"description"`
	CreatedAt   time.Time `pg:"created_at"`
	UpdatedAt   time.Time `pg:"updated_at"`
}

type CreateRoleModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateRoleModel struct {
	CreateTodoModel
}