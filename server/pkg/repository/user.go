package repository

import (
	"context"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
)

type UserRepository interface {
	Add(ctx context.Context, user model.NewUser) (int, error)
	GetById(ctx context.Context, id int) (*model.User, error)
	GetIdIfExists(ctx context.Context, username string, password string) (*int, error)
	Clear(ctx context.Context) error
}
