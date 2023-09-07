package service

import (
	"context"
	"fmt"
	"github/tekeoglan/discord-clone/model"
	"time"

	"github.com/google/uuid"
)

const SESSION_PREFIX = "session"

type sessionService struct {
	cacheRepository   model.CacheRepository
	sessionExpiration time.Duration
}

func NewSessionService(cacheRepository model.CacheRepository) model.SessionService {
	return &sessionService{
		cacheRepository:   cacheRepository,
		sessionExpiration: time.Hour * 24,
	}
}

func (ss *sessionService) CreateSession(c context.Context, userId string) (string, error) {
	uuid := uuid.New().String()
	key := fmt.Sprintf("%s:%s", SESSION_PREFIX, uuid)

	err := ss.cacheRepository.Set(c, key, userId, ss.sessionExpiration)
	if err != nil {
		return "", err
	}

	return key, err
}

func (ss *sessionService) RetriveSession(c context.Context, sessionId string) (string, error) {
	return ss.cacheRepository.Get(c, sessionId)
}
