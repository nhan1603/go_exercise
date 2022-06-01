package model

type MakeRelationship struct {
	FromFriend, ToFriend string
}

type CommonFriend struct {
	FirstUser, SecondUser string
}

type UpdateInfo struct {
	Sender, Message string
}
