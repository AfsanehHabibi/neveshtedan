package repository_test

import (
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestWhenAddingMultipleEntriesCanRetrieveThemAllLater(t *testing.T) {
	userId := 9485
	templateId := 6543

	id1, err := writingEntry.Add(ctx, model.NewWritingEntry{TemplateID: templateId}, userId)
	assert.NoError(t, err)
	id2, err := writingEntry.Add(ctx, model.NewWritingEntry{TemplateID: templateId}, userId)
	assert.NoError(t, err)

	assert.NotEqual(t, id1, id2)
	entries, err := writingEntry.GetAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, entries, 2)
	writingEntry.Clear(ctx)
}

func TestWhenAddingOneEntriesCanRetrieveItById(t *testing.T) {
	userId := 9485
	templateId := 6543

	id, err := writingEntry.Add(ctx, model.NewWritingEntry{TemplateID: templateId}, userId)
	assert.NoError(t, err)

	entry, err := writingEntry.GetById(ctx, id)
	assert.NoError(t, err)
	assert.NotNil(t, entry)
	assert.Equal(t, templateId, entry.TemplateID)
	assert.Equal(t, userId, entry.UserID)
	writingEntry.Clear(ctx)
}
