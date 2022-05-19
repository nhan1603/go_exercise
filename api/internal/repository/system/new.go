package system

import (
	"context"

	"gobase/api/pkg/db/pg"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	// CheckDB will check if calls to DB are successful or not
	CheckDB(context.Context) error
	AddFriend(ctx context.Context, email1, email2 string) error
	CreateUser(ctx context.Context, email string) error
	FindFriendList(ctx context.Context, email string) ([]string, error)
	FindCommonFriends(ctx context.Context, email1, email2 string) ([]string, error)
	Subscribe(ctx context.Context, email1, email2 string) error
}

// New returns an implementation instance satisfying Repository
func New(dbConn pg.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn pg.ContextExecutor
}
