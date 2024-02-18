package repository_test

import (
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestWhenAddingMultipleEntriesCanRetrieveThemAllLater(t *testing.T) {
	for _, repo := range wEImp {
		userId := 9485
		templateId := 6543

		id1, err := repo.Add(ctx, model.NewWritingEntry{UserID: userId, TemplateID: templateId})
		assert.NoError(t, err)
		id2, err := repo.Add(ctx, model.NewWritingEntry{UserID: userId, TemplateID: templateId})
		assert.NoError(t, err)

		assert.NotEqual(t, id1, id2)
		entries, err := repo.GetAll(ctx)
		assert.NoError(t, err)
		assert.Len(t, entries, 2)
		repo.Clear(ctx)
	}
}

func TestWhenAddingOneEntriesCanRetrieveItById(t *testing.T) {
	for _, repo := range wEImp {
		userId := 9485
		templateId := 6543

		id, err := repo.Add(ctx, model.NewWritingEntry{UserID: userId, TemplateID: templateId})
		assert.NoError(t, err)

		entry, err := repo.GetById(ctx, id)
		assert.NoError(t, err)
		assert.NotNil(t, entry)
		assert.Equal(t, templateId, entry.TemplateID)
		assert.Equal(t, userId, entry.UserID)
		repo.Clear(ctx)
	}
}
