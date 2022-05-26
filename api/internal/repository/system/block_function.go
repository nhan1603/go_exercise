package system

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gobase/api/internal/repository/orm"
	"gobase/api/pkg/utils"
)

// CheckExistedBlock will check that the first email has already Blocked the second one or not
func (i impl) CheckExistedBlock(ctx context.Context, emailId1, emailId2 int) error {
	_, err := orm.Relationships(orm.RelationshipWhere.FirstEmailID.EQ(emailId1),
		orm.RelationshipWhere.SecondEmailID.EQ(emailId2), orm.RelationshipWhere.Status.EQ(BLOCK)).One(ctx, i.dbConn)

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil
		}
		return pkgerrors.WithStack(err)
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
