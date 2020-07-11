package application

// Application Layer:
// - List of uses cases given by business users, be aware about having big amount
// o lines code in this file, in that scenario It is recommended to have a folder bu userCase
// - connects the domain and the interfaces layers
import (
	"context"
	"ms-client/domain/entity"
	"ms-client/domain/service"
)

// ClientAppInterface provides adding operations.
type ClientAppInterface interface {
	AddClient(ctx context.Context, ID, IDType, FirstName string) error
	SaveClient(ctx context.Context, ID, IDType, FirstName string) error
	GetClient(ctx context.Context, ID string) (entity.Client, error)
	GetAllClients(ctx context.Context) ([]*entity.Client, error)
}

type clientApp struct {
	service service.ClientServiceInterface
}

// NewClientApp creates the client application from App Layer
func NewClientApp(service service.ClientServiceInterface) ClientAppInterface {
	return &clientApp{service}
}

// AddClient adds the given client to storage
func (app *clientApp) AddClient(ctx context.Context, ID, IDType, FirstName string) error {
	c := entity.NewClient(ID, IDType, FirstName)
	return app.service.AddClient(ctx, c)
}

// SaveClient save changes of given client to storage
func (app *clientApp) SaveClient(ctx context.Context, ID, IDType, FirstName string) error {
	c := entity.NewClient(ID, IDType, FirstName)
	return app.service.SaveClient(ctx, ID, c)
}

// GetClient searches a client given its Id, It returns multiple results client and error
func (app *clientApp) GetClient(ctx context.Context, ID string) (entity.Client, error) {
	return app.service.GetClient(ctx, ID)
}

// GetAllClients searches all clients into the storage
func (app *clientApp) GetAllClients(ctx context.Context) ([]*entity.Client, error) {
	return app.service.GetAllClients(ctx)
}
