package server

// Empty is an empty struct.
type Empty struct {
}

// UsersList contains list of Game users.
type UsersList struct {
	Users []*PublicUser
}

// PollTurn will be returned for every user while waiting for current players move.
type PollTurn struct {
	CurrentPlayer      string
	CurrentPlayerAlpha float32
	ShotPerformed      bool
}
