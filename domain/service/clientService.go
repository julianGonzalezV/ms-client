package service

// This component is in charge of handle domain business rules
import (
	"context"
	"ms-client/domain/entity"
	"ms-client/domain/repository"
)

// ClientServiceInterface interface that establishes functions to be implemented depending on client storage
type ClientServiceInterface interface {

	// AddClient saves a given client
	AddClient(ctx context.Context, c *entity.Client) error
	// GetClient returns the client with given ID
	GetClient(ctx context.Context, ID string) (entity.Client, error)
	// SaveClient modify client with given ID and given new data
	SaveClient(ctx context.Context, ID string, c *entity.Client) error
	// GetAllClients returns all existing clients into de storage
	GetAllClients(ctx context.Context) ([]*entity.Client, error)
}

type clientService struct {
	repository repository.ClientRepository
}

// NewClientService creates the client service for business logic
func NewClientService(repository repository.ClientRepository) ClientServiceInterface {
	return &clientService{repository}
}

// AddClient adds the given client to storage
func (service *clientService) AddClient(ctx context.Context, c *entity.Client) error {
	return service.repository.Create(ctx, c)
}

// SaveClient save changes of given client to storage
func (service *clientService) SaveClient(ctx context.Context, ID string, c *entity.Client) error {
	return service.repository.Update(ID, c)
}

// GetClient searches a client given its Id, It returns multiple results client and error
func (service *clientService) GetClient(ctx context.Context, ID string) (entity.Client, error) {
	return service.repository.FetchByID(ID)
}

// GetAllClients searches all clients into the storage
func (service *clientService) GetAllClients(ctx context.Context) ([]*entity.Client, error) {
	return service.repository.Fetch()
}
