package api

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gobase/api/internal/controller/rest"
	"gobase/api/pkg/httpserv"
	"net/http"
	"net/http/httptest"
	"testing"
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
			expErr:    &httpserv.Error{Status: http.StatusBadRequest, Code: "error in request body", Desc: "EOF"},
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

			mockController := rest.MockApiRestController{}
			mockController.On("UpdateReceiver", mock.Anything, mock.Anything, mock.Anything).Return(tc.resultDb, tc.errDb)

			// When:
			New(&mockController).UpdateReceiver().ServeHTTP(w, r)
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
