package model

// MakeRelationship represents the data of request making relationship
type MakeRelationship struct {
	FromFriend, ToFriend string
}

// CommonFriend represents the data of request getting common friend
type CommonFriend struct {
	FirstUser, SecondUser string
}

// UpdateInfo represents the data of request getting data for update receiver
type UpdateInfo struct {
	Sender, Message string
}
