package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"

	"github/tekeoglan/discord-clone/model"
)

const sessionPrefix = "session"

var config *sessionConfig

type sessionConfig struct {
	cookieExpr  int
	sessionExpr time.Duration
	path        string
	domain      string
	secure      bool
	httpOnly    bool
}

type sessionService struct {
	cacheRepository   model.CacheRepository
	sessionExpiration time.Duration
}

func NewSessionService(cacheRepository model.CacheRepository) model.SessionService {
	config = &sessionConfig{
		cookieExpr:  24 * 60 * 60,
		domain:      "localhost",
		path:        "/",
		secure:      false,
		httpOnly:    false,
		sessionExpr: time.Hour * 24,
	}

	if os.Getenv("ENV") == "production" {
		config.domain = "discord-clone.com"
		config.secure = true
		config.httpOnly = false
	}

	return &sessionService{
		cacheRepository:   cacheRepository,
		sessionExpiration: config.sessionExpr,
	}
}

func (ss *sessionService) CreateSession(c context.Context, userId string) (string, error) {
	uuid := uuid.New().String()
	key := fmt.Sprintf("%s:%s", sessionPrefix, uuid)

	err := ss.cacheRepository.Set(c, key, userId, ss.sessionExpiration)
	if err != nil {
		return "", err
	}

	return key, err
}

func (ss *sessionService) RetriveSession(c context.Context, sessionId string) (string, error) {
	return ss.cacheRepository.Get(c, sessionId)
}

func (ss *sessionService) RemoveSession(c context.Context, sessionId string) error {
	_, err := ss.cacheRepository.Delete(c, sessionId)
	return err
}

func (ss *sessionService) GetCokiPath() string {
	return config.path
}

func (ss *sessionService) GetCokiDomain() string {
	return config.domain
}

func (ss *sessionService) GetCokiExpr() int {
	return config.cookieExpr
}

func (ss *sessionService) GetExpr() time.Duration {
	return config.sessionExpr
}

func (ss *sessionService) IsCokiSecure() bool {
	return config.secure
}

func (ss *sessionService) IsCokiHttpOnly() bool {
	return config.httpOnly
}
