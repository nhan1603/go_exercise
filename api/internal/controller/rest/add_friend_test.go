package rest

import (
	"context"
	"errors"
	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gobase/api/internal/model"
	"gobase/api/internal/repository"
	"gobase/api/internal/repository/orm"
	"gobase/api/internal/repository/system"
	"testing"
)

func TestImpl_AddFriend(t *testing.T) {
	type arg struct {
		input           model.MakeRelationship
		mockDbOut       *orm.User
		mockDBAddErr    error
		mockDBFindErr   error
		mockDBCheckErr  error
		expDBMockCalled bool
		expErr          error
	}
	tcs := map[string]arg{
		"success": {
			input: model.MakeRelationship{
				FromFriend: "nhan.test123@test.com",
				ToFriend:   "nhan.test1234@test.com",
			},
			mockDbOut: &orm.User{
				Email: "nhan.test123@test.com",
			},
			expDBMockCalled: true,
		},
		"errDbFindUser": {
			input: model.MakeRelationship{
				FromFriend: "nhan.test123@test.com",
				ToFriend:   "nhan.test1234@test.com",
			},
			mockDbOut: &orm.User{
				Email: "nhan.test123@test.com",
			},
			expDBMockCalled: true,
			mockDBFindErr:   errors.New("error find user"),
			expErr:          errors.New("error find user"),
		},
		"errDbExistedFriend": {
			input: model.MakeRelationship{
				FromFriend: "nhan.test123@test.com",
				ToFriend:   "nhan.test1234@test.com",
			},
			mockDbOut: &orm.User{
				Email: "nhan.test123@test.com",
			},
			expDBMockCalled: true,
			mockDBCheckErr:  errors.New("error check existed friend"),
			expErr:          errors.New("error check existed friend"),
		},
		"errDbAdd": {
			input: model.MakeRelationship{
				FromFriend: "nhan.test123@test.com",
				ToFriend:   "nhan.test1234@test.com",
			},
			mockDbOut: &orm.User{
				Email: "nhan.test123@test.com",
			},
			expDBMockCalled: true,
			mockDBAddErr:    errors.New("error add friend"),
			expErr:          errors.New("error add friend"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			// Given:
			systemRepo := system.MockRepository{}
			if tc.expDBMockCalled {
				callAdd := systemRepo.On("AddFriend", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockDBAddErr)
				callFind := systemRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(tc.mockDbOut, tc.mockDBFindErr)
				callCheck := systemRepo.On("CheckExistedFriend", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockDBCheckErr)
				systemRepo.ExpectedCalls = []*mock.Call{}

				if tc.mockDBFindErr != nil {
					systemRepo.ExpectedCalls = append(systemRepo.ExpectedCalls, callFind)
				} else if tc.mockDBCheckErr != nil {
					systemRepo.ExpectedCalls = append(systemRepo.ExpectedCalls, callFind, callCheck)
				} else {
					systemRepo.ExpectedCalls = append(systemRepo.ExpectedCalls, callFind, callAdd, callCheck)
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
			err := c.AddFriend(context.Background(), tc.input)

			// Then:
			require.Equal(t, tc.expErr, pkgerrors.Cause(err))
			systemRepo.AssertExpectations(t)
		})
	}
}
