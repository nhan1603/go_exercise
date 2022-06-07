package user

import (
	"context"

	"gobase/api/internal/repository/orm"
	"gobase/api/pkg/db/pg"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	FindUserByEmail(context.Context, string) (orm.User, error)
	CreateUser(context.Context, string) (int, error)
}

// New returns an implementation instance satisfying Repository
func New(dbConn pg.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn pg.ContextExecutor
}
