package models

import "time"


type User struct {
	TableName struct{} `pg:"users"`
	Id             int `pg:"id,pk"`
	Fullname       string `pg:"fullname"`
	Email          string `pg:"email,unique"`
	Password       string `pg:"password"`
	ForgotPassword string `pg:"forgot_password"`
	RoleId         int `pg:"role_id"`
	CreatedAt      time.Time `pg:"created_at"`
	UpdatedAt      time.Time `pg:"updated_at"`
}

type RegisterModel struct {
	Fullname       string `json:"full_name" binding:"required"`
	Email          string	`json:"email" binding:"required"`
	Password       string	`json:"password" binding:"required"`
	ForgotPassword string	`json:"forgot_password" binding:"required"`
	RoleId         int	`json:"role_id" binding:"required"`
}

type LoginModel struct {
	Email          string	`json:"email" binding:"required"`
	Password       string	`json:"password" binding:"required"`
}

type MinimalUserModel struct {
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	RoleId         int `json:"role_id"`
}
