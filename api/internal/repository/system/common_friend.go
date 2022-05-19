package system

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/juliangruber/go-intersect"
	"github.com/pkg/errors"
	"gobase/api/pkg/utils"
)

// FindFriendList will find the list of email that is friend with the provided user
func (i impl) FindFriendList(ctx context.Context, email string) ([]string, error) {

	user1, err1 := i.findUserByEmail(ctx, email)
	if err1 != nil {
		return nil, err1
	}

	usrId := user1.ID

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

// FindCommonFriends will find the list of email that is common friend with the provided users
func (i impl) FindCommonFriends(ctx context.Context, email1, email2 string) ([]string, error) {
	listFriend1, err1 := i.FindFriendList(ctx, email1)

	listFriend2, err2 := i.FindFriendList(ctx, email2)

	if err := utils.MergeErr(err1, err2); err != nil {
		return nil, err
	}

	listCommonFr := intersect.Hash(listFriend1, listFriend2)

	listResult := make([]string, len(listCommonFr))
	for ind, val := range listCommonFr {
		listResult[ind] = fmt.Sprint(val)
	}
	return listResult, nil
}
