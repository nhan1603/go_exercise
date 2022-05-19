package health

import (
	"context"
	"encoding/json"
	"errors"
	pkgerrors "github.com/pkg/errors"
	"gobase/api/pkg/httpserv"
	"net/http"
)

type reqBody struct {
	Friends []string `json:"friends"`
}

// AddFriend add friend for two user
func (h Handler) AddFriend() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req reqBody

		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}
		if len(req.Friends) != 2 {
			return pkgerrors.WithStack(errors.New("invalid array length"))
		}

		errAdd := h.systemCtrl.AddFriend(r.Context(), req.Friends[0], req.Friends[1])

		if errors.Is(errAdd, context.Canceled) {
			return nil
		}

		if errAdd == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.CustomResponse{Success: true})
		}

		return errAdd
	})
}
