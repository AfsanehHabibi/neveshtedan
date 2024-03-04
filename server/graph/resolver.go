package graph

import "github.com/AfsanehHabibi/neveshtedan/pkg/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	WFRepo   repository.WritingEntryFieldRepository
	WERepo   repository.WritingEntryRepository
	UserRepo repository.UserRepository
}
