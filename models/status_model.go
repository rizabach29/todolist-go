package models

type CreateStatusModel struct {
	Name string `json:"name"`
}

type UpdateStatusModel struct {
	CreateStatusModel
}