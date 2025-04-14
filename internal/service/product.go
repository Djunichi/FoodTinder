package service

import (
	"context"
	"food-tinder/internal/model"
	"github.com/google/uuid"
)

type ProductService struct {
	productRepo ProductRepo
	voteRepo    VoteRepo
}

func NewProductService(productRepository ProductRepo, voteRepository VoteRepo) *ProductService {
	return &ProductService{productRepo: productRepository, voteRepo: voteRepository}
}

func (p *ProductService) GetAllProducts(ctx context.Context) ([]model.MachineProduct, error) {
	return p.productRepo.GetAllProducts(ctx)
}

func (p *ProductService) GetUnratedProducts(ctx context.Context, sessionId uuid.UUID) ([]model.MachineProduct, error) {

	votes, err := p.voteRepo.GetVotesBySession(ctx, sessionId)
	if err != nil {
		return []model.MachineProduct{}, err
	}

	voteIds := make([]uuid.UUID, len(votes))
	for _, vote := range votes {
		voteIds = append(voteIds, vote.ProductID)
	}

	return p.productRepo.GetProductsNotInList(ctx, voteIds)
}
