package relationship

import (
	"context"
	"fmt"
	"github.com/juliangruber/go-intersect"
	"gobase/api/internal/model"
	"gobase/api/pkg/utils"
)

// FindFriendList will return a list of friends of an email address
func (i impl) FindFriendList(ctx context.Context, email string) ([]string, error) {

	user1, err1 := i.repo.User().FindUserByEmail(ctx, email)
	if err1 != nil {
		return nil, err1
	}

	usrId := user1.ID

	return i.repo.Relationship().FindFriendList(ctx, usrId)
}

// FindCommonFriends will return a list of common friends between to email addresses
func (i impl) FindCommonFriends(ctx context.Context, input model.CommonFriend) ([]string, error) {

	listFriend1, err1 := i.FindFriendList(ctx, input.FirstUser)

	listFriend2, err2 := i.FindFriendList(ctx, input.SecondUser)

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
