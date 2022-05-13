package public

import (
	"net/http"

	"gobase/api/internal/controller/products"
	"gobase/api/pkg/httpserv"
)

var (
	// 4xx:
	errUnauthorized     = &httpserv.Error{Status: http.StatusForbidden, Code: "unauthorized", Desc: "Unauthorized"}
	errProductNotFound  = &httpserv.Error{Status: http.StatusBadRequest, Code: "product_not_found", Desc: "Product not found"}
	errProductNotActive = &httpserv.Error{Status: http.StatusBadRequest, Code: "product_not_active", Desc: "Product not active"}
)

func convertCtrlErr(err error) error {
	switch err {
	case products.ErrNotFound:
		return errProductNotFound
	case products.ErrNotActive:
		return errProductNotActive
	default:
		return err
	}
}
