package api

import (
	"encoding/json"
	"gobase/api/internal/model"
	"gobase/api/pkg/httpserv"
	"net/http"
)

type MakeFriendInput struct {
	Friends []string `json:"friends"`
}

// AddFriend add friend for two user
func (h ApiHandler) AddFriend() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req MakeFriendInput

		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		if err = req.validate(); err != nil {
			return err
		}

		errAdd := h.systemCtrl.AddFriend(r.Context(), model.MakeFriend{
			FromFriend: req.Friends[0],
			ToFriend:   req.Friends[1],
		})

		if errAdd != nil {
			return &httpserv.Error{Status: http.StatusBadRequest, Code: "error request", Desc: errAdd.Error()}
		}

		if errAdd == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.Response{Success: true})
		}

		return errAdd
	})
}

func (i MakeFriendInput) validate() error {
	if len(i.Friends) != 2 {
		return &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "Friendly message"}
	}

	if i.Friends[0] == i.Friends[1] {
		return &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "Friendly message"}
	}

	if i.Friends[0] == "" || i.Friends[1] == "" {
		return &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "Friendly message"}
	}

	return nil
}
