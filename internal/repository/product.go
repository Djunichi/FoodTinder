package repository

import (
	"context"
	"food-tinder/internal/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type ProductRepository struct {
	db  *mongo.Client
	log *zap.SugaredLogger
}

func NewProductRepository(db *mongo.Client, logger *zap.SugaredLogger) *ProductRepository {
	return &ProductRepository{db: db, log: logger}
}

func (p *ProductRepository) SaveProducts(ctx context.Context, products []model.MachineProduct) error {
	if len(products) == 0 {
		p.log.Warnf("[ProductRepository] No products to save")
		return nil
	}

	models := make([]mongo.WriteModel, 0, len(products))
	collection := p.db.Database("food-tinder").Collection("products")
	for _, product := range products {
		product.MongoId = product.ID.String()
		filter := bson.M{"_id": product.MongoId}
		update := bson.M{"$set": product}

		upsert := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)

		models = append(models, upsert)
	}

	result, err := collection.BulkWrite(ctx, models)
	if err != nil {
		if bwe, ok := err.(mongo.BulkWriteException); ok {
			for _, writeErr := range bwe.WriteErrors {
				p.log.Errorf("[ProductRepository] Bulk write error [%d]: %s", writeErr.Index, writeErr.Message)
			}
		}
		p.log.Errorf("[ProductRepository] BulkWrite failed: %v", err)
		return err
	}

	p.log.Infof("[ProductRepository] SaveProducts: Matched: %d, Modified: %d, Upserted: %d",
		result.MatchedCount, result.ModifiedCount, result.UpsertedCount)

	return nil
}

func (p *ProductRepository) GetAllProducts(ctx context.Context) ([]model.MachineProduct, error) {
	collection := p.db.Database("food-tinder").Collection("products")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []model.MachineProduct
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	for i := range products {
		if products[i].MongoId != "" {
			id, err := uuid.Parse(products[i].MongoId)
			if err != nil {
				p.log.Warnf("[ProductRepository] Invalid UUID in MongoID: %s", products[i].MongoId)
				continue
			}
			products[i].ID = id
		}
	}

	return products, nil
}

func (p *ProductRepository) GetProductsNotInList(ctx context.Context, ids []uuid.UUID) ([]model.MachineProduct, error) {
	collection := p.db.Database("food-tinder").Collection("products")

	// UUID → string
	excludedIDs := make([]interface{}, len(ids))
	for i, id := range ids {
		excludedIDs[i] = id.String()
	}

	filter := bson.M{
		"_id": bson.M{
			"$nin": excludedIDs,
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []model.MachineProduct
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	for i := range products {
		if products[i].MongoId != "" {
			id, err := uuid.Parse(products[i].MongoId)
			if err != nil {
				p.log.Warnf("[ProductRepository] Invalid UUID in MongoID: %s", products[i].MongoId)
				continue
			}
			products[i].ID = id
		}
	}

	return products, nil
}
