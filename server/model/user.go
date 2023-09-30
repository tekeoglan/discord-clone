package model

import "context"

const CollectionUser = "users"

type User struct {
	BaseModel `bson:",inline"`
	UserName  string `bson:"userName"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
	Image     string `bson:"image, omitempty"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	GetByID(c context.Context, id string) (User, error)
	GetByEmail(c context.Context, email string) (User, error)
}
