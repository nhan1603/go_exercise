package inventory

import (
	"errors"
)

var (
	// ErrNotFound means the item was not found
	ErrNotFound = errors.New("not found")
	// ErrUnexpectedRowsFound means there is a mismatch with expected vs actual no. of rows
	ErrUnexpectedRowsFound = errors.New("unexpected rows found")
	// ErrCacheMiss means there was a cache miss
	ErrCacheMiss = errors.New("cache miss")
)
