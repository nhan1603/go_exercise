package rest

import (
	"context"
	"gobase/api/internal/model"
)

// Block creates a block relation between email1 and email2
func (i impl) Block(ctx context.Context, input model.MakeRelationship) error {
	user1, err1 := i.repo.System().FindUserByEmail(ctx, input.FromFriend)
	if err1 != nil {
		return err1
	}

	user2, err2 := i.repo.System().FindUserByEmail(ctx, input.ToFriend)
	if err2 != nil {
		return err2
	}

	emailId1 := user1.ID
	emailId2 := user2.ID

	if err := i.repo.System().CheckExistedBlock(ctx, emailId1, emailId2); err != nil {
		return err
	}

	return i.repo.System().Block(ctx, emailId1, emailId2)
}
