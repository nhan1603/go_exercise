package relationship

import (
	"encoding/json"
	"gobase/api/internal/model"
	"gobase/api/pkg/httpserv"
	"net/http"
)

// MakeFriendInput struct for parsing body from request for making friends
type MakeFriendInput struct {
	Friends []string `json:"friends"`
}

// AddFriend add friend for two user
func (h ApiHandler) AddFriend() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req MakeFriendInput

		if err := decoder.Decode(&req); err != nil {
			return &httpserv.Error{Status: http.StatusBadRequest, Code: "request_body_error", Desc: "Invalid request body"}
		}

		if err := req.validate(); err != nil {
			return err
		}

		if err := h.relaCtrl.AddFriend(r.Context(), model.MakeRelationship{
			FromFriend: req.Friends[0],
			ToFriend:   req.Friends[1],
		}); err != nil {
			if err.Error() == "not found" {
				return &httpserv.Error{Status: http.StatusNotFound, Code: "invalid_email", Desc: err.Error()}
			}
			return err
		}

		httpserv.RespondJSON(r.Context(), w, httpserv.Response{Success: true})

		return nil
	})
}

func (i MakeFriendInput) validate() error {
	if len(i.Friends) != 2 {
		return &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "User has provided wrong amount of emails"}
	}

	if i.Friends[0] == i.Friends[1] {
		return &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "User has provided identical emails"}
	}

	if i.Friends[0] == "" || i.Friends[1] == "" {
		return &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "User has provided empty email"}
	}

	return nil
}
