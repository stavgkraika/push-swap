package internal

import (
	"reflect"
	"testing"
)

func newState(a, b []int) *State {
	return &State{A: NewStack(a), B: NewStack(b)}
}

func TestSa(t *testing.T) {
	s := newState([]int{1, 2, 3}, []int{})
	Sa(s)
	if !reflect.DeepEqual(s.A.Values(), []int{2, 1, 3}) {
		t.Fatalf("expected [2 1 3], got %v", s.A.Values())
	}
}

func TestSb(t *testing.T) {
	s := newState([]int{}, []int{1, 2, 3})
	Sb(s)
	if !reflect.DeepEqual(s.B.Values(), []int{2, 1, 3}) {
		t.Fatalf("expected [2 1 3], got %v", s.B.Values())
	}
}

func TestSs(t *testing.T) {
	s := newState([]int{1, 2}, []int{3, 4})
	Ss(s)
	if !reflect.DeepEqual(s.A.Values(), []int{2, 1}) {
		t.Fatalf("A: expected [2 1], got %v", s.A.Values())
	}
	if !reflect.DeepEqual(s.B.Values(), []int{4, 3}) {
		t.Fatalf("B: expected [4 3], got %v", s.B.Values())
	}
}

func TestPb(t *testing.T) {
	s := newState([]int{1, 2, 3}, []int{})
	Pb(s)
	if s.A.Size() != 2 || s.B.Size() != 1 {
		t.Fatalf("expected A size 2 and B size 1, got A=%d B=%d", s.A.Size(), s.B.Size())
	}
	v, _ := s.B.PeekTop()
	if v != 1 {
		t.Fatalf("expected 1 on top of B, got %d", v)
	}
}

func TestPb_EmptyA(t *testing.T) {
	s := newState([]int{}, []int{})
	Pb(s) // should be a no-op
	if s.A.Size() != 0 || s.B.Size() != 0 {
		t.Fatal("expected both stacks empty after Pb on empty A")
	}
}

func TestPa(t *testing.T) {
	s := newState([]int{}, []int{1, 2, 3})
	Pa(s)
	if s.B.Size() != 2 || s.A.Size() != 1 {
		t.Fatalf("expected B size 2 and A size 1, got A=%d B=%d", s.A.Size(), s.B.Size())
	}
	v, _ := s.A.PeekTop()
	if v != 1 {
		t.Fatalf("expected 1 on top of A, got %d", v)
	}
}

func TestPa_EmptyB(t *testing.T) {
	s := newState([]int{}, []int{})
	Pa(s) // should be a no-op
	if s.A.Size() != 0 || s.B.Size() != 0 {
		t.Fatal("expected both stacks empty after Pa on empty B")
	}
}

func TestRa(t *testing.T) {
	s := newState([]int{1, 2, 3}, []int{})
	Ra(s)
	if !reflect.DeepEqual(s.A.Values(), []int{2, 3, 1}) {
		t.Fatalf("expected [2 3 1], got %v", s.A.Values())
	}
}

func TestRb(t *testing.T) {
	s := newState([]int{}, []int{1, 2, 3})
	Rb(s)
	if !reflect.DeepEqual(s.B.Values(), []int{2, 3, 1}) {
		t.Fatalf("expected [2 3 1], got %v", s.B.Values())
	}
}

func TestRr(t *testing.T) {
	s := newState([]int{1, 2, 3}, []int{4, 5, 6})
	Rr(s)
	if !reflect.DeepEqual(s.A.Values(), []int{2, 3, 1}) {
		t.Fatalf("A: expected [2 3 1], got %v", s.A.Values())
	}
	if !reflect.DeepEqual(s.B.Values(), []int{5, 6, 4}) {
		t.Fatalf("B: expected [5 6 4], got %v", s.B.Values())
	}
}

func TestRra(t *testing.T) {
	s := newState([]int{1, 2, 3}, []int{})
	Rra(s)
	if !reflect.DeepEqual(s.A.Values(), []int{3, 1, 2}) {
		t.Fatalf("expected [3 1 2], got %v", s.A.Values())
	}
}

func TestRrb(t *testing.T) {
	s := newState([]int{}, []int{1, 2, 3})
	Rrb(s)
	if !reflect.DeepEqual(s.B.Values(), []int{3, 1, 2}) {
		t.Fatalf("expected [3 1 2], got %v", s.B.Values())
	}
}

func TestRrr(t *testing.T) {
	s := newState([]int{1, 2, 3}, []int{4, 5, 6})
	Rrr(s)
	if !reflect.DeepEqual(s.A.Values(), []int{3, 1, 2}) {
		t.Fatalf("A: expected [3 1 2], got %v", s.A.Values())
	}
	if !reflect.DeepEqual(s.B.Values(), []int{6, 4, 5}) {
		t.Fatalf("B: expected [6 4 5], got %v", s.B.Values())
	}
}
