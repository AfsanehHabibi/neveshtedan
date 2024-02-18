package repository

import (
	"context"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
)

type WritingEntryFieldRepository interface {
	Add(ctx context.Context, entryId int, field model.NewWritingEntryField) error
	AddAll(ctx context.Context, entryId int, fields []model.NewWritingEntryField) error
	GetAll(ctx context.Context, id int) ([]model.WritingEntryField, error)
	Clear(ctx context.Context) error
}
