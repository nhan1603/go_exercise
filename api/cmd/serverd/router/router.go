package router

import (
	"context"
	"gobase/api/internal/controller/products"
	"gobase/api/internal/handler/gql/public"
	"gobase/api/pkg/httpserv/gql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gobase/api/internal/handler/rest/health"
	"gobase/api/pkg/httpserv"
)

// Router defines the routes & handlers of the app
type Router struct {
	ctx                  context.Context
	corsOrigins          []string
	isGQLIntrospectionOn bool
	healthRESTHandler    health.Handler
	productCtrl          products.Controller
}

// Handler returns the Handler for use by the server
func (rtr Router) Handler() http.Handler {
	return httpserv.Handler(
		rtr.healthRESTHandler.CheckReadiness(),
		rtr.healthRESTHandler.Liveness(),
		rtr.healthRESTHandler.AddFriend(),
		rtr.healthRESTHandler.CreateUser(),
		rtr.healthRESTHandler.FindFriendList(),
		rtr.healthRESTHandler.FindCommonFriend(),
		rtr.healthRESTHandler.Subscribe(),
		rtr.healthRESTHandler.Block(),
		rtr.healthRESTHandler.UpdateReceiver(),
		rtr.routes)
}

func (rtr Router) routes(r chi.Router) {
	r.Group(rtr.public)
}

func (rtr Router) public(r chi.Router) {
	const prefix = "/gateway/public"

	r.Handle(prefix+"/graphql", gql.Handler(public.NewSchema(
		rtr.productCtrl),
		rtr.isGQLIntrospectionOn))
}
