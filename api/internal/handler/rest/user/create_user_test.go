package user

import (
	"bytes"
	"errors"
	"gobase/api/internal/controller/relationship"
	"gobase/api/internal/controller/user"
	"gobase/api/pkg/httpserv"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestApiHandler_CreateUser(t *testing.T) {
	tcs := map[string]struct {
		expErr    *httpserv.Error
		errDb     error
		body      []byte
		expStatus int
	}{
		"error_invalid_body": {
			expStatus: http.StatusBadRequest,
			expErr:    &httpserv.Error{Status: http.StatusBadRequest, Code: "request_body_error", Desc: "Invalid request body"},
			body:      []byte(``),
			errDb:     nil,
		},
		"empty_email": {
			expStatus: http.StatusBadRequest,
			expErr:    &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "User has provided empty email"},
			body:      []byte(`{"email":""}`),
			errDb:     nil,
		},
		"failFromDB": {
			expStatus: http.StatusInternalServerError,
			expErr:    &httpserv.Error{Status: http.StatusInternalServerError, Code: "internal_error", Desc: "Something went wrong"},
			body:      []byte(`{"email":"andy@example.com"}`),
			errDb:     errors.New("cannot create user"),
		},
		"failFromDBInvalidUser": {
			expStatus: http.StatusNotFound,
			expErr:    &httpserv.Error{Status: http.StatusNotFound, Code: "invalid_email", Desc: "not found"},
			body:      []byte(`{"email":"andy@example.com"}`),
			errDb:     errors.New("not found"),
		},
		"success": {
			expStatus: http.StatusOK,
			expErr:    nil,
			body:      []byte(`{"email":"andy@example.com"}`),
			errDb:     nil,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/_/create-user", bytes.NewReader(tc.body))
			w := httptest.NewRecorder()

			mockRelaCtrl := relationship.MockApiRestController{}
			mockUserCtrl := user.MockApiRestController{}
			mockUserCtrl.On("CreateUser", mock.Anything, mock.Anything).Return(0, tc.errDb)

			// When:
			New(&mockUserCtrl, &mockRelaCtrl).CreateUser().ServeHTTP(w, r)
			// Then:
			require.Equal(t, tc.expStatus, w.Code)
			var actErr httpserv.Error
			err := httpserv.ParseJSON(w.Result().Body, &actErr)
			if tc.expErr == nil {
				require.Error(t, err)
			} else {
				require.Nil(t, err)
				require.Equal(t, tc.expErr.Code, actErr.Code)
				require.Equal(t, tc.expErr.Desc, actErr.Desc)
			}
		})
	}
}
