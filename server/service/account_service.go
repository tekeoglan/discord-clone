package service

import (
	"context"
	"github/tekeoglan/discord-clone/model"

	"golang.org/x/crypto/bcrypt"
)

type accountService struct {
	userRepository model.UserRepository
}

func NewAccountService(userRepository model.UserRepository) model.AccountService {
	return &accountService{
		userRepository: userRepository,
	}
}

func (as *accountService) Login(c context.Context, email string, pass string) (model.User, error) {
	return as.userRepository.GetByEmail(c, email)
}

func (as *accountService) Register(c context.Context, user *model.User) error {
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(encryptPassword)

	return as.userRepository.Create(c, user)
}

func (as *accountService) GetByEmail(c context.Context, email string) (model.User, error) {
	return as.userRepository.GetByEmail(c, email)
}

func (as *accountService) IsEmailExist(c context.Context, email string) bool {
	_, err := as.userRepository.GetByEmail(c, email)
	return err == nil
}
