package model

import "context"

type User struct {
	BaseModel
	UserName string `bson:"userName"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Image    string `bson:"image, omitempty"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	GetByID(c context.Context, id string) (User, error)
}
