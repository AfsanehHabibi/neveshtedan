package repository

import (
	"context"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
)

type WritingEntryRepository interface {
	Add(ctx context.Context, entry model.NewWritingEntry) (int, error)
	GetById(ctx context.Context, id int) (*model.WritingEntry, error)
	GetAll(ctx context.Context) ([]model.WritingEntry, error)
	Clear(ctx context.Context) error
}
