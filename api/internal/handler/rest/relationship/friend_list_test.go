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

func TestApiHandler_FindFriendList(t *testing.T) {
	tcs := map[string]struct {
		expErr    *httpserv.Error
		errDb     error
		body      []byte
		resultDb  []string
		expStatus int
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
			body:      []byte(`{"email":""}`),
			errDb:     nil,
			resultDb:  []string{},
		},
		"failFromDB": {
			expStatus: http.StatusInternalServerError,
			expErr:    &httpserv.Error{Status: http.StatusInternalServerError, Code: "internal_error", Desc: "cannot find user"},
			body:      []byte(`{"email":"andy@example.com"}`),
			errDb:     errors.New("cannot find user"),
			resultDb:  []string{},
		},
		"failFromDBInvalidUser": {
			expStatus: http.StatusNotFound,
			expErr:    &httpserv.Error{Status: http.StatusNotFound, Code: "invalid_email", Desc: "Provided email does not exist"},
			body:      []byte(`{"email":"andy@example.com"}`),
			errDb:     errors.New("not found"),
		},
		"success": {
			expStatus: http.StatusOK,
			expErr:    nil,
			body:      []byte(`{"email":"andy@example.com"}`),
			errDb:     nil,
			resultDb:  []string{"test@gmail.com, test2@gmail.com"},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/api/friend-list", bytes.NewReader(tc.body))
			w := httptest.NewRecorder()

			mockRelaCtrl := relationship.MockApiRestController{}
			mockUserCtrl := user.MockApiRestController{}
			mockRelaCtrl.On("FindFriendList", mock.Anything, mock.Anything).Return(tc.resultDb, tc.errDb)

			// When:
			New(&mockUserCtrl, &mockRelaCtrl).FindFriendList().ServeHTTP(w, r)
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

func TestApiHandler_FindCommonFriend(t *testing.T) {
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
		"duplicate_email": {
			expStatus: http.StatusBadRequest,
			expErr:    &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "User has provided identical emails"},
			body:      []byte(`{"friends":["andy@example.com","andy@example.com"]}`),
			errDb:     nil,
			resultDb:  []string{},
		},
		"empty_email": {
			expStatus: http.StatusBadRequest,
			expErr:    &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "User has provided empty email"},
			body:      []byte(`{"friends":["","andy@example.com"]}`),
			errDb:     nil,
			resultDb:  []string{},
		},
		"invalid_range": {
			expStatus: http.StatusBadRequest,
			expErr:    &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "User has provided wrong amount of emails"},
			body:      []byte(`{"friends":["","andy@example.com","andy@example.com"]}`),
			errDb:     nil,
			resultDb:  []string{},
		},
		"failFromDB": {
			expStatus: http.StatusInternalServerError,
			expErr:    &httpserv.Error{Status: http.StatusInternalServerError, Code: "internal_error", Desc: "cannot create new friendship"},
			body:      []byte(`{"friends":["andy@example.com","john@example.com"]}`),
			errDb:     errors.New("cannot create new friendship"),
			resultDb:  []string{},
		},
		"failFromDBInvalidUser": {
			expStatus: http.StatusNotFound,
			expErr:    &httpserv.Error{Status: http.StatusNotFound, Code: "invalid_email", Desc: "Provided email does not exist"},
			body:      []byte(`{"friends":["andy@example.com","john@example.com"]}`),
			errDb:     errors.New("not found"),
			resultDb:  []string{},
		},
		"success": {
			expStatus: http.StatusOK,
			expErr:    nil,
			body:      []byte(`{"friends":["andy@example.com","john@example.com"]}`),
			errDb:     nil,
			resultDb:  []string{"test@gmail.com, test2@gmail.com"},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/api/common-friend", bytes.NewReader(tc.body))
			w := httptest.NewRecorder()

			mockRelaCtrl := relationship.MockApiRestController{}
			mockUserCtrl := user.MockApiRestController{}
			mockRelaCtrl.On("FindCommonFriends", mock.Anything, mock.Anything).Return(tc.resultDb, tc.errDb)

			// When:
			New(&mockUserCtrl, &mockRelaCtrl).FindCommonFriend().ServeHTTP(w, r)
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
