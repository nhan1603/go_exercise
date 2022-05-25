package system

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"gobase/api/internal/repository/orm"
	"gobase/api/pkg/utils"
)

// CheckExistedSubscribe will check that the first email has already subscribed the second one or not
func (i impl) CheckExistedSubscribe(ctx context.Context, emailId1, emailId2 int) error {

	relaObj := &orm.Relationship{}

	sel := "*"

	query := fmt.Sprintf(
		"select %s from \"relationship\" where \"first_email_id\"=$1 and \"second_email_id\" = $2 and (\"status\" = $3 "+
			"or \"status\" = $4)", sel,
	)

	q := queries.Raw(query, emailId1, emailId2, SUBSCRIBE, BLOCK)

	err := q.Bind(ctx, i.dbConn, relaObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil
		}
		return errors.Wrap(err, "orm: unable to select from user")
	}

	return errors.New("Cannot create new subscription.")
}

// Subscribe will create a subscription for two email
func (i impl) Subscribe(ctx context.Context, email1, email2 int) error {

	relaSubscribe := orm.Relationship{
		FirstEmailID:  email1,
		SecondEmailID: email2,
		Status:        SUBSCRIBE,
	}

	errInsert := relaSubscribe.Insert(ctx, i.dbConn, boil.Infer())

	return utils.MergeErrDB(errInsert)
}
