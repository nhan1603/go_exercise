package user

import (
	"gobase/api/internal/controller/relationship"
	"gobase/api/internal/controller/user"
)

// ApiHandler is the web handler for this pkg
type ApiHandler struct {
	userCtrl user.ApiRestController
	relaCtrl relationship.ApiRestController
}

// New instantiates a new Handler and returns it
func New(userCtrl user.ApiRestController, relaCtrl relationship.ApiRestController) ApiHandler {
	return ApiHandler{userCtrl: userCtrl, relaCtrl: relaCtrl}
}
