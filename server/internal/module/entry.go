package module

import (
	"context"
	"errors"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/internal/auth"
	"github.com/AfsanehHabibi/neveshtedan/pkg/util"
)

func (m NeveshtedanModule) CreateWritingEntry(ctx context.Context, input model.NewWritingEntry) (*model.WritingEntry, error) {
	userId := auth.GetUseFromContext(ctx)
	if userId == nil {
		return nil, errors.New("access denied")
	}

	id, err := m.entryRep.Add(context.Background(), input, *userId)
	if err != nil {
		return nil, err
	}

	fields := util.RemoveNilElements(input.Fields)
	err = m.entryFieldRep.AddAll(context.Background(), id, fields)
	if err != nil {
		return nil, err
	}
	oFields := make([]*model.WritingEntryField, 0, len(fields))
	for _, v := range fields {
		oFields = append(oFields, model.ConvertNewToWritingEntryField(v))
	}
	return &model.WritingEntry{ID: id, UserID: *userId, TemplateID: input.TemplateID, Fields: oFields}, nil
}

//TODO(improve performance or add filter)
func (m NeveshtedanModule) Entries(ctx context.Context) ([]*model.WritingEntry, error) {
	entries, err := m.entryRep.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var results = make([]*model.WritingEntry, 0, len(entries))
	for _, entry := range entries {
		result, err := m.Entry(ctx, entry.ID)
		if err == nil {
			results = append(results, result)
		} else {
			return nil, err
		}
	}

	return results, nil
}

func (m NeveshtedanModule) Entry(ctx context.Context, id int) (*model.WritingEntry, error) {
	entry, err := m.entryRep.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	fields, err := m.entryFieldRep.GetAll(ctx, entry.ID)
	if err != nil {
		return nil, err
	}

	entry.Fields = util.ConvertToPointerArray(fields)
	return entry, nil
}
