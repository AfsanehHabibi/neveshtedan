package repository_test

import (
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestWhenAddingOneTextTemplateFieldCanRetrieveItLater(t *testing.T) {
	defer templateFieldRep.Clear(ctx)

	templateId := 123
	err := templateFieldRep.AddAll(ctx, templateId, []model.NewWritingTemplateField{{Name: "name", Description: "description",
		Type: model.FieldTypeText}})
	assert.NoError(t, err)
	fields, err := templateFieldRep.GetAll(ctx, templateId)
	assert.NoError(t, err)
	assert.Len(t, fields, 1)
	assert.Equal(t, "name", fields[0].Name)
	assert.Equal(t, "description", fields[0].Description)
	assert.Equal(t, model.FieldTypeText, fields[0].Type)
}

func TestWhenAddingMultipleTemplateFieldsFromEachTypeCanRetrieveThemLater(t *testing.T) {
	defer templateFieldRep.Clear(ctx)

	entryId := 123
	fields := make([]model.NewWritingTemplateField, 0, 4)
	fields = append(fields, model.NewWritingTemplateField{Name: "text Field", Description: "text Field des", Type: model.FieldTypeText})
	fields = append(fields, model.NewWritingTemplateField{Name: "image Field", Description: "image Field des", Type: model.FieldTypeImage})
	fields = append(fields, model.NewWritingTemplateField{Name: "video Field", Description: "video Field des", Type: model.FieldTypeVideo})
	fields = append(fields, model.NewWritingTemplateField{Name: "number Field", Description: "number Field des", Type: model.FieldTypeNumber})

	err := templateFieldRep.AddAll(ctx, entryId, fields)

	assert.NoError(t, err)
	oFields, err := templateFieldRep.GetAll(ctx, entryId)
	assert.NoError(t, err)
	assert.Len(t, oFields, 4)
}
