package model

import (
	"github.com/google/uuid"
	"time"
)

type Vote struct {
	VoteId    uuid.UUID `gorm:"primary_key" db:"vote_id"`
	SessionID uuid.UUID `db:"session_id"`
	ProductID uuid.UUID `db:"product_id"`
	Score     int32     `db:"score"`
	CreatedAt time.Time `db:"created_at"`
}

type VoteScoreStats struct {
	ProductID uuid.UUID `json:"product_id"`
	AvgScore  float32   `json:"avg_score"`
	VoteCount int64     `json:"vote_count"`
}
