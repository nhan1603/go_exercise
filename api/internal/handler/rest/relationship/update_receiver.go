package relationship

import (
	"encoding/json"
	"gobase/api/internal/model"
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
			return &httpserv.Error{Status: http.StatusBadRequest, Code: "request_body_error", Desc: "Invalid request body"}
		}

		if err = req.validate(); err != nil {
			return err
		}

		listReceiver, errFind := h.relaCtrl.UpdateReceiver(r.Context(), model.UpdateInfo{Sender: req.Sender, Message: req.Text})

		if errFind == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.UpdateReceiveResponse{
				Success:    true,
				Recipients: listReceiver,
			})
		} else {
			if errFind.Error() == "not found" {
				return &httpserv.Error{Status: http.StatusNotFound, Code: "invalid_email", Desc: errFind.Error()}
			}
			return errFind
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
