package health

import (
	"encoding/json"
	"gobase/api/pkg/httpserv"
	"net/http"
)

// Block creates a block relationship between two emails
func (h Handler) Block() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req SubscribeInput

		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		errBlock := h.systemCtrl.Block(r.Context(), req.Requestor, req.Target)

		if errBlock != nil {
			return httpserv.Error{Status: http.StatusBadRequest, Code: "error request", Desc: errBlock.Error()}
		}

		if errBlock == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.Response{Success: true})
		}

		return errBlock
	})
}
