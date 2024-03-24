package logic

import (
	"context"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
)

type Neveshtedan interface {
	CreateWritingEntry(ctx context.Context, input model.NewWritingEntry) (*model.WritingEntry, error)
	CreateWritingTemplate(ctx context.Context, input model.NewWritingTemplate) (int, error)
	CreateUser(ctx context.Context, input model.NewUser) (string, error)
	Login(ctx context.Context, input model.Login) (string, error)
	RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error)

	Entries(ctx context.Context) ([]*model.WritingEntry, error)
	Templates(ctx context.Context) ([]*model.WritingTemplate, error)
	Template(ctx context.Context, id int) (*model.WritingTemplate, error)
}