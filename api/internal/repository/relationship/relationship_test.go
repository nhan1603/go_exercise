package relationship

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

func TestImpl_CheckExistedFriend(t *testing.T) {
	os.Setenv("DB_URL", "postgres://gobase:@localhost:5432/gobase?sslmode=disable")
	type arg struct {
		givenCtx context.Context
		expErr   error
		email1   string
		email2   string
		existed  bool
	}

	tcs := map[string]arg{
		"success": {
			givenCtx: context.Background(),
			expErr:   nil,
			email1:   "nhan.tran3@test.com",
			email2:   "nhan.tran4@test.com",
			existed:  false,
		},
		"existed_email": {
			givenCtx: context.Background(),
			expErr:   errors.New("Cannot create new friendship."),
			email1:   "nhan.tran3@test.com",
			email2:   "nhan.tran4@test.com",
			existed:  true,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testutil.WithTxDB(t, func(dbConn pg.BeginnerExecutor) {
				// Given:
				repo := New(dbConn)

				user1 := orm.User{Email: tc.email1}
				require.NoError(t, user1.Insert(context.Background(), dbConn, boil.Infer()))

				user2 := orm.User{Email: tc.email2}
				require.NoError(t, user2.Insert(context.Background(), dbConn, boil.Infer()))

				if tc.existed {
					o := orm.Relationship{
						FirstEmailID:  user1.ID,
						SecondEmailID: user2.ID,
						Status:        RelationshipTypeFriend,
					}
					require.NoError(t, o.Insert(context.Background(), dbConn, boil.Infer()))
				}

				// When:
				err := repo.CheckExistedFriend(context.Background(), user1.ID, user2.ID)
				if tc.expErr == nil {
					require.NoError(t, err)
				} else {
					require.Equal(t, tc.expErr.Error(), err.Error())
				}
			})
		})
	}
}

func TestImpl_AddFriend(t *testing.T) {
	os.Setenv("DB_URL", "postgres://gobase:@localhost:5432/gobase?sslmode=disable")
	type arg struct {
		givenCtx context.Context
		expErr   error
		email1   string
		email2   string
	}
	tcs := map[string]arg{
		"success": {
			givenCtx: context.Background(),
			expErr:   nil,
			email1:   "nhan.tran3@test.com",
			email2:   "nhan.tran4@test.com",
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testutil.WithTxDB(t, func(dbConn pg.BeginnerExecutor) {
				// Given:
				repo := New(dbConn)

				user1 := orm.User{Email: tc.email1}
				require.NoError(t, user1.Insert(context.Background(), dbConn, boil.Infer()))

				user2 := orm.User{Email: tc.email2}
				require.NoError(t, user2.Insert(context.Background(), dbConn, boil.Infer()))

				// When:
				require.NoError(t, repo.AddFriend(context.Background(), user1.ID, user2.ID))
			})
		})
	}
}
