package products

import (
	"context"
	"fmt"

	pkgerrors "github.com/pkg/errors"
	"gobase/api/internal/model"
	"gobase/api/internal/repository/inventory"
)

// Get gets a single product from DB
func (i impl) Get(ctx context.Context, extID string) (model.Product, error) {
	slice, err := i.repo.Inventory().ListProducts(ctx, inventory.ProductsFilter{ExtID: extID})
	if err != nil {
		return model.Product{}, err
	}

	if len(slice) != 1 {
		return model.Product{}, pkgerrors.WithStack(fmt.Errorf("%w. got: %d", ErrUnexpectedProducts, len(slice)))
	}

	return slice[0], nil
}
