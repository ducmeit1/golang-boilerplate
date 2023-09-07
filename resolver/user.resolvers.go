package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golang-boilerplate/ent"
	graphql1 "golang-boilerplate/graphql"
	"golang-boilerplate/model"
)

// User is the resolver for the user field.
func (r *mutationResolver) User(ctx context.Context) (*ent.UserOps, error) {
	return &ent.UserOps{}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.UserQueries, error) {
	return &ent.UserQuery{}, nil
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	return obj.ID.String(), nil
}

// Create is the resolver for the create field.
func (r *userOpsResolver) Create(ctx context.Context, obj *ent.UserOps, input model.CreateUserInput) (*ent.User, error) {
	return r.service.User().Create(ctx, input)
}

// List is the resolver for the list field.
func (r *userQueriesResolver) List(ctx context.Context, obj *model.UserQueries, filter *ent.UserFilterInput) (*ent.UserConnection, error) {
	panic(fmt.Errorf("not implemented: List - list"))
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

// User returns graphql1.UserResolver implementation.
func (r *Resolver) User() graphql1.UserResolver { return &userResolver{r} }

// UserOps returns graphql1.UserOpsResolver implementation.
func (r *Resolver) UserOps() graphql1.UserOpsResolver { return &userOpsResolver{r} }

// UserQueries returns graphql1.UserQueriesResolver implementation.
func (r *Resolver) UserQueries() graphql1.UserQueriesResolver { return &userQueriesResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userOpsResolver struct{ *Resolver }
type userQueriesResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *userQueryResolver) List(ctx context.Context, obj *ent.UserQuery, filter *ent.UserFilterInput) (*ent.UserConnection, error) {
	return r.service.User().List(ctx, *filter)
}
func (r *Resolver) UserQuery() graphql1.UserQueryResolver { return &userQueryResolver{r} }

type userQueryResolver struct{ *Resolver }
