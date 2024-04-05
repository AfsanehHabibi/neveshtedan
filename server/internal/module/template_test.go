package module

import (
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestAfterCreatingTemplateCanRetrieveIt(t *testing.T) {
	defer clear()

	ctx = logsBasicUserInAndFillContext()
	field := model.NewWritingTemplateField{Name: "basic", Type: model.FieldTypeText}
	result, err := module.CreateWritingTemplate(ctx, model.NewWritingTemplate{Title: "title", Description: "desc",
		Fields: []*model.NewWritingTemplateField{&field}})
	assert.NoError(t, err)
	assert.NotNil(t, result)

	template, err := module.Template(ctx, result)
	assert.NoError(t, err)
	assert.NotNil(t, template)
	assert.Equal(t, model.WritingTemplateField{Name: "basic", Type: model.FieldTypeText}, *template.Fields[0])
}

func TestAfterOneUserCreatingMultipleTemplatesCanRetrieveThemAll(t *testing.T) {
	defer clear()

	ctx = logsBasicUserInAndFillContext()
	field := model.NewWritingTemplateField{Name: "basic", Type: model.FieldTypeText}
	_, err := module.CreateWritingTemplate(ctx, model.NewWritingTemplate{Title: "title", Description: "desc",
		Fields: []*model.NewWritingTemplateField{&field}})
	assert.NoError(t, err)
	field2 := model.NewWritingTemplateField{Name: "detailed", Type: model.FieldTypeVideo}
	_, err = module.CreateWritingTemplate(ctx, model.NewWritingTemplate{Title: "title2", Description: "desc2",
		Fields: []*model.NewWritingTemplateField{&field2}})
	assert.NoError(t, err)

	templates, err := module.Templates(ctx)
	assert.NoError(t, err)
	assert.Len(t, templates, 2)
}
