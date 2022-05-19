package products

import (
	"context"
	"gobase/api/internal/model"
	"gobase/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	// List gets a list of products from DB
	List(ctx context.Context) ([]model.Product, error)
	// Get gets a single product from DB
	Get(ctx context.Context, extID string) (model.Product, error)
	// Create creates a product in DB
	Create(ctx context.Context, inp CreateInput) (model.Product, error)
	// Delete deletes a product from DB by marking the status as deleted
	Delete(ctx context.Context, extID string) error
	// GetActiveCount gets the count of active products
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) Controller {
	return impl{repo: repo}
}

type impl struct {
	repo repository.Registry
}
