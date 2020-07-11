package repository

import (
	"context"
	"ms-client/domain/entity"
)

// ClientRepository interface that establishes functions to be implemented depending on client storage
type ClientRepository interface {
	// Create saves a given client
	Create(ctx context.Context, c *entity.Client) error
	// Fetch return all clients saved in storage
	Fetch() ([]*entity.Client, error)
	// FetchByID returns the client with given ID
	FetchByID(ID string) (entity.Client, error)
	// Delete remove a client with given ID
	Delete(ID string) error
	// Update modify client with given ID and given new data
	Update(ID string, c *entity.Client) error
}
