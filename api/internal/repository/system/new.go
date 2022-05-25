package system

import (
	"context"
	"gobase/api/internal/repository/orm"

	"gobase/api/pkg/db/pg"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	// CheckDB will check if calls to DB are successful or not
	CheckDB(context.Context) error
	FindUserByEmail(context.Context, string) (*orm.User, error)
	CheckExistedFriend(context.Context, int, int) error
	AddFriend(context.Context, int, int) error
	CreateUser(context.Context, string) (int, error)
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
