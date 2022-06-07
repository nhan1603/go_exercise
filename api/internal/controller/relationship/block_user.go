package relationship

import (
	"context"
	"gobase/api/internal/model"
)

// Block creates a block relation between email1 and email2
func (i impl) Block(ctx context.Context, input model.MakeRelationship) error {
	user1, err := i.repo.User().FindUserByEmail(ctx, input.FromFriend)
	if err != nil {
		return err
	}

	user2, err := i.repo.User().FindUserByEmail(ctx, input.ToFriend)
	if err != nil {
		return err
	}

	emailId1 := user1.ID
	emailId2 := user2.ID

	if err := i.repo.Relationship().CheckExistedBlock(ctx, emailId1, emailId2); err != nil {
		return err
	}

	return i.repo.Relationship().Block(ctx, emailId1, emailId2)
}
