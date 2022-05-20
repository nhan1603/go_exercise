package health

import (
	"context"
	"encoding/json"
	"errors"
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

		errAdd := h.systemCtrl.Block(r.Context(), req.Requestor, req.Target)

		if errors.Is(errAdd, context.Canceled) {
			return nil
		}

		if errAdd == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.Response{Success: true})
		}

		return errAdd
	})
}
