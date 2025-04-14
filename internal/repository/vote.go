package repository

import (
	"context"
	"errors"
	"fmt"
	"food-tinder/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoteRepository struct {
	db *gorm.DB
}

func NewVoteRepository(db *gorm.DB) *VoteRepository {
	return &VoteRepository{db: db}
}

func (v *VoteRepository) CreateVote(ctx context.Context, vote *model.Vote) error {
	err := v.db.WithContext(ctx).Create(vote).Error
	if err != nil {
		return fmt.Errorf("[VoteRepository] failed to create vote: %w", err)
	}

	return nil
}

func (v *VoteRepository) CreateVotes(ctx context.Context, votes []model.Vote) error {
	err := v.db.WithContext(ctx).Create(votes).Error
	if err != nil {
		return fmt.Errorf("[VoteRepository] failed to create votes: %w", err)
	}
	return nil
}

func (v *VoteRepository) GetVote(ctx context.Context, id uuid.UUID) (*model.Vote, error) {
	var vote model.Vote
	err := v.db.WithContext(ctx).Model(&model.Vote{}).Where("vote_id = ?", id).First(&vote).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("[VoteRepository] failed to get vote: %w", err)
	}
	return nil, nil
}

func (v *VoteRepository) DeleteVote(ctx context.Context, id uuid.UUID) error {
	err := v.db.WithContext(ctx).Delete(&model.Vote{}, "vote_id = ?", id)
	if err != nil {
		return fmt.Errorf("[VoteRepository] failed to delete vote: %w", err)
	}
	return nil
}

func (v *VoteRepository) UpdateVote(ctx context.Context, vote *model.Vote) error {
	err := v.db.WithContext(ctx).Where("vote_id = ?", vote.VoteId).Model(&model.Vote{}).Updates(vote).Error
	if err != nil {
		return fmt.Errorf("[VoteRepository] failed to update vote: %w", err)
	}
	return nil
}

func (v *VoteRepository) UpdateVotes(ctx context.Context, vote []model.Vote) error {
	err := v.db.WithContext(ctx).Model(&model.Vote{}).Updates(vote).Error
	if err != nil {
		return fmt.Errorf("[VoteRepository] failed to update votes: %w", err)
	}
	return nil
}

func (v *VoteRepository) GetVotesBySession(ctx context.Context, sessionId uuid.UUID) ([]model.Vote, error) {
	var votes []model.Vote
	err := v.db.WithContext(ctx).Model(&model.Vote{}).Where("session_id = ?", sessionId).Find(&votes).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("[VoteRepository] failed to get votes by session: %w", err)
	}
	return votes, nil
}

func (v *VoteRepository) GetVotesByProduct(ctx context.Context, productId uuid.UUID) ([]model.Vote, error) {
	var votes []model.Vote
	err := v.db.WithContext(ctx).Model(&model.Vote{}).Where("product_id = ?", productId).Find(&votes).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("[VoteRepository] failed to get votes by product: %w", err)
	}
	return votes, nil
}

func (v *VoteRepository) GetProductScoreStats(ctx context.Context) ([]model.VoteScoreStats, error) {
	var results []model.VoteScoreStats
	err := v.db.Model(&model.Vote{}).
		Select("product_id, AVG(score) as avg_score, COUNT(*) as vote_count").
		Group("product_id").
		Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("[VoteRepository] failed to get product score stats: %w", err)
	}

	return results, nil
}
