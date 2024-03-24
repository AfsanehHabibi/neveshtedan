package module

import (
	"github.com/AfsanehHabibi/neveshtedan/pkg/logic"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NeveshtedanModule struct {
	entryRep         repository.WritingEntryRepository
	entryFieldRep    repository.WritingEntryFieldRepository
	userRep          repository.UserRepository
	templateRep      repository.WritingTemplateRepository
	templateFieldRep repository.WritingTemplateFieldRepository
}

func NewNeveshtedanModule(dbCon *pgxpool.Pool) logic.Neveshtedan {
	entryRep := postgres.NewPostgresWritingEntryRepository(dbCon)
	entryFieldRep := postgres.NewPostgresWritingEntryFieldRepository(dbCon)
	userRep := postgres.NewUserRepository(dbCon)
	templateRep := postgres.NewPostgresWritingTemplateRepository(dbCon)
	templateFieldRep := postgres.NewPostgresWritingTemplateFieldRepository(dbCon)
	return &NeveshtedanModule{
		entryRep:         entryRep,
		entryFieldRep:    entryFieldRep,
		userRep:          userRep,
		templateRep:      templateRep,
		templateFieldRep: templateFieldRep}
}
