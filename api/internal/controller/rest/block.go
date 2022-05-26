package rest

import (
	"context"
	"github.com/friendsofgo/errors"
	"gobase/api/internal/repository"
)

// Block creates a block relation between email1 and email2
func (i impl) Block(ctx context.Context, email1, email2 string) error {
	newCtx := context.Background()
	return i.repo.DoInTx(newCtx, func(ctx context.Context, repo repository.Registry) error {
		if email1 == email2 {
			return errors.New("Duplicate email input")
		}

		user1, err1 := repo.System().FindUserByEmail(ctx, email1)
		if err1 != nil {
			return err1
		}

		user2, err2 := repo.System().FindUserByEmail(ctx, email2)
		if err2 != nil {
			return err2
		}

		emailId1 := user1.ID
		emailId2 := user2.ID

		if err := repo.System().CheckExistedBlock(ctx, emailId1, emailId2); err != nil {
			return err
		}

		return repo.System().Block(ctx, emailId1, emailId2)
	}, nil)
}
