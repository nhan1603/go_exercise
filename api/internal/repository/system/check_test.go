package system

import (
	"context"
	"os"
	"testing"

	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gobase/api/pkg/db/pg"
	"gobase/api/pkg/testutil"
)

func TestImpl_CheckDB(t *testing.T) {
	os.Setenv("DB_URL", "postgres://gobase:@localhost:5432/gobase?sslmode=disable")
	cancelledCtx, c := context.WithCancel(context.Background())
	c()
	type arg struct {
		givenCtx context.Context
		expErr   error
	}
	tcs := map[string]arg{
		"success": {givenCtx: context.Background()},
		"ctx_cancelled": {
			givenCtx: cancelledCtx,
			expErr:   context.Canceled,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testutil.WithTxDB(t, func(dbConn pg.BeginnerExecutor) {
				// Given:
				repo := New(dbConn)

				// When:
				err := repo.CheckDB(tc.givenCtx)

				// Then:
				require.Equal(t, tc.expErr, pkgerrors.Cause(err))
			})
		})
	}
}
