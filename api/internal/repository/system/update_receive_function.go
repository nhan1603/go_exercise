package system

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"strings"
)

// UpdateReceiver list all the emails that receive update from source email
func (i impl) UpdateReceiver(ctx context.Context, email, message string) ([]string, error) {

	// find users mentioned in the update
	words := strings.Fields(message)

	var emailList []string

	for _, word := range words {
		if strings.Contains(word, "@") {
			emailList = append(emailList, word)
		}
	}

	user, err := i.findUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	emailId := user.ID

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

	rows, errs := i.dbConn.Query(query, emailId, BLOCK, pq.Array(emailList), emailId, BLOCK)

	if errs != nil {
		if errors.Cause(errs) == sql.ErrNoRows {
			return result, nil
		}
		return nil, errors.Wrap(errs, "orm: unable to select from user and relationship")
	}
	defer rows.Close()

	var receiveEmail string

	for rows.Next() {
		err := rows.Scan(&receiveEmail)
		if err != nil {
			return nil, errors.Wrap(err, "orm: unable to select from user and relationship")
		}
		result = append(result, receiveEmail)
	}

	return result, nil
}
