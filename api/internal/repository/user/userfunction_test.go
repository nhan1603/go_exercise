package user

import (
	"context"
	"gobase/api/internal/repository/orm"
	"gobase/api/pkg/db/pg"
	"gobase/api/pkg/testutil"
	"os"
	"testing"

	"github.com/friendsofgo/errors"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestImpl_CreateUser(t *testing.T) {
	os.Setenv("DB_URL", "postgres://gobase:@localhost:5432/gobase?sslmode=disable")
	type arg struct {
		expErr        error
		checkIfExists bool
		email         string
	}
	tcs := map[string]arg{
		"success": {
			email: "nhan.tran3@test.com",
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testutil.WithTxDB(t, func(dbConn pg.BeginnerExecutor) {
				// Given:
				repo := New(dbConn)

				if tc.checkIfExists {
					o := orm.User{Email: tc.email}
					require.NoError(t, o.Insert(context.Background(), dbConn, boil.Infer()))
				}

				// When:
				_, err := repo.CreateUser(context.Background(), tc.email)
				if tc.expErr == nil {
					require.NoError(t, err)
				} else {
					require.Equal(t, tc.expErr, err)
				}
			})
		})
	}
}

func TestImpl_FindUserByEmail(t *testing.T) {
	os.Setenv("DB_URL", "postgres://gobase:@localhost:5432/gobase?sslmode=disable")
	type arg struct {
		givenCtx context.Context
		expErr   error
		email    string
		existed  bool
	}

	tcs := map[string]arg{
		"success": {
			givenCtx: context.Background(),
			expErr:   nil,
			email:    "nhan.tran3@test.com",
			existed:  true,
		},
		"existed_email": {
			givenCtx: context.Background(),
			expErr:   errors.New("not found"),
			email:    "nhan.tran3@test.com",
			existed:  false,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testutil.WithTxDB(t, func(dbConn pg.BeginnerExecutor) {
				// Given:
				repo := New(dbConn)

				if tc.existed {
					o := orm.User{Email: tc.email}
					require.NoError(t, o.Insert(context.Background(), dbConn, boil.Infer()))
				}

				// When:
				_, err := repo.FindUserByEmail(context.Background(), tc.email)
				if tc.expErr == nil {
					require.NoError(t, err)
				} else {
					require.Equal(t, tc.expErr.Error(), err.Error())
				}
			})
		})
	}
}
