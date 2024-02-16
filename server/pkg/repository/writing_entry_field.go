package repository

import "github.com/AfsanehHabibi/neveshtedan/graph/model"

type WritingEntryFieldRepository interface {
	Add(entryId int, field model.NewWritingEntryField) error
	AddAll(entryId int, fields []model.NewWritingEntryField) error
	GetAll(id int) ([]model.WritingEntryField, error)
}
