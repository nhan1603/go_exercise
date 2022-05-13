package products

import (
	"context"
	"gobase/api/internal/model"
)

// CreateInput holds input params for creating the product
type CreateInput struct {
	Name  string
	Desc  string
	Price int64
}

// Create creates the product
func (i impl) Create(ctx context.Context, inp CreateInput) (model.Product, error) {
	product := model.Product{
		Name:        inp.Name,
		Description: inp.Desc,
		Status:      model.ProductStatusInactive,
		Price:       inp.Price,
	}

	product, err := i.repo.Inventory().CreateProduct(ctx, product)
	return product, err
}
