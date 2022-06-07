package user

import (
	"context"
	"database/sql"
	"errors"
	"gobase/api/internal/repository/orm"

	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// FindUserByEmail will find the user entity with the corresponding email
func (i impl) FindUserByEmail(ctx context.Context, email string) (orm.User, error) {
	u, err := orm.Users(orm.UserWhere.Email.EQ(email)).One(ctx, i.dbConn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return orm.User{}, ErrNotFound
		}

		return orm.User{}, pkgerrors.WithStack(err)
	}

	return *u, nil
}

// CreateUser creates a new user entity with the corresponding email
func (i impl) CreateUser(ctx context.Context, email string) (int, error) {
	userEntity := orm.User{
		Email: email,
	}

	return userEntity.ID, pkgerrors.WithStack(userEntity.Insert(ctx, i.dbConn, boil.Infer()))
}
