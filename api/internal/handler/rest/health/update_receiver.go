package health

import (
	"encoding/json"
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

		if errFind == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.UpdateReceiveResponse{
				Success:    true,
				Recipients: listReceiver,
			})
		} else {
			return httpserv.Error{Status: http.StatusBadRequest, Code: "error request", Desc: errFind.Error()}
		}

		return errFind
	})
}