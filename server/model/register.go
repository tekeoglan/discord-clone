package model

import "context"

type Register struct {
	UserName string `form:"userName" binding:"required"`
	Email    string `form:"email" binding:"required, email"`
	Password string `form:"password" binding:"required"`
}

type RegisterResponse struct {
}

type RegisterService interface {
	Create(c context.Context, user *User) error
}
