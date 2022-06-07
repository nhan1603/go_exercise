package repository

import (
	"context"
	"gobase/api/internal/repository/relationship"
	"gobase/api/internal/repository/user"
	"time"

	"github.com/cenkalti/backoff/v4"
	pkgerrors "github.com/pkg/errors"
	"gobase/api/pkg/db/pg"
)

// Registry is the registry of all the domain specific repositories and also provides transaction capabilities.
type Registry interface {
	// User returns the user repo
	User() user.Repository
	// Relationship returns the relationship repo
	Relationship() relationship.Repository
	// DoInTx wraps operations within a db tx
	DoInTx(ctx context.Context, txFunc func(ctx context.Context, txRepo Registry) error, overrideBackoffPolicy backoff.BackOff) error
}

// New returns a new instance of Registry
func New(dbConn pg.BeginnerExecutor) Registry {
	return impl{
		dbConn:       dbConn,
		user:         user.New(dbConn),
		relationship: relationship.New(dbConn),
	}
}

type impl struct {
	dbConn       pg.BeginnerExecutor // Only used to start DB txns
	tx           pg.ContextExecutor  // Only used to keep track if txn has already been started to prevent devs from accidentally creating nested txns
	user         user.Repository
	relationship relationship.Repository
}

// User returns the user repo
func (i impl) User() user.Repository {
	return i.user
}

// Relationship returns the relationship repo
func (i impl) Relationship() relationship.Repository {
	return i.relationship
}

// DoInTx wraps operations within a db tx
func (i impl) DoInTx(ctx context.Context, txFunc func(ctx context.Context, txRepo Registry) error, overrideBackoffPolicy backoff.BackOff) error {
	if i.tx != nil {
		return pkgerrors.WithStack(errNestedTx)
	}

	if overrideBackoffPolicy == nil {
		overrideBackoffPolicy = pg.ExponentialBackOff(3, time.Minute)
	}

	return pg.TxWithBackOff(ctx, overrideBackoffPolicy, i.dbConn, func(tx pg.ContextExecutor) error {
		newI := impl{
			tx:           tx,
			user:         user.New(tx),
			relationship: relationship.New(tx),
		}
		return txFunc(ctx, newI)
	})
}
