package inventory

import (
	"context"

	"gobase/api/internal/model"
	"gobase/api/pkg/db/pg"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	// ListProducts gets a list of products from DB
	ListProducts(context.Context, ProductsFilter) ([]model.Product, error)
	// CreateProduct saves product in DB
	CreateProduct(context.Context, model.Product) (model.Product, error)
	// UpdateProductStatus updates the product status in DB
	UpdateProductStatus(context.Context, model.Product) (model.Product, error)
	// GetActiveProductsCountFromDB gets active products count from DB
	GetActiveProductsCountFromDB(ctx context.Context) (int64, error)
}

// New returns an implementation instance satisfying Repository
func New(dbConn pg.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn pg.ContextExecutor
}
