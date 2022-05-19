package generator

import (
	pkgerrors "github.com/pkg/errors"
	"gobase/api/pkg/snowflake"
)

// Snowflake generators per table.
var (
	// ProductIDSNF the snowflake generator for Product table's ID in DB
	ProductIDSNF *snowflake.Generator
)

// InitSnowflakeGenerators initializes all the snowflake generators
func InitSnowflakeGenerators() error {
	var err error

	if ProductIDSNF == nil {
		ProductIDSNF, err = snowflake.New()
		if err != nil {
			return pkgerrors.WithStack(err)
		}
	}

	return nil
}
