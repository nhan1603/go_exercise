package utils

import (
	"database/sql"
	"github.com/pkg/errors"
)

func MergeErr(errs ...error) error {
	var returnErr error = nil
	for _, err := range errs {
		if err != nil {
			returnErr = errors.Wrap(returnErr, err.Error())
		}
	}
	return returnErr
}

func MergeErrDB(errs ...error) error {
	var returnErr error = nil
	for _, err := range errs {
		if errors.Cause(err) == sql.ErrNoRows {
			return sql.ErrNoRows
		}
		if err != nil {
			returnErr = errors.Wrap(returnErr, err.Error())
		}
	}
	return returnErr
}
