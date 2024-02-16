package inmemory

import (
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestWhenAddingMultipleEntriesCanRetrieveThemAllLater(t *testing.T) {
	repo := NewInMemoryWritingEntryRepository()
	userId := 9485
	templateId := 6543

	id1, err := repo.Add(model.NewWritingEntry{UserID: userId, TemplateID: templateId})
	assert.NoError(t, err)
	id2, err := repo.Add(model.NewWritingEntry{UserID: userId, TemplateID: templateId})
	assert.NoError(t, err)

	assert.NotEqual(t, id1, id2)
	entries, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, entries, 2)
}

func TestWhenAddingOneEntriesCanRetrieveItById(t *testing.T) {
	repo := NewInMemoryWritingEntryRepository()
	userId := 9485
	templateId := 6543

	id, err := repo.Add(model.NewWritingEntry{UserID: userId, TemplateID: templateId})
	assert.NoError(t, err)

	entry, err := repo.GetById(id)
	assert.NoError(t, err)
	assert.NotNil(t, entry)
	assert.Equal(t, templateId, entry.TemplateID)
	assert.Equal(t, userId, entry.UserID)
}
