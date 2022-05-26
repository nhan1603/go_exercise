package rest

import (
	"context"
	"strings"
)

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
