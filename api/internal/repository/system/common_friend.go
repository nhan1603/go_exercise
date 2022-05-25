package system

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

// FindFriendList will find the list of email that is friend with the provided user
func (i impl) FindFriendList(ctx context.Context, usrId int) ([]string, error) {

	sel := "usr.\"email\""

	query := fmt.Sprintf(
		`select %s from "relationship" rela left join "user" usr on usr."id" = rela."second_email_id" 
				where rela."first_email_id"=$1 and rela."status" = $2`, sel,
	)

	var result []string

	rows, errs := i.dbConn.Query(query, usrId, FRIEND)

	if errs != nil {
		if errors.Cause(errs) == sql.ErrNoRows {
			return result, nil
		}
		return nil, errors.Wrap(errs, "orm: unable to select from user and relationship")
	}
	defer rows.Close()

	var frEmail string

	for rows.Next() {
		err := rows.Scan(&frEmail)
		if err != nil {
			return nil, errors.Wrap(err, "orm: unable to select from user and relationship")
		}
		result = append(result, frEmail)
	}

	return result, nil
}
