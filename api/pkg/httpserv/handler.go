package httpserv

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Handler(
	addFriend, createUser, getFriendList, getCommonFriend,
	subscribe, block, updateReceiver http.HandlerFunc) http.Handler {
	r := chi.NewRouter()
	// TODO: add middleware here
	r.Use(
		// render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger, // log relationship request calls
		// middleware.DefaultCompress, // compress results, mostly gzipping assets and json
		middleware.StripSlashes, // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,    // recover from panics without crashing server
	)

	r.Post("/api/add-friend", addFriend)

	r.Post("/api/create-user", createUser)

	r.Post("/api/friend-list", getFriendList)

	r.Post("/api/common-friend", getCommonFriend)

	r.Post("/api/subscribe", subscribe)

	r.Post("/api/block", block)

	r.Post("/api/update-receiver", updateReceiver)

	return r
}
