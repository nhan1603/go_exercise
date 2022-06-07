package relationship

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

func TestApiHandler_UpdateReceiver(t *testing.T) {
	tcs := map[string]struct {
		expErr    *httpserv.Error
		errDb     error
		body      []byte
		expStatus int
		resultDb  []string
	}{
		"error_invalid_body": {
			expStatus: http.StatusBadRequest,
			expErr:    &httpserv.Error{Status: http.StatusBadRequest, Code: "request_body_error", Desc: "Invalid request body"},
			body:      []byte(``),
			errDb:     nil,
			resultDb:  []string{},
		},
		"empty_email": {
			expStatus: http.StatusBadRequest,
			expErr:    &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "User has provided empty email"},
			body:      []byte(`{"sender":"","text":"andy@example.com"}`),
			errDb:     nil,
			resultDb:  []string{},
		},
		"failFromDB": {
			expStatus: http.StatusInternalServerError,
			expErr:    &httpserv.Error{Status: http.StatusInternalServerError, Code: "internal_error", Desc: "Something went wrong"},
			body:      []byte(`{"sender":"andy@example.com","text":"john@example.com"}`),
			errDb:     errors.New("cannot find user"),
			resultDb:  []string{},
		},
		"failFromDBInvalidUser": {
			expStatus: http.StatusNotFound,
			expErr:    &httpserv.Error{Status: http.StatusNotFound, Code: "invalid_email", Desc: "Provided email does not exist"},
			body:      []byte(`{"sender":"andy@example.com","text":"john@example.com"}`),
			errDb:     errors.New("not found"),
			resultDb:  []string{},
		},
		"success": {
			expStatus: http.StatusOK,
			expErr:    nil,
			body:      []byte(`{"sender":"andy@example.com","text":"john@example.com"}`),
			errDb:     nil,
			resultDb:  []string{"test@gmail.com, test2@gmail.com"},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/_/update-receiver", bytes.NewReader(tc.body))
			w := httptest.NewRecorder()

			mockRelaCtrl := relationship.MockApiRestController{}
			mockUserCtrl := user.MockApiRestController{}
			mockRelaCtrl.On("UpdateReceiver", mock.Anything, mock.Anything, mock.Anything).Return(tc.resultDb, tc.errDb)

			// When:
			New(&mockUserCtrl, &mockRelaCtrl).UpdateReceiver().ServeHTTP(w, r)
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
