package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	Products *ProductRepository
	Sessions *SessionRepository
	Votes    *VoteRepository
	logger   *zap.SugaredLogger
}

func NewRepositoryContainer(db *gorm.DB, mongo *mongo.Client, logger *zap.SugaredLogger) *Container {
	return &Container{
		Products: NewProductRepository(mongo, logger),
		Sessions: NewSessionRepository(db),
		Votes:    NewVoteRepository(db),
	}
}
