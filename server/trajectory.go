package server

// Trajectory stores information of generated arrow trajectory.
type Trajectory struct {
	ArrowStates  []ArrowState
	CollidedWith string
}

// ArrowState is the state of single simulation frame for the arrow.
type ArrowState struct {
	Position    Vector
	Orientation float32
}
