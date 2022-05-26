package system

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"gobase/api/internal/repository/orm"
	"gobase/api/pkg/utils"
)

// CheckExistedFriend will check a relationship between two emails has already existed or not
func (i impl) CheckExistedFriend(ctx context.Context, emailId1, emailId2 int) error {

	relaObj := &orm.Relationship{}

	sel := "*"

	query := fmt.Sprintf(
		`select %s from "relationship" where 
			(("first_email_id"=$1 and "second_email_id" = $2) OR ("first_email_id"=$3 and "second_email_id" = $4))
			AND ("status" = $5 OR "status" = $6)`, sel,
	)

	q := queries.Raw(query, emailId1, emailId2, emailId2, emailId1, FRIEND, BLOCK)

	err := q.Bind(ctx, i.dbConn, relaObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil
		}

		return errors.Wrap(err, "orm: unable to select from user")
	}

	return errors.New("Cannot create new friendship.")
}

// AddFriend will create a relationship entity for two email
func (i impl) AddFriend(ctx context.Context, emailId1, emailId2 int) error {

	relaFriend1 := orm.Relationship{
		FirstEmailID:  emailId1,
		SecondEmailID: emailId2,
		Status:        FRIEND,
	}

	relaFriend2 := orm.Relationship{
		FirstEmailID:  emailId2,
		SecondEmailID: emailId1,
		Status:        FRIEND,
	}

	errInsert := relaFriend1.Insert(ctx, i.dbConn, boil.Infer())

	errInsert2 := relaFriend2.Insert(ctx, i.dbConn, boil.Infer())

	return utils.MergeErrDB(errInsert, errInsert2)
}
