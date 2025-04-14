package repository

import (
	"context"
	"errors"
	"fmt"
	"food-tinder/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (s *SessionRepository) CreateSession(ctx context.Context, session *model.Session) (*model.Session, error) {
	err := s.db.WithContext(ctx).Create(session).Error
	if err != nil {
		return nil, fmt.Errorf("[SessionRepository] failed to create session: %w", err)
	}

	return session, nil
}

func (s *SessionRepository) GetSessionById(ctx context.Context, id uuid.UUID) (*model.Session, error) {
	err := s.db.WithContext(ctx).Model(&model.Session{}).Where("session_id = ?", id).Find(&model.Session{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("[SessionRepository] failed to get session: %w", err)
	}
	return nil, nil
}

func (s *SessionRepository) GetActiveSessions(ctx context.Context) ([]model.Session, error) {
	var sessions []model.Session
	err := s.db.WithContext(ctx).Model(&model.Session{}).Where("status = ?", model.SessionStatusActive).Error
	if err != nil {
		return nil, fmt.Errorf("[SessionRepository] failed to get active sessions: %w", err)
	}
	return sessions, nil
}

func (s *SessionRepository) DeleteSession(ctx context.Context, id uuid.UUID) error {
	err := s.db.WithContext(ctx).Delete(&model.Session{}, "session_id = ?", id)
	if err != nil {
		return fmt.Errorf("[SessionRepository] failed to delete session: %w", err)
	}
	return nil
}

func (s *SessionRepository) UpdateSession(ctx context.Context, session *model.Session) error {
	err := s.db.WithContext(ctx).Where("session_id = ?", session.SessionId).Model(&model.Vote{}).Updates(session).Error
	if err != nil {
		return fmt.Errorf("[SessionRepository] failed to update session: %w", err)
	}
	return nil
}
