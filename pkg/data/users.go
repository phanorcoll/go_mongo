// Package data provides data  
// this package will handle the direct interaction with the data store and any
// associated logic to process responses created from the mongo package.
package data

import (
	"context"

	"github.com/phanorcoll/go_mongo/pkg/config"
	"github.com/phanorcoll/go_mongo/pkg/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// IUserProvider interface    the data package must satisfy this contract.
type IUserProvider interface {
	CreateAccount(user *domain.User) error
	UserNameExists(username string) (bool, error)
	FindByUsername(username string) (*domain.User, error)
}

type UserProvider struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

// NewUserProvider function  
func NewUserProvider(cfg *config.Settings, mongo *mongo.Client) IUserProvider {
	userCollection := mongo.Database(cfg.DbName).Collection("users")
	return &UserProvider{
		userCollection: userCollection,
		ctx:            context.TODO(),
	}
}

func (u UserProvider) CreateAccount(user *domain.User) error {
	_, err := u.userCollection.InsertOne(u.ctx, user)
	if err != nil {
		return errors.Wrap(err, "Error inserting user")
	}
	return nil
}

func (u UserProvider) UserNameExists(username string) (bool, error) {
	var userFound *domain.User
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	if err := u.userCollection.FindOne(u.ctx, filter).Decode(&userFound); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrap(err, "Error finding by username")
	}
	return true, nil
}

func (u UserProvider) FindByUsername(username string) (*domain.User, error) {
	var userFound domain.User
	filter := bson.D{primitive.E{Key: "useername", Value: username}}
	if err := u.userCollection.FindOne(u.ctx, filter).Decode(&userFound); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error finding my username")
	}

	return &userFound, nil
}
