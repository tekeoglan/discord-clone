package model

import "context"

type RegisterRequest struct {
	UserName string `form:"userName" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" bindign:"required"`
}

type AccountService interface {
	Login(c context.Context, email string, pass string) (User, error)
	Register(c context.Context, user *User) error
	GetByEmail(c context.Context, email string) (User, error)
	IsEmailExist(c context.Context, email string) bool
	FetchUser(c context.Context, id string) (User, error)
}
