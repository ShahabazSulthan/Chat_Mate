package repositories

import (
	"chatting/pkg/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (u *UserRepository) Create(user *domain.User) error {
	_, err := u.Collection.InsertOne(context.Background(), user)
	return err
}

func (u *UserRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	err := u.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	return &user, err
}
