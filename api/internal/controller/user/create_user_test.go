package user

import (
	"context"
	"database/sql"
	"errors"
	frerrors "github.com/friendsofgo/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gobase/api/internal/repository"
	"gobase/api/internal/repository/orm"
	"gobase/api/internal/repository/user"
	"testing"
)

func TestImpl_CreateUser(t *testing.T) {
	type arg struct {
		input           string
		mockDbOut       orm.User
		mockDBFindErr   error
		mockDBCreateErr error
		expDBMockCalled bool
		expErr          error
	}
	tcs := map[string]arg{
		"success": {
			input:           "nhan.test12345@test.com",
			mockDbOut:       orm.User{},
			mockDBFindErr:   sql.ErrNoRows,
			expDBMockCalled: true,
		},
		"errDbExistedUser": {
			input: "nhan.test123@test.com",
			mockDbOut: orm.User{
				Email: "nhan.test123@test.com",
			},
			expDBMockCalled: true,
			mockDBFindErr:   nil,
			expErr:          frerrors.New("Existed email input."),
		},
		"errDbCreate": {
			input:           "nhan.test123@test.com",
			mockDbOut:       orm.User{},
			mockDBFindErr:   sql.ErrNoRows,
			expDBMockCalled: true,
			mockDBCreateErr: errors.New("error from database"),
			expErr:          errors.New("error from database"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			// Given:
			userRepo := user.MockRepository{}
			if tc.expDBMockCalled {
				callFind := userRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(tc.mockDbOut, tc.mockDBFindErr)
				callCreate := userRepo.On("CreateUser", mock.Anything, mock.Anything).Return(0, tc.mockDBCreateErr)
				userRepo.ExpectedCalls = []*mock.Call{}

				if tc.mockDBFindErr == nil {
					userRepo.ExpectedCalls = append(userRepo.ExpectedCalls, callFind)
				} else {
					userRepo.ExpectedCalls = append(userRepo.ExpectedCalls, callFind, callCreate)
				}
			}

			repo := repository.MockRegistry{}
			if tc.expDBMockCalled {
				repo.ExpectedCalls = []*mock.Call{
					repo.On("User").Return(&userRepo),
				}
			}

			c := New(&repo)

			// When:
			id, err := c.CreateUser(context.Background(), tc.input)

			// Then:
			require.Equal(t, 0, id)
			if tc.expErr == nil {
				require.NoError(t, err)
			} else {
				require.Equal(t, tc.expErr.Error(), err.Error())
			}
			userRepo.AssertExpectations(t)
			repo.AssertExpectations(t)
		})
	}
}
