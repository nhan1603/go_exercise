package rest

import (
	"context"
	"github.com/friendsofgo/errors"
)

// CreateUser will create a new user for the email
func (i impl) CreateUser(ctx context.Context, email string) (int, error) {
	if _, err := i.repo.System().FindUserByEmail(ctx, email); err == nil {
		return 0, errors.New("Existed email input.")
	}
	return i.repo.System().CreateUser(ctx, email)
}
