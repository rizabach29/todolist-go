package models

type RegisterModel struct {
	Fullname       string `json:"full_name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	ForgotPassword string `json:"forgot_password" binding:"required"`
	RoleId         int    `json:"role_id" binding:"required"`
}

type LoginModel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
