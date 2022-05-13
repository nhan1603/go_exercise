package health

import "gobase/api/internal/controller/system"

// Handler is the web handler for this pkg
type Handler struct {
	systemCtrl system.Controller
}

// New instantiates a new Handler and returns it
func New(systemCtrl system.Controller) Handler {
	return Handler{systemCtrl: systemCtrl}
}
