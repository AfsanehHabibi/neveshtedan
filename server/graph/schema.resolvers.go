package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"fmt"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
)

// CreateWritingEntry is the resolver for the createWritingEntry field.
func (r *mutationResolver) CreateWritingEntry(ctx context.Context, input model.NewWritingEntry) (*model.WritingEntry, error) {
	panic(fmt.Errorf("not implemented: CreateWritingEntry - createWritingEntry"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// Entries is the resolver for the entries field.
func (r *queryResolver) Entries(ctx context.Context) ([]*model.WritingEntry, error) {
	panic(fmt.Errorf("not implemented: Entries - entries"))
}

// Templates is the resolver for the templates field.
func (r *queryResolver) Templates(ctx context.Context) ([]*model.WritingTemplate, error) {
	var links []*model.WritingTemplate
	dummyLink := model.WritingTemplate{
		Title:  "Basic",
		ID:     1234,
		Fields: []string{"apple", "banana", "orange"},
	}
	links = append(links, &dummyLink)
	return links, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
