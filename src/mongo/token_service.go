package mongo

import (
	"context"

	"github.com/Sippata/medods-test/src/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TokenService is an implementation of the app.TokenService interface for MongoDB
type TokenService struct {
	DB *mongo.Database
}

// Get refersh token from collection
func (s *TokenService) Get(token string, userID string) (string, error) {
	collection := s.DB.Collection("refresh_tokens")

	filter := app.RefreshToken{
		UserID: userID,
		Token:  token,
	}
	projection := bson.D{{Key: "token", Value: 1}}

	var result app.RefreshToken
	err := collection.FindOne(
		context.TODO(),
		filter,
		options.FindOne().SetProjection(projection),
	).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.Token, err
}

// Add given refresh token in collection or return error
// if operation is failure
func (s *TokenService) Add(token string, userID string) error {
	elem := app.RefreshToken{
		UserID: userID,
		Token:  token,
	}
	collection := s.DB.Collection("refresh_tokens")
	_, err := collection.InsertOne(context.TODO(), elem)
	return err
}

// Remove refresh token from the collection using a uuid match
func (s *TokenService) Remove(token string, userID string) error {
	filter := bson.D{
		{Key: "userid", Value: userID},
		{Key: "token", Value: token},
	}

	collection := s.DB.Collection("refresh_tokens")
	_, err := collection.DeleteOne(context.TODO(), filter)
	return err
}

// RemoveByUserID refresh token from the collection
// using a user id match
func (s *TokenService) RemoveByUserID(userID string) error {
	filter := bson.D{{Key: "userid", Value: userID}}

	collection := s.DB.Collection("refersh_tokens")
	_, err := collection.DeleteMany(context.TODO(), filter)
	return err
}
