// Package internal provides core utilities for the push-swap program,
// including the Stack data structure and argument parsing.
package internal

// Stack is a LIFO (last-in, first-out) data structure.
// The top element is always at index 0 of the underlying slice.
type Stack struct {
	items []int
}

// NewStack creates a Stack pre-loaded with the given values.
// The slice is copied so the caller's original slice is never mutated.
func NewStack(values []int) *Stack {
	cp := make([]int, len(values))
	copy(cp, values)
	return &Stack{items: cp}
}

// PushTop inserts v at the top (index 0) of the stack.
func (s *Stack) PushTop(v int) {
	s.items = append([]int{v}, s.items...)
}

// PopTop removes and returns the top element.
// Returns (0, false) if the stack is empty.
func (s *Stack) PopTop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	v := s.items[0]
	s.items = s.items[1:] // drop the top element
	return v, true
}

// PeekTop returns the top element without removing it.
// Returns (0, false) if the stack is empty.
func (s *Stack) PeekTop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[0], true
}

// Size returns the number of elements currently in the stack.
func (s *Stack) Size() int {
	return len(s.items)
}

// IsEmpty reports whether the stack contains no elements.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Values returns a copy of the stack's elements in top-to-bottom order.
// The copy ensures the internal slice cannot be modified by the caller.
func (s *Stack) Values() []int {
	cp := make([]int, len(s.items))
	copy(cp, s.items)
	return cp
}

// SwapTopTwo swaps the two topmost elements (push-swap operation "sa"/"sb").
// Does nothing if the stack has fewer than 2 elements.
func (s *Stack) SwapTopTwo() {
	if len(s.items) < 2 {
		return
	}
	s.items[0], s.items[1] = s.items[1], s.items[0]
}

// Rotate moves the top element to the bottom (push-swap operation "ra"/"rb").
// Does nothing if the stack has fewer than 2 elements.
func (s *Stack) Rotate() {
	if len(s.items) < 2 {
		return
	}

	first := s.items[0]
	s.items = append(s.items[1:], first) // top goes to the end
}

// ReverseRotate moves the bottom element to the top (push-swap operation "rra"/"rrb").
// Does nothing if the stack has fewer than 2 elements.
func (s *Stack) ReverseRotate() {
	if len(s.items) < 2 {
		return
	}

	last := s.items[len(s.items)-1]
	s.items = append([]int{last}, s.items[:len(s.items)-1]...) // bottom goes to the top
}
