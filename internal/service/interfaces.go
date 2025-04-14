package service

import (
	"context"
	"food-tinder/internal/model"
	"github.com/google/uuid"
)

type SessionRepo interface {
	CreateSession(ctx context.Context, session *model.Session) (*model.Session, error)
	GetSessionById(ctx context.Context, id uuid.UUID) (*model.Session, error)
	UpdateSession(ctx context.Context, session *model.Session) error
	GetActiveSessions(ctx context.Context) ([]model.Session, error)
}

type ProductRepo interface {
	GetAllProducts(ctx context.Context) ([]model.MachineProduct, error)
	GetProductsNotInList(ctx context.Context, ids []uuid.UUID) ([]model.MachineProduct, error)
}

type VoteRepo interface {
	GetVotesBySession(ctx context.Context, sessionId uuid.UUID) ([]model.Vote, error)
	CreateVote(ctx context.Context, vote *model.Vote) error
	CreateVotes(ctx context.Context, votes []model.Vote) error
	UpdateVote(ctx context.Context, vote *model.Vote) error
	UpdateVotes(ctx context.Context, vote []model.Vote) error
	GetProductScoreStats(ctx context.Context) ([]model.VoteScoreStats, error)
}
