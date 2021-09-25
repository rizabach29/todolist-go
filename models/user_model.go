package models

import "time"


type User struct {
	Id             int `sql:"id,pk"`
	Fullname       string `sql:"fullname"`
	Email          string `sql:"email,unique"`
	Password       string `sql:"password"`
	ForgotPassword string `sql:"forgot_password"`
	RoleId         int `sql:"role_id"`
	CreatedAt      time.Time `sql:"created_at"`
	UpdatedAt      time.Time `sql:"updated_at"`
}

type UserDB struct {
	tableName struct{} `sql:"users"`
	User
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
