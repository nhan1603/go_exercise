package relationship

import (
	"encoding/json"
	"gobase/api/internal/handler/rest/user"
	"gobase/api/internal/model"
	"gobase/api/pkg/httpserv"
	"net/http"
)

// FindFriendList create a user from email
func (h ApiHandler) FindFriendList() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req user.UserInfoInput

		if err := decoder.Decode(&req); err != nil {
			return &httpserv.Error{Status: http.StatusBadRequest, Code: "request_body_error", Desc: "Invalid request body"}
		}

		if err := req.Validate(); err != nil {
			return err
		}

		listFriend, err := h.relaCtrl.FindFriendList(r.Context(), req.Email)

		if err == nil {
			httpserv.RespondJSON(r.Context(), w, httpserv.FriendListResponse{
				Success: true,
				Friends: listFriend,
				Count:   len(listFriend),
			})
		} else {
			if err.Error() == "not found" {
				return &httpserv.Error{Status: http.StatusNotFound, Code: "invalid_email", Desc: err.Error()}
			}
		}

		return err
	})
}

// FindCommonFriend create a user from email
func (h ApiHandler) FindCommonFriend() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var req MakeFriendInput

		if err := decoder.Decode(&req); err != nil {
			return &httpserv.Error{Status: http.StatusBadRequest, Code: "request_body_error", Desc: "Invalid request body"}
		}

		if err := req.validate(); err != nil {
			return err
		}

		commonFriend, err := h.relaCtrl.FindCommonFriends(r.Context(),
			model.CommonFriend{
				FirstUser:  req.Friends[0],
				SecondUser: req.Friends[1],
			})

		if err != nil {
			if err.Error() == "not found" {
				return &httpserv.Error{Status: http.StatusNotFound, Code: "invalid_email", Desc: err.Error()}
			}
			return err
		}

		httpserv.RespondJSON(r.Context(), w, httpserv.FriendListResponse{
			Success: true,
			Friends: commonFriend,
			Count:   len(commonFriend),
		})

		return nil
	})
}
