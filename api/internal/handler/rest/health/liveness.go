package health

import (
	"net/http"

	"gobase/api/pkg/httpserv"
)

// Liveness is a default route to report that app is live
func (h Handler) Liveness() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		httpserv.RespondJSON(r.Context(), w, httpserv.CustomResponse{Success: true})

		return nil
	})
}
