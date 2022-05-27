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

func TestImpl_Subscribe(t *testing.T) {
	type arg struct {
		input              model.MakeRelationship
		mockDbOut          *orm.User
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
		"errDbExistedSubscribe": {
			input: model.MakeRelationship{
				FromFriend: "nhan.test123@test.com",
				ToFriend:   "nhan.test1234@test.com",
			},
			mockDbOut: &orm.User{
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
			mockDbOut: &orm.User{
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
			systemRepo := system.MockRepository{}
			if tc.expDBMockCalled {
				callSubscribe := systemRepo.On("Subscribe", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockDBSubscribeErr)
				callFind := systemRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(tc.mockDbOut, tc.mockDBFindErr)
				callCheck := systemRepo.On("CheckExistedSubscribe", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockDBCheckErr)
				systemRepo.ExpectedCalls = []*mock.Call{}

				if tc.mockDBFindErr != nil {
					systemRepo.ExpectedCalls = append(systemRepo.ExpectedCalls, callFind)
				} else if tc.mockDBCheckErr != nil {
					systemRepo.ExpectedCalls = append(systemRepo.ExpectedCalls, callFind, callCheck)
				} else {
					systemRepo.ExpectedCalls = append(systemRepo.ExpectedCalls, callFind, callSubscribe, callCheck)
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
			err := c.Subscribe(context.Background(), tc.input)

			// Then:
			require.Equal(t, tc.expErr, pkgerrors.Cause(err))
			systemRepo.AssertExpectations(t)
			repo.AssertExpectations(t)
		})
	}
}
