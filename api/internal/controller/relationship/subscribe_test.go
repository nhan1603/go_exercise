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

func TestImpl_Subscribe(t *testing.T) {
	type arg struct {
		input              model.MakeRelationship
		mockDbOut          orm.User
		mockDBSubscribeErr error
		mockDBFindErr      error
		mockDBCheckErr     error
		expDBMockCalled    bool
		expErr             error
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
		"errDbExistedSubscribe": {
			input: model.MakeRelationship{
				FromFriend: "nhan.test123@test.com",
				ToFriend:   "nhan.test1234@test.com",
			},
			mockDbOut: orm.User{
				Email: "nhan.test123@test.com",
			},
			expDBMockCalled: true,
			mockDBCheckErr:  errors.New("error check existed Subscribe"),
			expErr:          errors.New("error check existed Subscribe"),
		},
		"errDbSubscribe": {
			input: model.MakeRelationship{
				FromFriend: "nhan.test123@test.com",
				ToFriend:   "nhan.test1234@test.com",
			},
			mockDbOut: orm.User{
				Email: "nhan.test123@test.com",
			},
			expDBMockCalled:    true,
			mockDBSubscribeErr: errors.New("error Subscribe"),
			expErr:             errors.New("error Subscribe"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			// Given:
			userRepo := user.MockRepository{}
			relaRepo := relationship.MockRepository{}
			if tc.expDBMockCalled {
				callSubscribe := relaRepo.On("Subscribe", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockDBSubscribeErr)
				callFind := userRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(tc.mockDbOut, tc.mockDBFindErr)
				callCheck := relaRepo.On("CheckExistedSubscribe", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockDBCheckErr)

				relaRepo.ExpectedCalls = []*mock.Call{}
				userRepo.ExpectedCalls = []*mock.Call{}

				if tc.mockDBFindErr != nil {
					userRepo.ExpectedCalls = append(userRepo.ExpectedCalls, callFind)
				} else if tc.mockDBCheckErr != nil {
					userRepo.ExpectedCalls = append(userRepo.ExpectedCalls, callFind)
					relaRepo.ExpectedCalls = append(relaRepo.ExpectedCalls, callCheck)
				} else {
					userRepo.ExpectedCalls = append(userRepo.ExpectedCalls, callFind)
					relaRepo.ExpectedCalls = append(relaRepo.ExpectedCalls, callCheck, callSubscribe)
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
			err := c.Subscribe(context.Background(), tc.input)

			// Then:
			require.Equal(t, tc.expErr, pkgerrors.Cause(err))
			userRepo.AssertExpectations(t)
			relaRepo.AssertExpectations(t)
			repo.AssertExpectations(t)
		})
	}
}
