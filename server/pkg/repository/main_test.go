package repository_test

import (
	"context"
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres/util/test"
)

var (
	ctx               = context.Background()
	writingEntry      repository.WritingEntryRepository
	writingEntryField repository.WritingEntryFieldRepository
	userRep           repository.UserRepository
	templateRep       repository.WritingTemplateRepository
	templateFieldRep  repository.WritingTemplateFieldRepository
)

func TestMain(m *testing.M) {
	con := test.PreparePostgresForTest()
	writingEntry = postgres.NewPostgresWritingEntryRepository(con)
	writingEntryField = postgres.NewPostgresWritingEntryFieldRepository(con)
	userRep = postgres.NewUserRepository(con)
	templateRep = postgres.NewPostgresWritingTemplateRepository(con)
	templateFieldRep = postgres.NewPostgresWritingTemplateFieldRepository(con)
	m.Run()
}
