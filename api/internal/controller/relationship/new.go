package relationship

import (
	"context"
	"gobase/api/internal/model"
	"gobase/api/internal/repository"
)

// ApiRestController represents the specification of this pkg
type ApiRestController interface {
	AddFriend(context.Context, model.MakeRelationship) error
	FindFriendList(context.Context, string) ([]string, error)
	FindCommonFriends(context.Context, model.CommonFriend) ([]string, error)
	Subscribe(context.Context, model.MakeRelationship) error
	Block(context.Context, model.MakeRelationship) error
	UpdateReceiver(context.Context, string, string) ([]string, error)
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) ApiRestController {
	return impl{repo: repo}
}

type impl struct {
	repo repository.Registry
}
