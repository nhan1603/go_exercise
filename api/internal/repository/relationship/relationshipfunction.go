package relationship

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gobase/api/internal/repository/orm"
	"gobase/api/pkg/utils"
)

// CheckExistedFriend will check a relationship between two emails has already existed or not
func (i impl) CheckExistedFriend(ctx context.Context, emailId1, emailId2 int) error {

	_, err := orm.Relationships(qm.Expr(
		qm.Expr(qm.Where("first_email_id=?", emailId1), qm.And("second_email_id = ?", emailId2)),
		qm.Or2(qm.Expr(qm.Where("first_email_id=?", emailId2), qm.And("second_email_id = ?", emailId1)))),
		qm.Expr(qm.Where("status=?", FRIEND), qm.Or("status=?", BLOCK))).One(ctx, i.dbConn)

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil
		}
		return pkgerrors.WithStack(err)
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
