package module

import (
	"context"
	"errors"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/internal/auth"
	"github.com/AfsanehHabibi/neveshtedan/pkg/util"
)

func (m NeveshtedanModule) CreateWritingTemplate(ctx context.Context, input model.NewWritingTemplate) (int, error) {
	userId := auth.GetUseFromContext(ctx)
	if userId == nil {
		return 0, errors.New("access denied")
	}

	id, err := m.templateRep.Add(context.Background(), *userId, input)
	if err != nil {
		return 0, err
	}

	fields := util.RemoveNilElements(input.Fields)
	err = m.templateFieldRep.AddAll(context.Background(), id, fields)
	if err != nil {
		return 0, err
	}

	return id, nil
}

//TODO(improve performance or add filter)
func (m NeveshtedanModule) Templates(ctx context.Context) ([]*model.WritingTemplate, error) {
	templates, err := m.templateRep.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var results = make([]*model.WritingTemplate, 0, len(templates))
	for _, template := range templates {
		result, err := m.Template(ctx, template.ID)
		if err != nil {
			results = append(results, result)
		}
	}

	return results, nil
}

func (m NeveshtedanModule) Template(ctx context.Context, id int) (*model.WritingTemplate, error) {
	template, err := m.templateRep.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	fields, err := m.templateFieldRep.GetAll(ctx, template.ID)
	if err != nil {
		return nil, err
	}

	template.Fields = util.ConvertToPointerArray(fields)
	return template, nil
}
