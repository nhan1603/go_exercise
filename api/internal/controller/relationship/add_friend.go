package relationship

import (
	"context"
	"gobase/api/internal/model"
)

// AddFriend will create a friendship for two email
func (i impl) AddFriend(ctx context.Context, input model.MakeRelationship) error {
	user1, err := i.repo.User().FindUserByEmail(ctx, input.FromFriend)
	if err != nil {
		return err
	}

	user2, err := i.repo.User().FindUserByEmail(ctx, input.ToFriend)
	if err != nil {
		return err
	}

	if err := i.repo.Relationship().CheckExistedFriend(ctx, user1.ID, user2.ID); err != nil {
		return err
	}

	return i.repo.Relationship().AddFriend(ctx, user1.ID, user2.ID)
}
