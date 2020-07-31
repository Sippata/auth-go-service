package mongo

import (
	"context"
	"medods-test/app"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TokenService is an implementation of the app.TokenService interface for MongoDB
type TokenService struct {
	DB *mongo.Database
}

// Get app.RefreshTokenDetail from collection of "tokens"
func (s *TokenService) Get(uuid string) (*app.RefreshTokenDetail, error) {
	var result app.RefreshTokenDetail
	filter := bson.D{{Key: "UUID", Value: uuid}}

	collection := s.DB.Collection("tokens")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

// GetByAccessUUID return app.RefreshTokenDetail from collection of "tokens" by access uuid
func (s *TokenService) GetByAccessUUID(uuid string) (*app.RefreshTokenDetail, error) {
	var result app.RefreshTokenDetail
	filter := bson.D{{Key: "AccessUUID", Value: uuid}}

	collection := s.DB.Collection("tokens")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

// Add given app.RefreshTokenDetail in collection "tokens" or return error
// if operation is failure
func (s *TokenService) Add(token *app.RefreshTokenDetail) error {
	collection := s.DB.Collection("tokens")
	_, err := collection.InsertOne(context.TODO(), token)
	return err
}

// Remove app.RefreshTokenDetail from the "tokens" collection using a uuid match
func (s *TokenService) Remove(uuid string) error {
	filter := bson.D{{Key: "UUID", Value: uuid}}

	collection := s.DB.Collection("tokens")
	_, err := collection.DeleteOne(context.TODO(), filter)
	return err
}

// RemoveByUserID app.RefreshTokenDetail from the "token" collection
// using a user id match
func (s *TokenService) RemoveByUserID(userID string) error {
	filter := bson.D{{Key: "UserID", Value: userID}}

	collection := s.DB.Collection("tokens")
	_, err := collection.DeleteMany(context.TODO(), filter)
	return err
}
