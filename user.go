package demo

import (
	"context"
	"net/mail"
)

type User struct {
	ID       int        `json:"ID"`
	Name     string     `json:"Name"`
	Email    string     `json:"Email"`
	Birthday CustomTime `json:"Birthday"`
}

type UserRepository interface {
	Get(ctx context.Context, ID string) (*User, error)
	Add(ctx context.Context, user *User) error
}

func (u *User) ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
