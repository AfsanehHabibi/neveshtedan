package repository_test

import (
	"context"
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	inmemory "github.com/AfsanehHabibi/neveshtedan/pkg/repository/in_memory"
)

var (
	ctx    = context.Background()
	wEImp  []repository.WritingEntryRepository
	wEFImp []repository.WritingEntryFieldRepository
)

func TestMain(m *testing.M) {
	wEFImp = append(wEFImp, inmemory.NewInMemoryWritingEntryFieldRepository())
	wEImp = append(wEImp, inmemory.NewInMemoryWritingEntryRepository())
	m.Run()
}
