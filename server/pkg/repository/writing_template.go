package repository

import (
	"context"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
)

type WritingTemplateRepository interface {
	Add(ctx context.Context, userId int, template model.NewWritingTemplate) (int, error)
	GetById(ctx context.Context, id int) (*model.WritingTemplate, error)
	GetAll(ctx context.Context) ([]model.WritingTemplate, error)
	Clear(ctx context.Context) error
}
