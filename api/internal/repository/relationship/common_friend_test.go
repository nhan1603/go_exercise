package relationship

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gobase/api/internal/repository/orm"
	"gobase/api/pkg/db/pg"
	"gobase/api/pkg/testutil"
	"os"
	"testing"
)

func TestImpl_FindFriendList(t *testing.T) {
	os.Setenv("DB_URL", "postgres://gobase:@localhost:5432/gobase?sslmode=disable")
	type arg struct {
		givenCtx context.Context
		expErr   error
		email1   string
		email2   string
		expRes   []string
	}
	tcs := map[string]arg{
		"success": {
			givenCtx: context.Background(),
			expErr:   nil,
			email1:   "nhan.tran3@test.com",
			email2:   "nhan.tran4@test.com",
			expRes:   []string{"nhan.tran4@test.com"},
		},
		"empty": {
			givenCtx: context.Background(),
			expErr:   nil,
			email1:   "nhan.tran3@test.com",
			expRes:   []string(nil),
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testutil.WithTxDB(t, func(dbConn pg.BeginnerExecutor) {
				// Given:
				repo := New(dbConn)

				user1 := orm.User{Email: tc.email1}
				require.NoError(t, user1.Insert(context.Background(), dbConn, boil.Infer()))

				if tc.email2 != "" {
					user2 := orm.User{Email: tc.email2}
					require.NoError(t, user2.Insert(context.Background(), dbConn, boil.Infer()))
					// When:
					require.NoError(t, repo.AddFriend(context.Background(), user1.ID, user2.ID))
				}

				frList, err := repo.FindFriendList(context.Background(), user1.ID)
				require.NoError(t, err)

				require.Equal(t, frList, tc.expRes)

			})
		})
	}
}
