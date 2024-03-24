package repository

import (
	"context"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
)

type WritingTemplateFieldRepository interface {
	AddAll(ctx context.Context, templateId int, fields []model.NewWritingTemplateField) error
	GetAll(ctx context.Context, templateId int) ([]model.WritingTemplateField, error)
	Clear(ctx context.Context) error
}
