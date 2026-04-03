package internal

// Sa swaps the top two elements of stack A.
func Sa(s *State) {
	s.A.SwapTopTwo()
}

// Sb swaps the top two elements of stack B.
func Sb(s *State) {
	s.B.SwapTopTwo()
}

// Ss performs Sa and Sb simultaneously.
func Ss(s *State) {
	s.A.SwapTopTwo()
	s.B.SwapTopTwo()
}

// Pa pops the top element of stack B and pushes it onto stack A.
// Does nothing if stack B is empty.
func Pa(s *State) {
	v, ok := s.B.PopTop()
	if !ok {
		return
	}
	s.A.PushTop(v)
}

// Pb pops the top element of stack A and pushes it onto stack B.
// Does nothing if stack A is empty.
func Pb(s *State) {
	v, ok := s.A.PopTop()
	if !ok {
		return
	}
	s.B.PushTop(v)
}

// Ra rotates stack A upward: the top element moves to the bottom.
func Ra(s *State) {
	s.A.Rotate()
}

// Rb rotates stack B upward: the top element moves to the bottom.
func Rb(s *State) {
	s.B.Rotate()
}

// Rr performs Ra and Rb simultaneously.
func Rr(s *State) {
	s.A.Rotate()
	s.B.Rotate()
}

// Rra reverse-rotates stack A: the bottom element moves to the top.
func Rra(s *State) {
	s.A.ReverseRotate()
}

// Rrb reverse-rotates stack B: the bottom element moves to the top.
func Rrb(s *State) {
	s.B.ReverseRotate()
}

// Rrr performs Rra and Rrb simultaneously.
func Rrr(s *State) {
	s.A.ReverseRotate()
	s.B.ReverseRotate()
}
