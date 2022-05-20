package health

import (
	"context"
	"encoding/json"
	"errors"
	"gobase/api/pkg/httpserv"
	"net/http"
)

type updateStructDTO struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// UpdateReceiver returns a list of emails that will receive message from certain user
func (h Handler) UpdateReceiver() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req updateStructDTO

		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		listReceiver, errFind := h.systemCtrl.UpdateReceiver(r.Context(), req.Sender, req.Text)

		if errors.Is(errFind, context.Canceled) {
			return nil
		}

		if errFind == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.UpdateReceiveResponse{
				Success:    true,
				Recipients: listReceiver,
			})
		}

		return errFind
	})
}
