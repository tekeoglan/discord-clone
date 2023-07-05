package model

import "context"

type LoginRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" bindign:"required"`
}

type LoginService interface {
	GetUserByEmail(c context.Context, email string) (User, error)
}
