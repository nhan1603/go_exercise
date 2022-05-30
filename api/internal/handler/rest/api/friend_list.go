package api

import (
	"encoding/json"
	"gobase/api/internal/model"
	"gobase/api/pkg/httpserv"
	"net/http"
)

// FindFriendList create a user from email
func (h ApiHandler) FindFriendList() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req UserInfoInput

		err := decoder.Decode(&req)
		if err != nil {
			return &httpserv.Error{Status: http.StatusBadRequest, Code: "error in request body", Desc: err.Error()}
		}

		if err = req.validate(); err != nil {
			return err
		}

		listFriend, errFind := h.systemCtrl.FindFriendList(r.Context(), req.Email)

		if errFind == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.FriendListResponse{
				Success: true,
				Friends: listFriend,
				Count:   len(listFriend),
			})
		} else {
			return errFind
		}

		return errFind
	})
}

// FindCommonFriend create a user from email
func (h ApiHandler) FindCommonFriend() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req MakeFriendInput

		err := decoder.Decode(&req)
		if err != nil {
			return &httpserv.Error{Status: http.StatusBadRequest, Code: "error in request body", Desc: err.Error()}
		}

		if err = req.validate(); err != nil {
			return err
		}

		commonFriend, errFind := h.systemCtrl.FindCommonFriends(r.Context(),
			model.CommonFriend{
				FirstUser:  req.Friends[0],
				SecondUser: req.Friends[1],
			})

		if errFind != nil {
			return errFind
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
