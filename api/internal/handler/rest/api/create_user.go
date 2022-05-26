package api

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
			panic(err)
		}

		_, errCreate := h.systemCtrl.CreateUser(r.Context(), req.Email)

		if errCreate != nil {
			return httpserv.Error{Status: http.StatusBadRequest, Code: "error request", Desc: errCreate.Error()}
		}

		if errCreate == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.Response{Success: true})
		}

		return errCreate
	})
}
