package products

import (
	"errors"
)

var (
	// ErrNotActive means the product is not active
	ErrNotActive = errors.New("not active")
	// ErrNotFound means the product was not found
	ErrNotFound = errors.New("not found")
	// ErrUnexpectedProducts means there is mismatch in no. of products expected
	ErrUnexpectedProducts = errors.New("received unexpected num of products")
)
