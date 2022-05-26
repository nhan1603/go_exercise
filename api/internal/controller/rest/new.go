package rest

import (
	"context"
	"gobase/api/internal/model"
	"gobase/api/internal/repository"
)

// ApiRestController represents the specification of this pkg
type ApiRestController interface {
	AddFriend(context.Context, model.MakeFriend) error
	CreateUser(ctx context.Context, email string) (int, error)
	FindFriendList(ctx context.Context, email string) ([]string, error)
	FindCommonFriends(ctx context.Context, email1, email2 string) ([]string, error)
	Subscribe(ctx context.Context, email1, email2 string) error
	Block(ctx context.Context, email1, email2 string) error
	UpdateReceiver(ctx context.Context, email, message string) ([]string, error)
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) ApiRestController {
	return impl{repo: repo}
}

type impl struct {
	repo repository.Registry
}
