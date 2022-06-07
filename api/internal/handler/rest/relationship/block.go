package relationship

import (
	"encoding/json"
	"gobase/api/internal/model"
	"gobase/api/pkg/httpserv"
	"net/http"
)

// Block creates a block relationship between two emails
func (h ApiHandler) Block() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req SubscribeInput

		if err := decoder.Decode(&req); err != nil {
			return &httpserv.Error{Status: http.StatusBadRequest, Code: "request_body_error", Desc: "Invalid request body"}
		}

		if err := req.validate(); err != nil {
			return err
		}

		err := h.relaCtrl.Block(r.Context(), model.MakeRelationship{
			FromFriend: req.Requestor,
			ToFriend:   req.Target,
		})

		if err != nil {
			if err.Error() == "not found" {
				return &httpserv.Error{Status: http.StatusNotFound, Code: "invalid_email", Desc: err.Error()}
			}
			return err
		}

		httpserv.RespondJSON(r.Context(), w, httpserv.Response{Success: true})
		return nil
	})
}
