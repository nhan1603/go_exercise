package router

import (
	"context"
	"gobase/api/internal/controller/relationship"
	"gobase/api/internal/controller/user"
	relaRest "gobase/api/internal/handler/rest/relationship"
	userRest "gobase/api/internal/handler/rest/user"
)

// New creates and returns a new Router instance
func New(
	ctx context.Context,
	corsOrigin []string,
	isGQLIntrospectionOn bool,
	userCtrl user.ApiRestController,
	relaCtrl relationship.ApiRestController,
) Router {
	return Router{
		ctx:                  ctx,
		corsOrigins:          corsOrigin,
		isGQLIntrospectionOn: isGQLIntrospectionOn,
		userRESTHandler:      userRest.New(userCtrl, relaCtrl),
		relaRESTHandler:      relaRest.New(userCtrl, relaCtrl),
	}
}
