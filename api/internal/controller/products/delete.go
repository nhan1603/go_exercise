package products

import (
	"context"
	"errors"
	"fmt"

	pkgerrors "github.com/pkg/errors"
	"gobase/api/internal/model"
	"gobase/api/internal/repository"
	"gobase/api/internal/repository/inventory"
)

// Delete deletes a product from DB by marking the status as deleted
func (i impl) Delete(ctx context.Context, extID string) error {
	newCtx := context.Background()
	return i.repo.DoInTx(newCtx, func(ctx context.Context, repo repository.Registry) error {
		slice, err := repo.Inventory().ListProducts(ctx, inventory.ProductsFilter{ExtID: extID, WithLock: true})
		if err != nil {
			return err
		}
		if len(slice) != 1 {
			return pkgerrors.WithStack(fmt.Errorf("%w. got: %d", ErrUnexpectedProducts, len(slice)))
		}

		if slice[0].Status != model.ProductStatusInactive {
			return ErrNotActive
		}

		slice[0].Status = model.ProductStatusDeleted
		if _, err = repo.Inventory().UpdateProductStatus(ctx, slice[0]); err != nil {
			if errors.Is(err, inventory.ErrNotFound) {
				return ErrNotFound
			}
			return err
		}

		return nil
	}, nil)
}
