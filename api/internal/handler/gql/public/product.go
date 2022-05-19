package public

import (
	"context"
	"net/http"
	"strings"
	"time"

	"gobase/api/internal/controller/products"
	"gobase/api/internal/handler/gql/mod"
	"gobase/api/internal/model"
	"gobase/api/pkg/httpserv"
)

// CreateProductInput represents the input for create product
type CreateProductInput struct {
	Name        string
	Description string
	Status      model.ProductStatus
	Price       int64
}

// UpdateProductInput represents the input for update product
type UpdateProductInput struct {
	Name        string
	Description string
	Status      model.ProductStatus
	Price       int64
}

// CreateProduct creates a product
func (r *mutationResolver) CreateProduct(ctx context.Context, input CreateProductInput) (*mod.Product, error) {
	ctrlInp := products.CreateInput{
		Name:  strings.TrimSpace(input.Name),
		Desc:  strings.TrimSpace(input.Description),
		Price: input.Price,
	}
	if ctrlInp.Name == "" {
		return nil, &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_name", Desc: "Name given is invalid"}
	}
	if ctrlInp.Desc == "" {
		return nil, &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_description", Desc: "Description given is invalid"}
	}
	if ctrlInp.Price <= 0 {
		return nil, &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_price", Desc: "Price given is invalid"}
	}

	p, err := r.productsCtrl.Create(ctx, ctrlInp)
	if err != nil {
		return nil, convertCtrlErr(err)
	}

	return mod.NewProduct(p), nil
}

// GetProducts retrieves products
func (r *queryResolver) GetProducts(ctx context.Context) ([]*mod.Product, error) {
	// TODO: Add validation example
	// TODO: Add filtering example

	slice, err := r.productsCtrl.List(ctx)
	if err != nil {
		return nil, err
	}

	return mod.NewProducts(slice), nil
}

// UpdateProduct updates a product
func (r *mutationResolver) UpdateProduct(_ context.Context, externalID string, input UpdateProductInput) (*mod.Product, error) {
	if strings.TrimSpace(externalID) == "" {
		return nil, &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_ext_id", Desc: "External ID given is invalid"}
	}
	if strings.TrimSpace(input.Name) == "" {
		return nil, &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_name", Desc: "Name given is invalid"}
	}
	if strings.TrimSpace(input.Description) == "" {
		return nil, &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_description", Desc: "Description given is invalid"}
	}
	if input.Price <= 0 {
		return nil, &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_price", Desc: "Price given is invalid"}
	}

	return mod.NewProduct(model.Product{
		ID:          1,
		ExternalID:  "ext",
		Name:        "name",
		Description: "desc",
		Status:      model.ProductStatusActive,
		Price:       12345,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}), nil
}
