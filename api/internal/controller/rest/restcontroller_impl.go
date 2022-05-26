package rest

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/juliangruber/go-intersect"
	pkgerrors "github.com/pkg/errors"
	"gobase/api/internal/model"
	"gobase/api/pkg/utils"
	"strings"
)

// AddFriend will create a friendship for two email
func (i impl) AddFriend(ctx context.Context, input model.MakeFriend) error {
	user1, err := i.repo.System().FindUserByEmail(ctx, input.FromFriend)
	if err != nil {
		return err
	}

	user2, err := i.repo.System().FindUserByEmail(ctx, input.ToFriend)
	if err != nil {
		return err
	}

	if err := i.repo.System().CheckExistedFriend(ctx, user1.ID, user2.ID); err != nil {
		return pkgerrors.WithStack(err)
	}

	return pkgerrors.WithStack(i.repo.System().AddFriend(ctx, user1.ID, user2.ID))
}

// CreateUser will create a new user for the email
func (i impl) CreateUser(ctx context.Context, email string) (int, error) {
	if _, err := i.repo.System().FindUserByEmail(ctx, email); err == nil {
		return 0, errors.New("Existed email input.")
	}

	return i.repo.System().CreateUser(ctx, email)
}

// FindFriendList will return a list of friends of an email address
func (i impl) FindFriendList(ctx context.Context, email string) ([]string, error) {

	user1, err1 := i.repo.System().FindUserByEmail(ctx, email)
	if err1 != nil {
		return nil, err1
	}

	usrId := user1.ID

	return i.repo.System().FindFriendList(ctx, usrId)
}

// FindCommonFriends will return a list of common friends between to email addresses
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

// Subscribe will create a subscription for of the second email for the first email
func (i impl) Subscribe(ctx context.Context, email1, email2 string) error {

	if email1 == email2 {
		return errors.New("Duplicate email input")
	}

	user1, err1 := i.repo.System().FindUserByEmail(ctx, email1)
	if err1 != nil {
		return err1
	}

	user2, err2 := i.repo.System().FindUserByEmail(ctx, email2)
	if err2 != nil {
		return err2
	}

	emailId1 := user1.ID
	emailId2 := user2.ID

	if err := i.repo.System().CheckExistedSubscribe(ctx, emailId1, emailId2); err != nil {
		return err
	}

	return i.repo.System().Subscribe(ctx, emailId1, emailId2)
}

// Block creates a block relation between email1 and email2
func (i impl) Block(ctx context.Context, email1, email2 string) error {

	if email1 == email2 {
		return errors.New("Duplicate email input")
	}

	user1, err1 := i.repo.System().FindUserByEmail(ctx, email1)
	if err1 != nil {
		return err1
	}

	user2, err2 := i.repo.System().FindUserByEmail(ctx, email2)
	if err2 != nil {
		return err2
	}

	emailId1 := user1.ID
	emailId2 := user2.ID

	if err := i.repo.System().CheckExistedBlock(ctx, emailId1, emailId2); err != nil {
		return err
	}

	return i.repo.System().Block(ctx, emailId1, emailId2)
}

// UpdateReceiver returns a list of emails that will receive message from certain user
func (i impl) UpdateReceiver(ctx context.Context, email, message string) ([]string, error) {

	user, err := i.repo.System().FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// find users mentioned in the update
	words := strings.Fields(message)

	var emailList []string

	for _, word := range words {

		if word[0:1] == "@" {
			word = word[1:]
		}

		if strings.Contains(word, "@") {
			emailList = append(emailList, word)
		}
	}

	return i.repo.System().UpdateReceiver(ctx, user.ID, emailList)
}
