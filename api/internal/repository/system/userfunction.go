package system

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"gobase/api/internal/repository/orm"
)

// FindUserByEmail will find the user entity with the corresponding email
func (i impl) FindUserByEmail(ctx context.Context, email string) (*orm.User, error) {

	userObj := &orm.User{}

	sel := "*"

	query := fmt.Sprintf(
		"select %s from \"user\" where \"email\"=$1", sel,
	)

	q := queries.Raw(query, email)

	err := q.Bind(ctx, i.dbConn, userObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from user")
	}

	return userObj, nil
}

// CreateUser creates a new user entity with the corresponding email
func (i impl) CreateUser(ctx context.Context, email string) (int, error) {
	userEntity := orm.User{
		Email: email,
	}

	return userEntity.ID, pkgerrors.WithStack(userEntity.Insert(ctx, i.dbConn, boil.Infer()))
}
