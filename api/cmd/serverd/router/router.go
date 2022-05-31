package router

import (
	"context"
	"gobase/api/internal/handler/rest/relationship"
	"gobase/api/internal/handler/rest/user"
	"net/http"

	"gobase/api/pkg/httpserv"
)

// Router defines the routes & handlers of the app
type Router struct {
	ctx                  context.Context
	corsOrigins          []string
	isGQLIntrospectionOn bool
	userRESTHandler      user.ApiHandler
	relaRESTHandler      relationship.ApiHandler
}

// Handler returns the Handler for use by the server
func (rtr Router) Handler() http.Handler {
	return httpserv.Handler(
		rtr.relaRESTHandler.AddFriend(),
		rtr.userRESTHandler.CreateUser(),
		rtr.relaRESTHandler.FindFriendList(),
		rtr.relaRESTHandler.FindCommonFriend(),
		rtr.relaRESTHandler.Subscribe(),
		rtr.relaRESTHandler.Block(),
		rtr.relaRESTHandler.UpdateReceiver(),
	)
}
