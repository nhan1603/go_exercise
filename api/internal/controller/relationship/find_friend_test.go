package relationship

import (
	"context"
	"database/sql"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gobase/api/internal/repository"
	"gobase/api/internal/repository/orm"
	"gobase/api/internal/repository/relationship"
	"gobase/api/internal/repository/user"
)

func TestImpl_FindFriendList(t *testing.T) {
	type arg struct {
		input               string
		mockDbOut           orm.User
		mockDbLsFriendOut   []string
		mockDBFindErr       error
		mockDBFindFriendErr error
		expDBMockCalled     bool
		expErr              error
	}
	tcs := map[string]arg{
		"success": {
			input: "nhan.test12345@test.com",
			mockDbOut: orm.User{
				Email: "nhan.test12345@test.com",
			},
			mockDbLsFriendOut: []string{"nhan.tran3@testtest.com"},
			expDBMockCalled:   true,
		},
		"errDbNonExistedUser": {
			input:           "nhan.test123@test.com",
			mockDbOut:       orm.User{},
			mockDBFindErr:   sql.ErrNoRows,
			expDBMockCalled: true,
			expErr:          sql.ErrNoRows,
		},
		"errDbFindListFr": {
			input: "nhan.test123@test.com",
			mockDbOut: orm.User{
				Email: "nhan.test12345@test.com",
			},
			expDBMockCalled:     true,
			mockDBFindFriendErr: errors.New("orm: unable to select from user and relationship"),
			expErr:              errors.New("orm: unable to select from user and relationship"),
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			// Given:
			userRepo := user.MockRepository{}
			relaRepo := relationship.MockRepository{}
			if tc.expDBMockCalled {
				callFind := userRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(tc.mockDbOut, tc.mockDBFindErr)
				callFindFriend := relaRepo.On("FindFriendList", mock.Anything, mock.Anything).Return(tc.mockDbLsFriendOut, tc.mockDBFindFriendErr)
				userRepo.ExpectedCalls = []*mock.Call{}
				relaRepo.ExpectedCalls = []*mock.Call{}

				if tc.mockDBFindErr != nil {
					userRepo.ExpectedCalls = append(userRepo.ExpectedCalls, callFind)
				} else {
					userRepo.ExpectedCalls = append(userRepo.ExpectedCalls, callFind)
					relaRepo.ExpectedCalls = append(relaRepo.ExpectedCalls, callFindFriend)
				}
			}

			repo := repository.MockRegistry{}
			if tc.expDBMockCalled {
				repo.ExpectedCalls = []*mock.Call{
					repo.On("User").Return(&userRepo),
				}
				if tc.mockDBFindErr == nil {
					repo.ExpectedCalls = append(repo.ExpectedCalls, repo.On("Relationship").Return(&relaRepo))
				}
			}

			c := New(&repo)

			// When:
			lsFr, err := c.FindFriendList(context.Background(), tc.input)

			// Then:
			require.Equal(t, tc.mockDbLsFriendOut, lsFr)
			if tc.expErr == nil {
				require.NoError(t, err)
			} else {
				require.Equal(t, tc.expErr.Error(), err.Error())
			}
			userRepo.AssertExpectations(t)
			relaRepo.AssertExpectations(t)
			repo.AssertExpectations(t)
		})
	}

}
