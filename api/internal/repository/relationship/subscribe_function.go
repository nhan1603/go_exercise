package relationship

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gobase/api/internal/repository/orm"
)

// CheckExistedSubscribe will check that the first email has already subscribed the second one or not
func (i impl) CheckExistedSubscribe(ctx context.Context, emailId1, emailId2 int) error {
	_, err := orm.Relationships(qm.Where("first_email_id=?", emailId1), qm.Where("second_email_id = ?", emailId2),
		qm.Expr(qm.Where("status = ?", SUBSCRIBE), qm.Or("status = ?", BLOCK))).One(ctx, i.dbConn)

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil
		}
		return pkgerrors.WithStack(err)
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

	return pkgerrors.WithStack(relaSubscribe.Insert(ctx, i.dbConn, boil.Infer()))
}
