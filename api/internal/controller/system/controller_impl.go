package system

import "context"

// AddFriend will create a friendship for two email
func (i impl) AddFriend(ctx context.Context, email1, email2 string) error {
	return i.repo.System().AddFriend(ctx, email1, email2)
}

// CreateUser will create a new user for the email
func (i impl) CreateUser(ctx context.Context, email string) error {
	return i.repo.System().CreateUser(ctx, email)
}

// FindFriendList will return a list of friends of an email address
func (i impl) FindFriendList(ctx context.Context, email string) ([]string, error) {
	return i.repo.System().FindFriendList(ctx, email)
}

// FindCommonFriends will return a list of common friends between to email addresses
func (i impl) FindCommonFriends(ctx context.Context, email1, email2 string) ([]string, error) {
	return i.repo.System().FindCommonFriends(ctx, email1, email2)
}

// Subscribe will create a subscription for of the second email for the first email
func (i impl) Subscribe(ctx context.Context, email1, email2 string) error {
	return i.repo.System().Subscribe(ctx, email1, email2)
}

// Block creates a block relation between email1 and email2
func (i impl) Block(ctx context.Context, email1, email2 string) error {
	return i.repo.System().Block(ctx, email1, email2)
}

// UpdateReceiver returns a list of emails that will receive message from certain user
func (i impl) UpdateReceiver(ctx context.Context, email, message string) ([]string, error) {
	return i.repo.System().UpdateReceiver(ctx, email, message)
}
