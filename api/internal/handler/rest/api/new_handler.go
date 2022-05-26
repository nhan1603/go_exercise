package api

import (
	"gobase/api/internal/controller/rest"
)

// ApiHandler is the web handler for this pkg
type ApiHandler struct {
	systemCtrl rest.ApiRestController
}

// New instantiates a new Handler and returns it
func New(systemCtrl rest.ApiRestController) ApiHandler {
	return ApiHandler{systemCtrl: systemCtrl}
}
