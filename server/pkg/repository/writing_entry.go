package repository

import "github.com/AfsanehHabibi/neveshtedan/graph/model"

type WritingEntryRepository interface {
	Add(entry model.NewWritingEntry) (int, error)
	GetById(id int) (*model.WritingEntry, error)
	GetAll() ([]model.WritingEntry, error)
}
