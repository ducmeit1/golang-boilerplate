package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"tvf-saleshub-backend/ent"
	graphql1 "tvf-saleshub-backend/graphql"
)

// NewUser is the resolver for the newUser field.
func (r *mutationResolver) NewUser(ctx context.Context, input ent.NewUserInput) (*ent.User, error) {
	return r.client.User.Create().SetName(input.Name).Save(ctx)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return r.client.User.Query().Paginate(ctx, after, first, before, last, ent.WithUserOrder(orderBy))
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	return obj.ID.String(), nil
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

// User returns graphql1.UserResolver implementation.
func (r *Resolver) User() graphql1.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
