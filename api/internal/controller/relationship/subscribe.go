package relationship

import (
	"context"
	"gobase/api/internal/model"
)

// Subscribe will create a subscription for of the second email for the first email
func (i impl) Subscribe(ctx context.Context, input model.MakeRelationship) error {
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

	if err := i.repo.Relationship().CheckExistedSubscribe(ctx, emailId1, emailId2); err != nil {
		return err
	}

	return i.repo.Relationship().Subscribe(ctx, emailId1, emailId2)
}
