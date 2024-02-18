package inmemory

import (
	"context"
	"sync"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	rModel "github.com/AfsanehHabibi/neveshtedan/pkg/repository/model"
)

type InMemoryWritingEntryFieldRepository struct {
	fields map[int][]rModel.WritingEntryField
	mu     sync.RWMutex
}

func NewInMemoryWritingEntryFieldRepository() repository.WritingEntryFieldRepository {
	return &InMemoryWritingEntryFieldRepository{
		fields: make(map[int][]rModel.WritingEntryField),
	}
}

func (repo *InMemoryWritingEntryFieldRepository) Add(ignored context.Context, entryId int, field model.NewWritingEntryField) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	newField := graphEntryFieldToRepository(entryId, field)

	repo.fields[entryId] = append(repo.fields[entryId], newField)

	return nil
}

func (repo *InMemoryWritingEntryFieldRepository) AddAll(ignored context.Context, entryId int, fields []model.NewWritingEntryField) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, field := range fields {
		newField := graphEntryFieldToRepository(entryId, field)

		repo.fields[entryId] = append(repo.fields[entryId], newField)
	}

	return nil
}

func (repo *InMemoryWritingEntryFieldRepository) GetAll(ignored context.Context, id int) ([]model.WritingEntryField, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	fields := repo.fields[id]
	converted := make([]model.WritingEntryField, 0, len(fields))
	for _, field := range fields {
		converted = append(converted, repositoryEntryFieldToGraph(field))
	}
	return converted, nil
}

func (repo *InMemoryWritingEntryFieldRepository) Clear(ctx context.Context) error {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	repo.fields = make(map[int][]rModel.WritingEntryField)
	return nil
}

func repositoryEntryFieldToGraph(input rModel.WritingEntryField) model.WritingEntryField {
	return model.WritingEntryField{
		Name:  input.Name,
		Value: input.Value,
	}
}

func graphEntryFieldToRepository(entryId int, input model.NewWritingEntryField) rModel.WritingEntryField {
	return rModel.WritingEntryField{
		Name:    input.Name,
		Value:   input.Value,
		EntryId: entryId,
	}
}
