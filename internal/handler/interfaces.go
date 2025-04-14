package handler

import (
	"context"
	"food-tinder/internal/dto"
	"food-tinder/internal/model"
	"github.com/google/uuid"
)

type ProductSvc interface {
	GetUnratedProducts(ctx context.Context, sessionId uuid.UUID) ([]model.MachineProduct, error)
	GetAllProducts(ctx context.Context) ([]model.MachineProduct, error)
}

type SessionSvc interface {
	CreateSession(ctx context.Context) (*model.Session, error)
	GetSessionById(ctx context.Context, id uuid.UUID) (*model.Session, error)
	GetActiveSessions(ctx context.Context) ([]model.Session, error)
}

type VoteSvc interface {
	CreateVote(ctx context.Context, req *dto.CreateVoteReq) error
	CreateVotes(ctx context.Context, req *dto.CreateVotesReq) error
	UpdateVote(ctx context.Context, req *dto.UpdateVoteReq) error
	UpdateVotes(ctx context.Context, req *dto.UpdateVotesReq) error
	GetVotesBySession(ctx context.Context, sessionID uuid.UUID) ([]model.Vote, error)
	GetAggregatedScoresByAllSessions(ctx context.Context) ([]model.VoteScoreStats, error)
}
