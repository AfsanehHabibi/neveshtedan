package repository_test

import (
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestWhenAddingOneTextFieldCanRetrieveItLater(t *testing.T) {
	defer writingEntryField.Clear(ctx)

	entryId := 123
	value := "some value"
	a := &value
	err := writingEntryField.Add(ctx, entryId, model.NewWritingEntryField{Name: "name", Text: a, Type: model.FieldTypeText})
	assert.NoError(t, err)
	fields, err := writingEntryField.GetAll(ctx, entryId)
	assert.NoError(t, err)
	assert.Len(t, fields, 1)
	assert.Equal(t, "name", fields[0].Name)
	assert.Equal(t, model.TextValue{Text: "some value"}, fields[0].Value.(model.TextValue))
}

func TestWhenAddingMultipleFieldsFromEachTypeCanRetrieveThemLater(t *testing.T) {
	defer writingEntryField.Clear(ctx)

	entryId := 123
	fields := make([]model.NewWritingEntryField, 0, 4)
	var text = "text value"
	textField := model.NewWritingEntryField{Name: "text Field", Text: &text, Type: model.FieldTypeText}
	fields = append(fields, textField)
	var imgUrl = "image value"
	imgField := model.NewWritingEntryField{Name: "image Field", URL: &imgUrl, Type: model.FieldTypeImage}
	fields = append(fields, imgField)
	var videoUrl = "video value"
	videoField := model.NewWritingEntryField{Name: "video Field", URL: &videoUrl, Type: model.FieldTypeVideo}
	fields = append(fields, videoField)
	var number = 56.78
	numberField := model.NewWritingEntryField{Name: "number Field", Number: &number, Type: model.FieldTypeNumber}
	fields = append(fields, numberField)

	err := writingEntryField.AddAll(ctx, entryId, fields)

	assert.NoError(t, err)
	oFields, err := writingEntryField.GetAll(ctx, entryId)
	assert.NoError(t, err)
	assert.Len(t, oFields, 4)
}
