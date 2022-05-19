package system

import (
	"context"

	"gobase/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	// CheckReadiness checks if the system is ready for operation or not
	CheckReadiness(ctx context.Context) error
	AddFriend(ctx context.Context, email1, email2 string) error
	CreateUser(ctx context.Context, email string) error
	FindFriendList(ctx context.Context, email string) ([]string, error)
	FindCommonFriends(ctx context.Context, email1, email2 string) ([]string, error)
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) Controller {
	return impl{repo: repo}
}

type impl struct {
	repo repository.Registry
}
