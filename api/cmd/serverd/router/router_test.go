package router

import (
	"context"
	"net/http"
	"sort"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestRouter_Handler(t *testing.T) {
	// Given:
	r := New(
		context.Background(),
		[]string{},
		true,
		nil,
		nil,
		nil,
	)

	expectedRoutes := []string{
		"GET /_/live", "GET /_/ready",

		"CONNECT /gateway/public/graphql",
		"DELETE /gateway/public/graphql",
		"OPTIONS /gateway/public/graphql",
		"GET /gateway/public/graphql",
		"HEAD /gateway/public/graphql",
		"PATCH /gateway/public/graphql",
		"POST /gateway/public/graphql",
		"PUT /gateway/public/graphql",
		"TRACE /gateway/public/graphql",
	}
	sort.Strings(expectedRoutes)
	var routesFound []string

	handler, ok := r.Handler().(chi.Router)
	if !ok {
		require.FailNow(t, "handler is not a chi router")
	}

	// When:
	err := chi.Walk(
		handler,
		func(method string, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
			routesFound = append(routesFound, method+" "+route) // TODO: Add middleware stack check also if possible.
			return nil
		},
	)
	require.NoError(t, err)
	sort.Strings(routesFound)

	// Then:
	require.EqualValues(t, expectedRoutes, routesFound)
}
