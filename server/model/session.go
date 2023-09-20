package model

import (
	"context"
	"time"
)

type SessionService interface {
	CreateSession(c context.Context, userId string) (string, error)
	RetriveSession(c context.Context, sessionId string) (string, error)
	GetCokiPath() string
	GetCokiDomain() string
	GetCokiExpr() int
	GetExpr() time.Duration
	IsCokiSecure() bool
	IsCokiHttpOnly() bool
}
