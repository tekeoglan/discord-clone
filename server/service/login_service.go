package service

import (
	"context"
	"github/tekeoglan/discord-clone/model"
	"time"
)

type loginService struct {
	userRepository model.UserRepository
	contextTimeout time.Duration
}

func NewLoginService(userRepository model.UserRepository, timeout time.Duration) model.LoginService {
	return &loginService{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (ls *loginService) GetUserByEmail(c context.Context, email string) (model.User, error) {
	return ls.userRepository.GetByEmail(c, email)
}
