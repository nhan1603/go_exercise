//go:generate rm -rf vendor
//go:generate go run github.com/99designs/gqlgen generate

package public

import (
	"github.com/99designs/gqlgen/graphql"
	"gobase/api/internal/controller/products"
)

// NewSchema returns the GraphQL schema
func NewSchema(
	productsCtrl products.Controller,
) graphql.ExecutableSchema {
	cfg := Config{
		Resolvers: &resolver{
			productsCtrl: productsCtrl,
		},
	}

	return NewExecutableSchema(cfg)
}

type resolver struct {
	productsCtrl products.Controller
}

// Query returns the QueryResolver
func (r *resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// Mutation returns the MutationResolver
func (r *resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct {
	*resolver
}

type mutationResolver struct {
	*resolver
}
