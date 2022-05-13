package mod

import (
	"time"

	"gobase/api/internal/model"
)

// Product is the gql representation of Product
type Product struct {
	m model.Product
}

// ExternalID returns the external ID
func (s Product) ExternalID() string {
	return s.m.ExternalID
}

// Name returns the name
func (s Product) Name() string {
	return s.m.Name
}

// Description returns the description
func (s Product) Description() string {
	return s.m.Description
}

// Status returns the status
func (s Product) Status() model.ProductStatus {
	return s.m.Status
}

// Price returns the price
func (s Product) Price() int64 {
	return s.m.Price
}

// CreatedAt returns the created at
func (s Product) CreatedAt() time.Time {
	return s.m.CreatedAt
}

// UpdatedAt returns the updated at
func (s Product) UpdatedAt() time.Time {
	return s.m.UpdatedAt
}

// NewProduct converts model.Product into Product
func NewProduct(m model.Product) *Product {
	return &Product{m: m}
}

// NewProducts converts slice of model.Product into slice of Product
func NewProducts(slice []model.Product) []*Product {
	result := make([]*Product, len(slice))
	for i, m := range slice {
		result[i] = &Product{m: m}
	}
	return result
}
