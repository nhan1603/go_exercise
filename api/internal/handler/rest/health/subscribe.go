package health

import (
	"context"
	"encoding/json"
	"errors"
	"gobase/api/pkg/httpserv"
	"net/http"
)

type reqStructDTO struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

// Subscribe creates a subscription between two emails
func (h Handler) Subscribe() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req reqStructDTO

		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		errAdd := h.systemCtrl.Subscribe(r.Context(), req.Requestor, req.Target)

		if errors.Is(errAdd, context.Canceled) {
			return nil
		}

		if errAdd == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.CustomResponse{Success: true})
		}

		return errAdd
	})
}
