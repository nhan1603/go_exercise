package rest

import (
	"context"
	"github.com/friendsofgo/errors"
	"gobase/api/internal/repository"
)

// Subscribe will create a subscription for of the second email for the first email
func (i impl) Subscribe(ctx context.Context, email1, email2 string) error {
	newCtx := context.Background()
	return i.repo.DoInTx(newCtx, func(ctx context.Context, repo repository.Registry) error {
		if email1 == email2 {
			return errors.New("Duplicate email input")
		}

		user1, err1 := i.repo.System().FindUserByEmail(ctx, email1)
		if err1 != nil {
			return err1
		}

		user2, err2 := i.repo.System().FindUserByEmail(ctx, email2)
		if err2 != nil {
			return err2
		}

		emailId1 := user1.ID
		emailId2 := user2.ID

		if err := i.repo.System().CheckExistedSubscribe(ctx, emailId1, emailId2); err != nil {
			return err
		}

		return i.repo.System().Subscribe(ctx, emailId1, emailId2)
	}, nil)
}
