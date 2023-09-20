package service

import (
	"context"
	"fmt"
	"github/tekeoglan/discord-clone/model"
	"os"
	"time"

	"github.com/google/uuid"
)

const SESSION_PREFIX = "session"
const COOKIE_PREFIX = "session_id"

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
		path:        "/",
		domain:      "localhost",
		secure:      false,
		httpOnly:    true,
		sessionExpr: time.Hour * 24,
	}

	if os.Getenv("ENV") == "production" {
		config.path = "/me"
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
