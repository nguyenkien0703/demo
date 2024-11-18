package user

import (
	"context"
	"my-project/model"

	"entgo.io/ent"
)

// Repository is the interface for the user repository.

type Repository interface {
	// Create a new user
	Create(ctx context.Context, input model.CreateUserInput) (*ent.userr, error)
}

type impl struct {
	client *ent.Client
}

func (i impl) Create(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	userRecord, err := i.client.User.
		Create().
		SetDisplayName(input.DisplayName).
		SetEmail(input.Email).
		SetPassword(input.Password).
		SetRole(input.Role).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return userRecord, nil
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}

}
