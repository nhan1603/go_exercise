package relationship

import (
	"encoding/json"
	"gobase/api/pkg/httpserv"
	"net/http"
)

type UpdateReceiveInput struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// UpdateReceiver returns a list of emails that will receive message from certain user
func (h ApiHandler) UpdateReceiver() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req UpdateReceiveInput

		err := decoder.Decode(&req)
		if err != nil {
			return &httpserv.Error{Status: http.StatusBadRequest, Code: "error in request body", Desc: err.Error()}
		}

		if err = req.validate(); err != nil {
			return err
		}

		listReceiver, errFind := h.relaCtrl.UpdateReceiver(r.Context(), req.Sender, req.Text)

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

func (i UpdateReceiveInput) validate() error {
	if i.Sender == "" {
		return &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "User has provided empty email"}
	}

	return nil
}
