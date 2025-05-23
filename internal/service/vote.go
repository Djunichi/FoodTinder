package service

import (
	"context"
	"food-tinder/internal/dto"
	"food-tinder/internal/model"
	"github.com/google/uuid"
)

type VoteService struct {
	voteRepo VoteRepo
}

func NewVoteService(voteRepo VoteRepo) *VoteService {
	return &VoteService{
		voteRepo: voteRepo,
	}
}

func (v *VoteService) CreateVote(ctx context.Context, req *dto.CreateVoteReq) error {

	sessionId, err := uuid.Parse(req.SessionID)
	if err != nil {
		return err
	}
	productId, err := uuid.Parse(req.Vote.ProductID)
	if err != nil {
		return err
	}

	vote := model.Vote{
		VoteId:    uuid.New(),
		SessionID: sessionId,
		ProductID: productId,
		Score:     req.Vote.Score,
	}

	return v.voteRepo.CreateVote(ctx, &vote)
}

func (v *VoteService) CreateVotes(ctx context.Context, req *dto.CreateVotesReq) error {
	votes := make([]model.Vote, len(req.Votes))

	sessionId, err := uuid.Parse(req.SessionID)
	if err != nil {
		return err
	}

	for _, vote := range req.Votes {
		productId, err := uuid.Parse(vote.ProductID)
		if err != nil {
			return err
		}
		votes = append(votes, model.Vote{
			VoteId:    uuid.New(),
			SessionID: sessionId,
			ProductID: productId,
			Score:     vote.Score,
		})
	}

	return v.voteRepo.CreateVotes(ctx, votes)
}

func (v *VoteService) UpdateVote(ctx context.Context, req *dto.UpdateVoteReq) error {

	sessionId, err := uuid.Parse(req.SessionID)
	if err != nil {
		return err
	}
	productId, err := uuid.Parse(req.Vote.ProductID)
	if err != nil {
		return err
	}
	voteId, err := uuid.Parse(req.Vote.VoteId)
	if err != nil {
		return err
	}

	vote := model.Vote{
		VoteId:    voteId,
		SessionID: sessionId,
		ProductID: productId,
		Score:     req.Vote.Score,
	}

	return v.voteRepo.UpdateVote(ctx, &vote)
}

func (v *VoteService) UpdateVotes(ctx context.Context, req *dto.UpdateVotesReq) error {
	votes := make([]model.Vote, len(req.Votes))

	sessionId, err := uuid.Parse(req.SessionID)
	if err != nil {
		return err
	}

	for _, vote := range req.Votes {
		voteId, err := uuid.Parse(vote.VoteId)
		if err != nil {
			return err
		}

		productId, err := uuid.Parse(vote.ProductID)
		if err != nil {
			return err
		}

		votes = append(votes, model.Vote{
			VoteId:    voteId,
			SessionID: sessionId,
			ProductID: productId,
			Score:     vote.Score,
		})
	}

	return v.voteRepo.UpdateVotes(ctx, votes)
}

func (v *VoteService) GetVotesBySession(ctx context.Context, sessionID uuid.UUID) ([]model.Vote, error) {
	return v.voteRepo.GetVotesBySession(ctx, sessionID)
}

func (v *VoteService) GetAggregatedScoresByAllSessions(ctx context.Context) ([]model.VoteScoreStats, error) {
	return v.voteRepo.GetProductScoreStats(ctx)
}
