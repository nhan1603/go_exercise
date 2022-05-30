package api

import (
	"encoding/json"
	"gobase/api/internal/model"
	"gobase/api/pkg/httpserv"
	"net/http"
)

// SubscribeInput struct for parsing body from request for subscribe
type SubscribeInput struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

// Subscribe creates a subscription between two emails
func (h ApiHandler) Subscribe() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req SubscribeInput

		err := decoder.Decode(&req)
		if err != nil {
			return &httpserv.Error{Status: http.StatusBadRequest, Code: "error in request body", Desc: err.Error()}
		}

		if err = req.validate(); err != nil {
			return err
		}

		if err = h.systemCtrl.Subscribe(r.Context(), model.MakeRelationship{FromFriend: req.Requestor, ToFriend: req.Target}); err != nil {
			return err
		}

		httpserv.RespondJSON(r.Context(), w, httpserv.Response{Success: true})

		return nil
	})
}

func (i SubscribeInput) validate() error {
	if i.Requestor == i.Target {
		return &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "User has provided identical emails"}
	}

	if i.Requestor == "" || i.Target == "" {
		return &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "User has provided empty email"}
	}

	return nil
}
