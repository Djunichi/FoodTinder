package service

import (
	"food-tinder/internal/repository"
)

type Container struct {
	ProductService *ProductService
	SessionService *SessionService
	VoteService    *VoteService
}

func NewServiceContainer(repos *repository.Container) *Container {
	return &Container{
		ProductService: NewProductService(repos.Products, repos.Votes),
		SessionService: NewSessionService(repos.Sessions),
		VoteService:    NewVoteService(repos.Votes),
	}
}
