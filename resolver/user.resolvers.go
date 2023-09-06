package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-boilerplate/ent"
	graphql1 "golang-boilerplate/graphql"
	"golang-boilerplate/model"
)

// User is the resolver for the user field.
func (r *mutationResolver) User(ctx context.Context) (*ent.UserOps, error) {
	return &ent.UserOps{}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*ent.UserQuery, error) {
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
func (r *userQueryResolver) List(ctx context.Context, obj *ent.UserQuery, filter *ent.UserFilterInput) (*ent.UserConnection, error) {
	return r.service.User().List(ctx, *filter)
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

// User returns graphql1.UserResolver implementation.
func (r *Resolver) User() graphql1.UserResolver { return &userResolver{r} }

// UserOps returns graphql1.UserOpsResolver implementation.
func (r *Resolver) UserOps() graphql1.UserOpsResolver { return &userOpsResolver{r} }

// UserQuery returns graphql1.UserQueryResolver implementation.
func (r *Resolver) UserQuery() graphql1.UserQueryResolver { return &userQueryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userOpsResolver struct{ *Resolver }
type userQueryResolver struct{ *Resolver }
