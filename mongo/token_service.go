package mongo

import (
	"context"

	"github.com/Sippata/auth-go-service/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TokenService is an implementation of the app.TokenService interface for MongoDB
type TokenService struct {
	Instance   *DBInstance
	Collection *mongo.Collection
}

// Get refersh token from collection
func (s *TokenService) Get(token string, userID string) (string, error) {
	hash, _ := app.GenerateHash(token)

	filter := app.RefreshToken{
		UserID:    userID,
		TokenHash: hash,
	}
	projection := bson.D{{Key: "tokenhash", Value: 1}}

	var result app.RefreshToken
	err := s.Instance.WithTransaction(func() error {
		return s.Collection.FindOne(
			context.TODO(),
			filter,
			options.FindOne().SetProjection(projection),
		).Decode(&result)
	})
	if err != nil {
		return "", err
	}
	return string(result.TokenHash), err
}

// Add given refresh token in collection or return error
// if operation is failure
func (s *TokenService) Add(token string, userID string) error {
	hash, _ := app.GenerateHash(token)
	elem := app.RefreshToken{
		UserID:    userID,
		TokenHash: hash,
	}
	err := s.Instance.WithTransaction(func() error {
		_, err := s.Collection.InsertOne(context.TODO(), elem)
		return err
	})
	return err
}

// Remove refresh token from the collection using a uuid match
func (s *TokenService) Remove(token string, userID string) error {
	hash, _ := app.GenerateHash(token)

	filter := bson.D{
		{Key: "userid", Value: userID},
		{Key: "tokenhash", Value: hash},
	}

	err := s.Instance.WithTransaction(func() error {
		_, err := s.Collection.DeleteOne(context.TODO(), filter)
		return err
	})
	return err
}

// RemoveByUserID refresh token from the collection
// using a user id match
func (s *TokenService) RemoveByUserID(userID string) error {
	filter := bson.D{{Key: "userid", Value: userID}}

	err := s.Instance.WithTransaction(func() error {
		_, err := s.Collection.DeleteMany(context.TODO(), filter)
		return err
	})
	return err
}
