package rest

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gobase/api/internal/repository"
	"gobase/api/internal/repository/orm"
	"gobase/api/internal/repository/system"
	"testing"
)

func TestImpl_UpdateReceiver(t *testing.T) {
	type arg struct {
		email             string
		message           string
		mockDbOut         *orm.User
		mockDbLsEmailOut  []string
		mockDBFindErr     error
		mockDBReceiverErr error
		expDBMockCalled   bool
		expErr            error
	}

	tcs := map[string]arg{
		"success": {
			email: "nhan.test12345@test.com",
			mockDbOut: &orm.User{
				Email: "nhan.test12345@test.com",
			},
			mockDbLsEmailOut: []string{"nhan.tran3@testtest.com"},
			expDBMockCalled:  true,
		},
		"errDbNonExistedUser": {
			email:           "nhan.test123@test.com",
			mockDbOut:       nil,
			mockDBFindErr:   sql.ErrNoRows,
			expDBMockCalled: true,
			expErr:          sql.ErrNoRows,
		},
		"errDbFindListEmail": {
			email: "nhan.test123@test.com",
			mockDbOut: &orm.User{
				Email: "nhan.test12345@test.com",
			},
			expDBMockCalled:   true,
			mockDBReceiverErr: errors.New("orm: unable to select from user and relationship"),
			expErr:            errors.New("orm: unable to select from user and relationship"),
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			// Given:
			systemRepo := system.MockRepository{}
			if tc.expDBMockCalled {
				callFind := systemRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(tc.mockDbOut, tc.mockDBFindErr)
				callFindEmail := systemRepo.On("UpdateReceiver", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockDbLsEmailOut, tc.mockDBReceiverErr)
				systemRepo.ExpectedCalls = []*mock.Call{}

				if tc.mockDBFindErr != nil {
					systemRepo.ExpectedCalls = append(systemRepo.ExpectedCalls, callFind)
				} else {
					systemRepo.ExpectedCalls = append(systemRepo.ExpectedCalls, callFind, callFindEmail)
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
			lsFr, err := c.UpdateReceiver(context.Background(), tc.email, tc.message)

			// Then:
			require.Equal(t, tc.mockDbLsEmailOut, lsFr)
			if tc.expErr == nil {
				require.NoError(t, err)
			} else {
				require.Equal(t, tc.expErr.Error(), err.Error())
			}
			systemRepo.AssertExpectations(t)
			repo.AssertExpectations(t)
		})
	}
}
