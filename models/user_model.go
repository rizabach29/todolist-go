package models

type MinimalUserModel struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	RoleId   int    `json:"role_id"`
}

type UpdateUserModel struct {
	Fullname string `json:"fullname"`
	RoleId   int    `json:"role_id"`
}