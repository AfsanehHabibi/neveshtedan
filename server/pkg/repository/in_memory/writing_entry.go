package inmemory

import (
	"context"
	"sync"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	rModel "github.com/AfsanehHabibi/neveshtedan/pkg/repository/model"
)

type InMemoryWritingRepository struct {
	entries   map[int]rModel.WritingEntry
	mu        sync.RWMutex
	idCounter int
	idMu      sync.RWMutex
}

func NewInMemoryWritingEntryRepository() repository.WritingEntryRepository {
	return &InMemoryWritingRepository{
		entries:   make(map[int]rModel.WritingEntry),
		idCounter: 0,
	}
}

func (repo *InMemoryWritingRepository) Add(ignored context.Context, entry model.NewWritingEntry) (int, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	id := repo.NewId()
	rEntry := rModel.WritingEntry{Id: id, TemplateId: entry.TemplateID, UserId: entry.UserID}
	repo.entries[id] = rEntry

	return id, nil
}

func (repo *InMemoryWritingRepository) NewId() int {
	repo.idMu.Lock()
	defer repo.idMu.Unlock()

	repo.idCounter++
	return repo.idCounter - 1
}

func (repo *InMemoryWritingRepository) GetAll(ignored context.Context) ([]model.WritingEntry, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	var entries []model.WritingEntry
	for _, val := range repo.entries {
		entries = append(entries, model.WritingEntry{ID: val.Id, TemplateID: val.TemplateId, UserID: val.UserId})
	}
	return entries, nil
}

func (repo *InMemoryWritingRepository) GetById(ignored context.Context, id int) (*model.WritingEntry, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if value, is := repo.entries[id]; is {
		return &model.WritingEntry{ID: value.Id, TemplateID: value.TemplateId, UserID: value.UserId}, nil
	}
	return nil, nil
}

func (repo *InMemoryWritingRepository) Clear(ctx context.Context) error {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	repo.entries = make(map[int]rModel.WritingEntry)
	return nil
}
