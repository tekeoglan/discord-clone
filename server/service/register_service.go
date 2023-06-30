package service

import (
	"context"
	"github/tekeoglan/discord-clone/model"
	"time"
)

type registerService struct {
	userRepository model.UserRepository
	contextTimeout time.Duration
}

func NewRegisterService(userRepository model.UserRepository, timeout time.Duration) model.RegisterService {
	return &registerService{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (rs *registerService) Create(c context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(c, rs.contextTimeout)
	defer cancel()
	return rs.userRepository.Create(ctx, user)
}
