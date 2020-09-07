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

// GameStatus tells the status of the game.
type GameStatus struct {
	Started       bool
	WorldSettings WorldSettings
}

// WorldSettings define settings of the world.
type WorldSettings struct {
	ShootTimeout int32
}
