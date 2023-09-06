package repository

import (
	"golang-boilerplate/ent"
	"golang-boilerplate/repository/user"
)

// RepositoryRegistry is the interface for the repository registry.
type RepositoryRegistry interface {
	User() user.Repository
}

// impl is the implementation of the repository registry.
type impl struct {
	user user.Repository
}

// New creates a new repository registry.
func New(client *ent.Client) RepositoryRegistry {
	return &impl{
		user: user.New(client),
	}
}

// User returns the user repository.
func (i impl) User() user.Repository {
	return i.user
}
