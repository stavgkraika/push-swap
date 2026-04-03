package internal

type State struct {
	A *Stack
	B *Stack
}

func NewState(initialA []int) *State {
	return &State{
		A: NewStack(initialA),
		B: NewStack(nil),
	}
}