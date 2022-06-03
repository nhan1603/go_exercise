package relationship

import (
	"context"
	"gobase/api/internal/model"
	"strings"
)

// UpdateReceiver returns a list of emails that will receive message from certain user
func (i impl) UpdateReceiver(ctx context.Context, input model.UpdateInfo) ([]string, error) {
	user, err := i.repo.User().FindUserByEmail(ctx, input.Sender)
	if err != nil {
		return nil, err
	}

	// find users mentioned in the update
	words := strings.Fields(input.Message)

	var emailList []string

	for _, word := range words {

		if word[0:1] == "@" {
			word = word[1:]
		}

		if strings.Contains(word, "@") {
			emailList = append(emailList, word)
		}
	}

	return i.repo.Relationship().UpdateReceiver(ctx, user.ID, emailList)
}
