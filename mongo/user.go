package mongo

import (
	"context"
	"example.com/demo"
	"go.mongodb.org/mongo-driver/bson"
)

type userRepository struct {
	db *db
}

func NewUserRepository(db *db) demo.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Get(ctx context.Context, ID string) (*demo.User, error) {
	collections := u.db.client.Database("demo").Collection("users")

	filter := bson.M{"id": ID}

	var user demo.User

	err := collections.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) Add(ctx context.Context, user *demo.User) error {
	collections := u.db.client.Database("demo").Collection("users")

	_, err := collections.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}
