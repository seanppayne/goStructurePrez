package mocks

import (
	"context"
	"example.com/demo"
)

type UserRepository struct {
	GetUserFunc func(ctx context.Context, ID string) (*demo.User, error)
	AddUserFunc func(ctx context.Context, user *demo.User) error
}

func (u *UserRepository) Get(ctx context.Context, ID string) (*demo.User, error) {
	return u.GetUserFunc(ctx, ID)
}

func (u *UserRepository) Add(ctx context.Context, user *demo.User) error {
	return u.AddUserFunc(ctx, user)
}
