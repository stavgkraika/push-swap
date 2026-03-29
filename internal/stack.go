package ps

type Stack struct {
	items []int
}

func NewStack(values []int) *Stack {
	cp := make([]int, len(values))
	copy(cp, values)
	return &Stack{items: cp}
}

func (s *Stack) PushTop(v int) {
	s.items = append([]int{v}, s.items...)
}

func (s *Stack) PopTop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	v := s.items[0]
	s.items = s.items[1:]
	return v, true
}

func (s *Stack) PeekTop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[0], true
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Values() []int {
	cp := make([]int, len(s.items))
	copy(cp, s.items)
	return cp
}

func (s *Stack) SwapTopTwo() {
	if len(s.items) < 2 {
		return
	}
	s.items[0], s.items[1] = s.items[1], s.items[0]
}

func (s *Stack) Rotate() {
	if len(s.items) < 2 {
		return
	}

	first := s.items[0]
	s.items = append(s.items[1:], first)
}

func (s *Stack) ReverseRotate() {
	if len(s.items) < 2 {
		return
	}

	last := s.items[len(s.items)-1]
	s.items = append([]int{last}, s.items[:len(s.items)-1]...)
}
