package relationship

import (
	"context"
	"errors"
	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gobase/api/internal/model"
	"gobase/api/internal/repository"
	"gobase/api/internal/repository/orm"
	"gobase/api/internal/repository/relationship"
	"gobase/api/internal/repository/user"
	"testing"
)

func TestImpl_Block(t *testing.T) {
	type arg struct {
		input           model.MakeRelationship
		mockDbOut       orm.User
		mockDBBlockErr  error
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
		"errDbExistedBlock": {
			input: model.MakeRelationship{
				FromFriend: "nhan.test123@test.com",
				ToFriend:   "nhan.test1234@test.com",
			},
			mockDbOut: orm.User{
				Email: "nhan.test123@test.com",
			},
			expDBMockCalled: true,
			mockDBCheckErr:  errors.New("error check existed block"),
			expErr:          errors.New("error check existed block"),
		},
		"errDbBlock": {
			input: model.MakeRelationship{
				FromFriend: "nhan.test123@test.com",
				ToFriend:   "nhan.test1234@test.com",
			},
			mockDbOut: orm.User{
				Email: "nhan.test123@test.com",
			},
			expDBMockCalled: true,
			mockDBBlockErr:  errors.New("error block"),
			expErr:          errors.New("error block"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			// Given:
			userRepo := user.MockRepository{}
			relaRepo := relationship.MockRepository{}
			if tc.expDBMockCalled {
				callBlock := relaRepo.On("Block", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockDBBlockErr)
				callFind := userRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(tc.mockDbOut, tc.mockDBFindErr)
				callCheck := relaRepo.On("CheckExistedBlock", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockDBCheckErr)

				userRepo.ExpectedCalls = []*mock.Call{}
				relaRepo.ExpectedCalls = []*mock.Call{}

				if tc.mockDBFindErr != nil {
					userRepo.ExpectedCalls = append(userRepo.ExpectedCalls, callFind)
				} else if tc.mockDBCheckErr != nil {
					userRepo.ExpectedCalls = append(userRepo.ExpectedCalls, callFind)
					relaRepo.ExpectedCalls = append(relaRepo.ExpectedCalls, callCheck)
				} else {
					userRepo.ExpectedCalls = append(userRepo.ExpectedCalls, callFind)
					relaRepo.ExpectedCalls = append(relaRepo.ExpectedCalls, callCheck, callBlock)
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
			err := c.Block(context.Background(), tc.input)

			// Then:
			require.Equal(t, tc.expErr, pkgerrors.Cause(err))
			userRepo.AssertExpectations(t)
			relaRepo.AssertExpectations(t)
			repo.AssertExpectations(t)
		})
	}
}
