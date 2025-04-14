package service

import (
	"context"
	"food-tinder/internal/model"
	"github.com/google/uuid"
)

type SessionService struct {
	sessionRepo SessionRepo
}

func NewSessionService(sessionRepo SessionRepo) *SessionService {
	return &SessionService{
		sessionRepo: sessionRepo,
	}
}

func (s *SessionService) CreateSession(ctx context.Context) (*model.Session, error) {
	session := &model.Session{
		SessionId: uuid.New(),
		Status:    model.SessionStatusActive,
	}
	return s.sessionRepo.CreateSession(ctx, session)
}

func (s *SessionService) GetSessionById(ctx context.Context, id uuid.UUID) (*model.Session, error) {
	return s.sessionRepo.GetSessionById(ctx, id)
}

func (s *SessionService) GetActiveSessions(ctx context.Context) ([]model.Session, error) {
	return s.sessionRepo.GetActiveSessions(ctx)
}
