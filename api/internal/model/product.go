package model

import (
	"time"
)

// ProductStatus represents the status of the product
type ProductStatus string

const (
	// ProductStatusActive means the product is active
	ProductStatusActive ProductStatus = "ACTIVE"
	// ProductStatusInactive means the product is inactive
	ProductStatusInactive ProductStatus = "INACTIVE"
	// ProductStatusDeleted means the product is deleted. This is for archival only
	ProductStatusDeleted ProductStatus = "DELETED"
)

// AllProductStatus is a list of all ProductStatus
var AllProductStatus = []ProductStatus{ProductStatusActive, ProductStatusInactive, ProductStatusDeleted}

// String converts to string value
func (p ProductStatus) String() string {
	return string(p)
}

// IsValid checks if plan status is valid
func (p ProductStatus) IsValid() bool {
	return p == ProductStatusActive || p == ProductStatusInactive || p == ProductStatusDeleted
}

// Product represents the product to be sold
type Product struct {
	ID          int
	ExternalID  string
	Name        string
	Description string
	Status      ProductStatus
	Price       int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
