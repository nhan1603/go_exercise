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

// CheckExistedBlock will check that the first email has already Blocked the second one or not
func (i impl) CheckExistedBlock(ctx context.Context, emailId1, emailId2 int) error {

	relaObj := &orm.Relationship{}

	sel := "*"

	query := fmt.Sprintf(
		"select %s from \"relationship\" where \"first_email_id\"=$1 and \"second_email_id\" = $2 and \"status\" = $3", sel,
	)

	q := queries.Raw(query, emailId1, emailId2, BLOCK)

	err := q.Bind(ctx, i.dbConn, relaObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil
		}
		return errors.Wrap(err, "orm: unable to select from user")
	}

	return errors.New("Cannot block user.")
}

// Block creates a block-relationship from the first email
func (i impl) Block(ctx context.Context, emailId1, emailId2 int) error {

	// Delete all other relationship
	deleteQuery := "DELETE from \"relationship\" where \"first_email_id\"=$1 and \"second_email_id\" = $2"

	_, errDelete := i.dbConn.Exec(deleteQuery, emailId1, emailId2)

	relaBlock := orm.Relationship{
		FirstEmailID:  emailId1,
		SecondEmailID: emailId2,
		Status:        BLOCK,
	}

	errInsert := relaBlock.Insert(ctx, i.dbConn, boil.Infer())

	return utils.MergeErrDB(errInsert, errDelete)
}
