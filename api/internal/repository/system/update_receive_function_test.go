package system

import (
	"context"
	"github.com/stretchr/testify/require"
	"gobase/api/pkg/db/pg"
	"gobase/api/pkg/testutil"
	"os"
	"testing"
)

func TestImpl_UpdateReceiver(t *testing.T) {
	os.Setenv("DB_URL", "postgres://gobase:@localhost:5432/gobase?sslmode=disable")
	type arg struct {
		expErr         error
		sendEmail      string
		subscribeEmail string
		friendEmail    string
		blockEmail     string
		mentionedEmail []string
		expRes         []string
	}
	tcs := map[string]arg{
		"success": {
			sendEmail:      "nhan.tran@test.com",
			subscribeEmail: "nhan.tran1@test.com",
			friendEmail:    "nhan.tran2@test.com",
			blockEmail:     "nhan.tran3@test.com",
			mentionedEmail: []string{"nhan.tran4@test.com", "nhan.tran3@test.com", "nhan.tran5@test.com"},
			expRes:         []string{"nhan.tran1@test.com", "nhan.tran2@test.com", "nhan.tran4@test.com"},
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			testutil.WithTxDB(t, func(dbConn pg.BeginnerExecutor) {
				// Given:
				repo := New(dbConn)

				id, errr := repo.CreateUser(context.Background(), tc.sendEmail)

				id1, err1 := repo.CreateUser(context.Background(), tc.subscribeEmail)

				id2, err2 := repo.CreateUser(context.Background(), tc.friendEmail)

				id3, err3 := repo.CreateUser(context.Background(), tc.blockEmail)

				_, err4 := repo.CreateUser(context.Background(), tc.mentionedEmail[0])

				require.NoError(t, errr)
				require.NoError(t, err1)
				require.NoError(t, err2)
				require.NoError(t, err3)
				require.NoError(t, err4)

				// When:

				require.NoError(t, repo.Subscribe(context.Background(), id1, id))
				require.NoError(t, repo.AddFriend(context.Background(), id2, id))
				require.NoError(t, repo.Block(context.Background(), id3, id))

				result, err := repo.UpdateReceiver(context.Background(), id, tc.mentionedEmail)

				require.NoError(t, err)

				require.Equal(t, result, tc.expRes)
			})
		})
	}
}
