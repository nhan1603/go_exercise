package system

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gobase/api/pkg/db/pg"
	"gobase/api/pkg/testutil"
	"os"
	"testing"
)

func TestImpl_Subscribe(t *testing.T) {
	os.Setenv("DB_URL", "postgres://gobase:@localhost:5432/gobase?sslmode=disable")
	type arg struct {
		expErr error
		email1 string
		email2 string
	}
	tcs := map[string]arg{
		"success": {
			email1: "nhan.tran3@test.com",
			email2: "nhan.tran4@test.com",
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testutil.WithTxDB(t, func(dbConn pg.BeginnerExecutor) {
				// Given:
				repo := New(dbConn)

				id1, err1 := repo.CreateUser(context.Background(), tc.email1)

				id2, err2 := repo.CreateUser(context.Background(), tc.email1)

				require.NoError(t, err1)
				require.NoError(t, err2)

				// When:
				err := repo.Subscribe(context.Background(), id1, id2)
				if tc.expErr == nil {
					require.NoError(t, err)
				} else {
					require.Equal(t, tc.expErr, err)
				}
			})
		})
	}
}

func TestImpl_CheckExistedSubscribe(t *testing.T) {
	os.Setenv("DB_URL", "postgres://gobase:@localhost:5432/gobase?sslmode=disable")
	type arg struct {
		expErr  error
		email1  string
		email2  string
		existed bool
	}

	tcs := map[string]arg{
		"success": {
			email1:  "nhan.tran3@test.com",
			email2:  "nhan.tran4@test.com",
			existed: false,
		},
		"existed_subscribe": {
			expErr:  errors.New("Cannot create new subscription."),
			email1:  "nhan.tran3@test.com",
			email2:  "nhan.tran4@test.com",
			existed: true,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testutil.WithTxDB(t, func(dbConn pg.BeginnerExecutor) {
				// Given:
				repo := New(dbConn)

				id1, err1 := repo.CreateUser(context.Background(), tc.email1)

				id2, err2 := repo.CreateUser(context.Background(), tc.email2)

				require.NoError(t, err1)
				require.NoError(t, err2)

				if tc.existed {
					require.NoError(t, repo.Subscribe(context.Background(), id1, id2))
				}

				// When:
				err := repo.CheckExistedSubscribe(context.Background(), id1, id2)
				if tc.expErr == nil {
					require.NoError(t, err)
				} else {
					require.Equal(t, tc.expErr.Error(), err.Error())
				}
			})
		})
	}
}
