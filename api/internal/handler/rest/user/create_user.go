package user

import (
	"encoding/json"
	"gobase/api/pkg/httpserv"
	"net/http"
)

type UserInfoInput struct {
	Email string `json:"email"`
}

// CreateUser create a user from email
func (h ApiHandler) CreateUser() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req UserInfoInput

		err := decoder.Decode(&req)
		if err != nil {
			return &httpserv.Error{Status: http.StatusBadRequest, Code: "error in request body", Desc: err.Error()}
		}

		if err = req.Validate(); err != nil {
			return err
		}

		_, errCreate := h.userCtrl.CreateUser(r.Context(), req.Email)

		if errCreate != nil {
			return errCreate
		}

		if errCreate == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.Response{Success: true})
		}

		return errCreate
	})
}

func (i UserInfoInput) Validate() error {
	if i.Email == "" {
		return &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "User has provided empty email"}
	}
	return nil
}
