package rest

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"gobase/api/internal/model"
)

// AddFriend will create a friendship for two email
func (i impl) AddFriend(ctx context.Context, input model.MakeRelationship) error {
	user1, err := i.repo.System().FindUserByEmail(ctx, input.FromFriend)
	if err != nil {
		return err
	}

	user2, err := i.repo.System().FindUserByEmail(ctx, input.ToFriend)
	if err != nil {
		return err
	}

	if err := i.repo.System().CheckExistedFriend(ctx, user1.ID, user2.ID); err != nil {
		return pkgerrors.WithStack(err)
	}

	return pkgerrors.WithStack(i.repo.System().AddFriend(ctx, user1.ID, user2.ID))
}