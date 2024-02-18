package repository_test

import (
	"fmt"
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestWhenAddingOneFieldCanRetrieveItLater(t *testing.T) {
	for _, repo := range wEFImp {
		entryId := 123
		value := "some value"
		a := &value
		err := repo.Add(ctx, entryId, model.NewWritingEntryField{Name: "name", Value: a})
		assert.NoError(t, err)
		fields, err := repo.GetAll(ctx, entryId)
		assert.NoError(t, err)
		assert.Len(t, fields, 1)
		assert.Equal(t, "name", fields[0].Name)
		assert.Equal(t, "some value", *fields[0].Value)
		repo.Clear(ctx)
	}
}

func TestWhenAddingTwoFieldsSeparatelyRetrieveBothLater(t *testing.T) {
	for _, repo := range wEFImp {
		entryId := 123
		value := "some value"
		a := &value

		err := repo.Add(ctx, entryId, model.NewWritingEntryField{Name: "name", Value: a})
		assert.NoError(t, err)
		err = repo.Add(ctx, entryId, model.NewWritingEntryField{Name: "name2", Value: a})
		assert.NoError(t, err)

		fields, err := repo.GetAll(ctx, entryId)
		assert.NoError(t, err)
		assert.Len(t, fields, 2)
		repo.Clear(ctx)
	}
}

func TestWhenAddingMultipleFieldsCanRetrieveThemLater(t *testing.T) {
	for _, repo := range wEFImp {
		entryId := 123
		fields := make([]model.NewWritingEntryField, 0, 3)
		for i := 0; i < 3; i++ {
			value := "value " + fmt.Sprintf("%d", i)
			name := "name " + fmt.Sprintf("%d", i)
			a := &value
			field := model.NewWritingEntryField{Name: name, Value: a}
			fields = append(fields, field)
		}

		err := repo.AddAll(ctx, entryId, fields)

		assert.NoError(t, err)
		oFields, err := repo.GetAll(ctx, entryId)
		assert.NoError(t, err)
		assert.Len(t, oFields, 3)
		for i := 0; i < 3; i++ {
			assert.Equal(t, fields[i].Name, oFields[i].Name)
			assert.Equal(t, fields[i].Value, oFields[i].Value)
		}
		repo.Clear(ctx)
	}
}
