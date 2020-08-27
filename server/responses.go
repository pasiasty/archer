package server

// Empty is an empty struct.
type Empty struct {
}

// UsersList contains list of Game users.
type UsersList struct {
	Users []*PublicUser
}
