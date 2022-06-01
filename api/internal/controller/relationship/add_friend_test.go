package relationship

import (
	"context"
	"errors"
	"testing"

	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gobase/api/internal/model"
	"gobase/api/internal/repository"
	"gobase/api/internal/repository/orm"
	"gobase/api/internal/repository/relationship"
	"gobase/api/internal/repository/user"
)

func TestImpl_AddFriend(t *testing.T) {
	type arg struct {
		input           model.MakeRelationship
		mockDbOut       orm.User
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
			mockDbOut: orm.User{
				Email: "nhan.test123@test.com",
			},
			expDBMockCalled: true,
		},
		"errDbFindUser": {
			input: model.MakeRelationship{
				FromFriend: "nhan.test123@test.com",
				ToFriend:   "nhan.test1234@test.com",
			},
			mockDbOut: orm.User{
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
			mockDbOut: orm.User{
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
			mockDbOut: orm.User{
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
			userRepo := user.MockRepository{}
			relaRepo := relationship.MockRepository{}
			if tc.expDBMockCalled {
				callAdd := relaRepo.On("AddFriend", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockDBAddErr)
				callFind := userRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(tc.mockDbOut, tc.mockDBFindErr)
				callCheck := relaRepo.On("CheckExistedFriend", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockDBCheckErr)
				userRepo.ExpectedCalls = []*mock.Call{}
				relaRepo.ExpectedCalls = []*mock.Call{}

				if tc.mockDBFindErr != nil {
					userRepo.ExpectedCalls = append(userRepo.ExpectedCalls, callFind)
				} else if tc.mockDBCheckErr != nil {
					userRepo.ExpectedCalls = append(userRepo.ExpectedCalls, callFind)
					relaRepo.ExpectedCalls = append(relaRepo.ExpectedCalls, callCheck)
				} else {
					userRepo.ExpectedCalls = append(userRepo.ExpectedCalls, callFind)
					relaRepo.ExpectedCalls = append(relaRepo.ExpectedCalls, callCheck, callAdd)
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
			err := c.AddFriend(context.Background(), tc.input)

			// Then:
			require.Equal(t, tc.expErr, pkgerrors.Cause(err))
			userRepo.AssertExpectations(t)
			relaRepo.AssertExpectations(t)
			repo.AssertExpectations(t)
		})
	}
}
