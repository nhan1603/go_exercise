package rest

import (
	"context"
	"errors"
	frerrors "github.com/friendsofgo/errors"
	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gobase/api/internal/repository"
	"gobase/api/internal/repository/orm"
	"gobase/api/internal/repository/system"
	"testing"
)

func TestImpl_CreateUser(t *testing.T) {
	type arg struct {
		input           string
		mockDbOut       *orm.User
		mockDBFindErr   error
		mockDBCreateErr error
		expDBMockCalled bool
		expErr          error
	}
	tcs := map[string]arg{
		"success": {
			input:           "nhan.test12345@test.com123",
			mockDbOut:       nil,
			expDBMockCalled: true,
		},
		"errDbExistedUser": {
			input: "nhan.test123@test.com",
			mockDbOut: &orm.User{
				Email: "nhan.test123@test.com",
			},
			expDBMockCalled: true,
			mockDBFindErr:   errors.New("error existed user"),
			expErr:          frerrors.New("Existed email input."),
		},
		"errDbCreate": {
			input: "nhan.test123@test.com",
			mockDbOut: &orm.User{
				Email: "nhan.test123@test.com",
			},
			expDBMockCalled: true,
			mockDBCreateErr: errors.New("error from database"),
			expErr:          errors.New("error from database"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			// Given:
			systemRepo := system.MockRepository{}
			if tc.expDBMockCalled {
				callFind := systemRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(tc.mockDbOut, tc.mockDBFindErr)
				callCreate := systemRepo.On("CreateUser", mock.Anything, mock.Anything).Return(0, tc.mockDBCreateErr)
				systemRepo.ExpectedCalls = []*mock.Call{}

				if tc.mockDBFindErr != nil {
					systemRepo.ExpectedCalls = append(systemRepo.ExpectedCalls, callFind, callCreate)
				} else {
					systemRepo.ExpectedCalls = append(systemRepo.ExpectedCalls, callFind, callCreate)
				}
			}

			repo := repository.MockRegistry{}
			if tc.expDBMockCalled {
				repo.ExpectedCalls = []*mock.Call{
					repo.On("System").Return(&systemRepo),
				}
			}

			c := New(&repo)

			// When:
			id, err := c.CreateUser(context.Background(), tc.input)

			// Then:
			require.Equal(t, 0, id)
			require.Equal(t, tc.expErr, pkgerrors.Cause(err))
			systemRepo.AssertExpectations(t)
			repo.AssertExpectations(t)
		})
	}
}
