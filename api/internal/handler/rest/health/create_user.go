package health

import (
	"context"
	"encoding/json"
	"errors"
	"gobase/api/pkg/httpserv"
	"net/http"
)

type userInfo struct {
	Email string `json:"email"`
}

// CreateUser create a user from email
func (h Handler) CreateUser() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req userInfo

		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		errAdd := h.systemCtrl.CreateUser(r.Context(), req.Email)

		if errors.Is(errAdd, context.Canceled) {
			return nil
		}

		if errAdd == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.CustomResponse{Success: true})
		}

		return errAdd
	})
}
