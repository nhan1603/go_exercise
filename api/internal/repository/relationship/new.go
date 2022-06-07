package relationship

import (
	"context"
	"gobase/api/pkg/db/pg"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	CheckExistedFriend(context.Context, int, int) error
	AddFriend(context.Context, int, int) error
	FindFriendList(context.Context, int) ([]string, error)
	CheckExistedSubscribe(context.Context, int, int) error
	Subscribe(context.Context, int, int) error
	CheckExistedBlock(context.Context, int, int) error
	Block(context.Context, int, int) error
	UpdateReceiver(context.Context, int, []string) ([]string, error)
}

// New returns an implementation instance satisfying Repository
func New(dbConn pg.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn pg.ContextExecutor
}
