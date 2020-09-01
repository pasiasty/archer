package server

// Trajectory stores information of generated arrow trajectory.
type Trajectory struct {
	ArrowStates  []ArrowState
	KilledPlayer string
}

// ArrowState is the state of single simulation frame for the arrow.
type ArrowState struct {
	Time        float32
	Position    Point
	Orientation float32
}
