package inmemory

import (
	"fmt"
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestWhenAddingOneFieldCanRetrieveItLater(t *testing.T) {
	repo := NewInMemoryWritingEntryFieldRepository()
	entryId := 123
	value := "some value"
	a := &value
	err := repo.Add(entryId, model.NewWritingEntryField{Name: "name", Value: a})
	assert.NoError(t, err)
	fields, err := repo.GetAll(entryId)
	assert.NoError(t, err)
	assert.Len(t, fields, 1)
	assert.Equal(t, "name", fields[0].Name)
	assert.Equal(t, "some value", *fields[0].Value)
}

func TestWhenAddingTwoFieldsSeparatelyRetrieveBothLater(t *testing.T) {
	repo := NewInMemoryWritingEntryFieldRepository()
	entryId := 123
	value := "some value"
	a := &value

	err := repo.Add(entryId, model.NewWritingEntryField{Name: "name", Value: a})
	assert.NoError(t, err)
	err = repo.Add(entryId, model.NewWritingEntryField{Name: "name2", Value: a})
	assert.NoError(t, err)

	fields, err := repo.GetAll(entryId)
	assert.NoError(t, err)
	assert.Len(t, fields, 2)
}

func TestWhenAddingMultipleFieldsCanRetrieveThemLater(t *testing.T) {
	repo := NewInMemoryWritingEntryFieldRepository()
	entryId := 123
	fields := make([]model.NewWritingEntryField, 0, 3)
	for i := 0; i < 3; i++ {
		value := "value " + fmt.Sprintf("%d", i)
		name := "name " + fmt.Sprintf("%d", i)
		a := &value
		field := model.NewWritingEntryField{Name: name, Value: a}
		fields = append(fields, field)
	}

	err := repo.AddAll(entryId, fields)

	assert.NoError(t, err)
	oFields, err := repo.GetAll(entryId)
	assert.NoError(t, err)
	assert.Len(t, oFields, 3)
	for i := 0; i < 3; i++ {
		assert.Equal(t, fields[i].Name, oFields[i].Name)
		assert.Equal(t, fields[i].Value, oFields[i].Value)
	}
}
