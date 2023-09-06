package user

import (
	"context"
	"errors"
	"golang-boilerplate/ent"
	"golang-boilerplate/internal/util"
	"golang-boilerplate/model"
	"golang-boilerplate/repository"

	"go.uber.org/zap"
)

// Service is the interface for the user service.
type Service interface {
	// Create creates a new user.
	Create(ctx context.Context, input model.CreateUserInput) (*ent.User, error)
	// List lists all users by filter.
	List(ctx context.Context, filter ent.UserFilterInput) (*ent.UserConnection, error)
}

// impl is the implementation of the user service.
type impl struct {
	repoRegistry repository.RepositoryRegistry
	logger       *zap.Logger
}

// New creates a new user service.
func New(repoRegistry repository.RepositoryRegistry, logger *zap.Logger) Service {
	return &impl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

// Create creates a new user.
func (i impl) Create(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	user, err := i.repoRegistry.User().Create(ctx, input)
	if err != nil {
		i.logger.Error("create_user_has_failed", zap.Error(err))
		return nil, util.WrapGQLError(ctx, errors.Unwrap(err).Error(), "create_user_has_failed")
	}

	return user, nil
}

// List lists all users by filter.
func (i impl) List(ctx context.Context, filter ent.UserFilterInput) (*ent.UserConnection, error) {
	users, err := i.repoRegistry.User().List(ctx, filter)
	if err != nil {
		i.logger.Error("list_users_has_failed", zap.Error(err))
		return nil, util.WrapGQLError(ctx, errors.Unwrap(err).Error(), "list_user_has_failed")
	}

	return users, nil
}
