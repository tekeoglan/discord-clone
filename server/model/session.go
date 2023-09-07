package model

import (
	"context"
)

type SessionService interface {
	CreateSession(c context.Context, userId string) (string, error)
	RetriveSession(c context.Context, sessionId string) (string, error)
}
