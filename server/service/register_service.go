package service

import (
	"context"
	"github/tekeoglan/discord-clone/model"
)

type registerService struct {
	userRepository model.UserRepository
}

func NewRegisterService(userRepository model.UserRepository) model.RegisterService {
	return &registerService{
		userRepository: userRepository,
	}
}

func (rs *registerService) Create(c context.Context, user *model.User) error {
	return rs.userRepository.Create(c, user)
}

func (rs *registerService) GetUserByEmail(c context.Context, email string) (model.User, error) {
	return rs.userRepository.GetByEmail(c, email)
}
