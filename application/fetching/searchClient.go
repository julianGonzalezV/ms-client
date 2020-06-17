package fetching

import (
	"context"
	"ms-client/domain/model/client"
	"ms-client/domain/repository"
)

// Service provides gset/search operation.
type Service interface {
	GetClient(ctx context.Context, ID string) (*client.Client, error)
	GetAllClients(ctx context.Context) ([]*client.Client, error)
}

type service struct {
	repository repository.ClientRepository
}

// NewService creates a service with the necessary dependencies
func NewService(repository repository.ClientRepository) Service {
	return &service{repository}
}

// GetClient searches a client given its Id, It returns multiple results client and error
func (s *service) GetClient(ctx context.Context, ID string) (*client.Client, error) {
	return s.repository.FetchByID(ID)
}

// GetAllClients searches all clients into the storage
func (s *service) GetAllClients(ctx context.Context) ([]*client.Client, error) {
	return s.repository.Fetch()
}
