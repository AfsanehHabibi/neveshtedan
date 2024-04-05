package module

import (
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestAfterCreatingEntryCanRetrieveIt(t *testing.T) {
	defer clear()

	ctx = logsBasicUserInAndFillContext()
	text := "i filled this field"
	field := model.NewWritingEntryField{Name: "basic", Type: model.FieldTypeText, Text: &text}
	result, err := module.CreateWritingEntry(ctx, model.NewWritingEntry{TemplateID: 123, Fields: []*model.NewWritingEntryField{&field}})
	assert.NoError(t, err)
	assert.NotNil(t, result)

	entry, err := module.Entry(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, entry)
	assert.Equal(t, model.WritingEntryField{Name: "basic", Value: model.TextValue{Text: text}}, *entry.Fields[0])
}

func TestAfterCreatingMultipleEntriesCanRetrieveThemAll(t *testing.T) {
	defer clear()

	ctx = logsBasicUserInAndFillContext()
	text := "i filled this field"
	field := model.NewWritingEntryField{Name: "basic", Type: model.FieldTypeText, Text: &text}
	_, err := module.CreateWritingEntry(ctx, model.NewWritingEntry{TemplateID: 123, Fields: []*model.NewWritingEntryField{&field}})
	assert.NoError(t, err)
	text2 := "i filled this field again"
	field2 := model.NewWritingEntryField{Name: "detailed", Type: model.FieldTypeText, Text: &text2}
	_, err = module.CreateWritingEntry(ctx, model.NewWritingEntry{TemplateID: 1234, Fields: []*model.NewWritingEntryField{&field2}})
	assert.NoError(t, err)

	entries, err := module.Entries(ctx)
	assert.NoError(t, err)
	assert.Len(t, entries, 2)
}
