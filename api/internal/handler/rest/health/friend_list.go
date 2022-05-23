package health

import (
	"encoding/json"
	"errors"
	pkgerrors "github.com/pkg/errors"
	"gobase/api/pkg/httpserv"
	"net/http"
)

// FindFriendList create a user from email
func (h Handler) FindFriendList() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req userInfo

		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		listFriend, errFind := h.systemCtrl.FindFriendList(r.Context(), req.Email)

		if errFind == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.FriendListResponse{
				Success: true,
				Friends: listFriend,
				Count:   len(listFriend),
			})
		} else {
			return httpserv.Error{Status: http.StatusBadRequest, Code: "error request", Desc: errFind.Error()}
		}

		return errFind
	})
}

// FindCommonFriend create a user from email
func (h Handler) FindCommonFriend() http.HandlerFunc {
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

		commonFriend, errFind := h.systemCtrl.FindCommonFriends(r.Context(), req.Friends[0], req.Friends[1])

		if errFind != nil {
			return httpserv.Error{Status: http.StatusBadRequest, Code: "error request", Desc: errFind.Error()}
		}

		if errFind == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.FriendListResponse{
				Success: true,
				Friends: commonFriend,
				Count:   len(commonFriend),
			})
		}

		return errFind
	})
}
