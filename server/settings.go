package server

// GameSettings defines settings of the game.
type GameSettings struct {
	shootTimeout int32
	loopedWorld  bool
}

// GameSettingsOption defines single option of game settings.
type GameSettingsOption func(*GameSettings)

// WithShootTimeout defines shoot timeout of GameSettings.
func WithShootTimeout(timeout int32) GameSettingsOption {
	return func(gs *GameSettings) {
		gs.shootTimeout = timeout
	}
}

// WithLoopedWorld defines if world is looped.
func WithLoopedWorld(loopedWorld bool) GameSettingsOption {
	return func(gs *GameSettings) {
		gs.loopedWorld = loopedWorld
	}
}

// CreateGameSettings creates game settings.
func CreateGameSettings(opts ...GameSettingsOption) GameSettings {
	res := GameSettings{}

	for _, o := range opts {
		o(&res)
	}

	return res
}
