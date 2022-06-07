package user

import (
	"encoding/json"
	"gobase/api/pkg/httpserv"
	"net/http"
)

// UserInfoInput represents the parsing data from request body creating a new user
type UserInfoInput struct {
	Email string `json:"email"`
}

// CreateUser create a user from email
func (h ApiHandler) CreateUser() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req UserInfoInput

		if err := decoder.Decode(&req); err != nil {
			return &httpserv.Error{Status: http.StatusBadRequest, Code: "request_body_error", Desc: "Invalid request body"}
		}

		if err := req.Validate(); err != nil {
			return err
		}

		_, err := h.userCtrl.CreateUser(r.Context(), req.Email)

		if err != nil {
			if err.Error() == "Existed email input." {
				return &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_email", Desc: err.Error()}
			}
			return err

		}

		if err == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.Response{Success: true})
		}

		return err
	})
}

func (i UserInfoInput) Validate() error {
	if i.Email == "" {
		return &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "User has provided empty email"}
	}
	return nil
}
