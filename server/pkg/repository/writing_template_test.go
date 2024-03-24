package repository_test

import (
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestWhenAddingMultipleTemplatesCanRetrieveThemAllLater(t *testing.T) {
	defer templateRep.Clear(ctx)

	userId := 9485

	id1, err := templateRep.Add(ctx, userId, model.NewWritingTemplate{Title: "template title 1", Description: "template description 1"})
	assert.NoError(t, err)
	id2, err := templateRep.Add(ctx, userId, model.NewWritingTemplate{Title: "template title 2", Description: "template description 2"})
	assert.NoError(t, err)

	assert.NotEqual(t, id1, id2)
	templates, err := templateRep.GetAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, templates, 2)
}

func TestWhenAddingOneTemplateCanRetrieveItById(t *testing.T) {
	defer templateRep.Clear(ctx)

	userId := 9485

	title := "template title 1"
	description := "template description 1"
	id, err := templateRep.Add(ctx, userId, model.NewWritingTemplate{Title: title, Description: description})
	assert.NoError(t, err)

	template, err := templateRep.GetById(ctx, id)
	assert.NoError(t, err)
	assert.NotNil(t, template)
	assert.Equal(t, id, template.ID)
	assert.Equal(t, userId, template.UserID)
	assert.Equal(t, title, template.Title)
	assert.Equal(t, description, template.Description)
}
