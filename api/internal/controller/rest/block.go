package rest

import (
	"context"
	"github.com/friendsofgo/errors"
	"gobase/api/internal/model"
	"gobase/api/internal/repository"
)

// Block creates a block relation between email1 and email2
func (i impl) Block(ctx context.Context, input model.MakeRelationship) error {
	newCtx := context.Background()
	return i.repo.DoInTx(newCtx, func(ctx context.Context, repo repository.Registry) error {
		if input.FromFriend == input.ToFriend {
			return errors.New("Duplicate email input")
		}

		user1, err1 := repo.System().FindUserByEmail(ctx, input.FromFriend)
		if err1 != nil {
			return err1
		}

		user2, err2 := repo.System().FindUserByEmail(ctx, input.ToFriend)
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
