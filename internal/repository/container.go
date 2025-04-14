package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Container struct {
	Products *ProductRepository
	Sessions *SessionRepository
	Votes    *VoteRepository
}

func NewRepositoryContainer(db *gorm.DB, mongo *mongo.Client) *Container {
	return &Container{
		Products: NewProductRepository(mongo),
		Sessions: NewSessionRepository(db),
		Votes:    NewVoteRepository(db),
	}
}
