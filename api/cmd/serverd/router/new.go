package router

import (
	"context"
	"gobase/api/internal/controller/rest"
	"gobase/api/internal/handler/rest/api"

	"gobase/api/internal/controller/products"
	"gobase/api/internal/controller/system"
	"gobase/api/internal/handler/rest/health"
)

// New creates and returns a new Router instance
func New(
	ctx context.Context,
	corsOrigin []string,
	isGQLIntrospectionOn bool,
	systemCtrl system.Controller,
	apiCtrl rest.ApiRestController,
	productCtrl products.Controller,
) Router {
	return Router{
		ctx:                  ctx,
		corsOrigins:          corsOrigin,
		isGQLIntrospectionOn: isGQLIntrospectionOn,
		healthRESTHandler:    health.New(systemCtrl),
		apiRESTHandler:       api.New(apiCtrl),
		productCtrl:          productCtrl,
	}
}
