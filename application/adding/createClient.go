package adding

import (
	"context"
	"ms-client/domain/model/client"
	"ms-client/domain/repository"
)

// Service provides adding operations.
type Service interface {
	AddClient(ctx context.Context, ID, IDType, FirstName string) error
}

type service struct {
	repository repository.ClientRepository
}

// NewService creates an adding service with the necessary dependencies
func NewService(repository repository.ClientRepository) Service {
	return &service{repository}
}

// AddClient adds the given client to storage
func (s *service) AddClient(ctx context.Context, ID, IDType, FirstName string) error {
	c := client.New(ID, IDType, FirstName)
	return s.repository.Create(ctx, c)
}