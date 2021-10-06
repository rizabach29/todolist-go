package base

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