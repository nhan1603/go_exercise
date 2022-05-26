package api

import (
	"encoding/json"
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
			// return bad request
			return err
		}

		if err = h.systemCtrl.Subscribe(r.Context(), req.Requestor, req.Target); err != nil {
			return err
		}

		httpserv.RespondJSON(r.Context(), w, httpserv.Response{Success: true})

		return nil
	})
}
