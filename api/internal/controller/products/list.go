package products

import (
	"context"

	"gobase/api/internal/model"
	"gobase/api/internal/repository/inventory"
)

// List gets a list of products from DB
func (i impl) List(ctx context.Context) ([]model.Product, error) {
	// TODO: Add pagination example
	// TODO: Add filter example

	return i.repo.Inventory().ListProducts(ctx, inventory.ProductsFilter{
		Status: []model.ProductStatus{model.ProductStatusActive, model.ProductStatusInactive},
	})
}
