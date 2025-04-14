package repository

import (
	"context"
	"food-tinder/internal/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	db *mongo.Client
}

func NewProductRepository(db *mongo.Client) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) SaveProducts(ctx context.Context, products []model.MachineProduct) error {
	models := make([]mongo.WriteModel, 0, len(products))
	collection := p.db.Database("food-tinder").Collection("products")
	for _, product := range products {
		filter := bson.M{"_id": product.ID}
		update := bson.M{"$set": product}

		upsert := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)

		models = append(models, upsert)
	}

	if len(models) == 0 {
		return nil
	}
	_, err := collection.BulkWrite(ctx, models)
	return err
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

	return products, nil
}

func (p *ProductRepository) GetProductsNotInList(ctx context.Context, ids []uuid.UUID) ([]model.MachineProduct, error) {
	collection := p.db.Database("food-tinder").Collection("products")

	// []uuid.UUID → []interface{}
	idInterfaces := make([]interface{}, len(ids))
	for i, id := range ids {
		idInterfaces[i] = id
	}

	filter := bson.M{
		"_id": bson.M{
			"$nin": idInterfaces,
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

	return products, nil
}
