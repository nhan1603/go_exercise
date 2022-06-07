package relationship

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

// UpdateReceiver list all the emails that receive update from source email
func (i impl) UpdateReceiver(ctx context.Context, emailId int, emailList []string) ([]string, error) {

	// find users that can receive update from sender
	sel := "usr.\"email\""

	query := fmt.Sprintf(
		`SELECT DISTINCT %s FROM "relationship" rela RIGHT JOIN "user" usr ON usr."id" = rela."first_email_id" 
				WHERE (rela."second_email_id" = $1 and rela."status" != $2) 
				OR (usr."email" = any($3)
				AND (rela."first_email_id" IS NULL 
				OR usr."id" NOT IN (SELECT first_email_id from "relationship" WHERE
				second_email_id = $4 and status = $5)))`, sel,
	)

	var result []string

	rows, errs := i.dbConn.Query(query, emailId, RelationshipTypeBlock, pq.Array(emailList), emailId, RelationshipTypeBlock)

	if errs != nil {
		if errors.Cause(errs) == sql.ErrNoRows {
			return result, nil
		}
		return nil, errors.WithStack(errs)
	}
	defer rows.Close()

	var receiveEmail string

	for rows.Next() {
		err := rows.Scan(&receiveEmail)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		result = append(result, receiveEmail)
	}

	return result, nil
}
