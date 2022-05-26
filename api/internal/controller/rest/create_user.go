package rest

import (
	"context"
	"github.com/friendsofgo/errors"
	"gobase/api/internal/repository"
)

// CreateUser will create a new user for the email
func (i impl) CreateUser(ctx context.Context, email string) (int, error) {
	newCtx := context.Background()
	err := i.repo.DoInTx(newCtx, func(ctx context.Context, repo repository.Registry) error {
		if _, err := repo.System().FindUserByEmail(ctx, email); err == nil {
			return errors.New("Existed email input.")
		}
		_, err := repo.System().CreateUser(ctx, email)
		return err
	}, nil)
	return 0, err
}
