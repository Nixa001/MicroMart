package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Nixa001/micromart-services-product/internal/model"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) *ProductRepository {
	return &ProductRepository{collection: collection}
}

func (r *ProductRepository) ListProducts() ([]model.Product, error) {
	ctx := context.Background()
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []model.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) CreateProduct(product *model.Product) error {
	ctx := context.Background()
	_, err := r.collection.InsertOne(ctx, product)
	return err
}

func (r *ProductRepository) GetProduct(id string) (*model.Product, error) {
	ctx := context.Background()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var product model.Product
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) UpdateProduct(product *model.Product) error {
	ctx := context.Background()
	update := bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"category":    product.Category,
			"stock":       product.Stock,
		},
	}
	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": product.ID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no product found with given id")
	}
	return nil
}

func (r *ProductRepository) DeleteProduct(id string) error {
	ctx := context.Background()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no product found with given id")
	}
	return nil
}

func (r *ProductRepository) SearchProducts(query string) ([]model.Product, error) {
	ctx := context.Background()
	filter := bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": query, "$options": "i"}},
			{"description": bson.M{"$regex": query, "$options": "i"}},
		},
	}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []model.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetProductsByCategory(category string) ([]model.Product, error) {
	ctx := context.Background()
	filter := bson.M{"category": category}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []model.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) UpdateStock(id string, quantity int) error {
	ctx := context.Background()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"stock": quantity}}
	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no product found with given id")
	}
	return nil
}
