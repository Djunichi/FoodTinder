package model

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

type SessionStatus string

const (
	SessionStatusActive     SessionStatus = "ACTIVE"
	SessionStatusTerminated SessionStatus = "TERMINATED"
	SessionStatusExpired    SessionStatus = "EXPIRED"
)

type Session struct {
	SessionId uuid.UUID     `gorm:"primary_key" db:"session_id"`
	Status    SessionStatus `db:"status"`
	CreatedAt time.Time     `db:"created_at"`
}

func (s SessionStatus) ToString() string {
	return string(s)
}

func ResolveSessionStatus(value string) (SessionStatus, error) {
	switch strings.ToUpper(value) {
	case string(SessionStatusActive):
		return SessionStatusActive, nil
	case string(SessionStatusExpired):
		return SessionStatusExpired, nil
	case string(SessionStatusTerminated):
		return SessionStatusTerminated, nil
	default:
		return "", fmt.Errorf("unknown session status %s", value)
	}
}
