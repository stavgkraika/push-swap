package internal

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack([]int{1, 2, 3})
	if !reflect.DeepEqual(s.Values(), []int{1, 2, 3}) {
		t.Fatalf("unexpected values: %v", s.Values())
	}
}

func TestNewStack_IsolatesSlice(t *testing.T) {
	orig := []int{1, 2, 3}
	s := NewStack(orig)
	orig[0] = 99
	if s.Values()[0] == 99 {
		t.Fatal("NewStack should copy the slice, not reference it")
	}
}

func TestPushTop(t *testing.T) {
	s := NewStack([]int{2, 3})
	s.PushTop(1)
	if s.Values()[0] != 1 {
		t.Fatalf("expected 1 at top, got %v", s.Values()[0])
	}
}

func TestPopTop(t *testing.T) {
	s := NewStack([]int{1, 2, 3})
	v, ok := s.PopTop()
	if !ok || v != 1 {
		t.Fatalf("expected (1, true), got (%v, %v)", v, ok)
	}
	if s.Size() != 2 {
		t.Fatalf("expected size 2 after pop, got %d", s.Size())
	}
}

func TestPopTop_Empty(t *testing.T) {
	s := NewStack([]int{})
	v, ok := s.PopTop()
	if ok || v != 0 {
		t.Fatalf("expected (0, false) on empty stack, got (%v, %v)", v, ok)
	}
}

func TestPeekTop(t *testing.T) {
	s := NewStack([]int{5, 6})
	v, ok := s.PeekTop()
	if !ok || v != 5 {
		t.Fatalf("expected (5, true), got (%v, %v)", v, ok)
	}
	if s.Size() != 2 {
		t.Fatal("PeekTop should not remove the element")
	}
}

func TestPeekTop_Empty(t *testing.T) {
	s := NewStack([]int{})
	v, ok := s.PeekTop()
	if ok || v != 0 {
		t.Fatalf("expected (0, false) on empty stack, got (%v, %v)", v, ok)
	}
}

func TestIsEmpty(t *testing.T) {
	s := NewStack([]int{})
	if !s.IsEmpty() {
		t.Fatal("expected IsEmpty true")
	}
	s.PushTop(1)
	if s.IsEmpty() {
		t.Fatal("expected IsEmpty false after push")
	}
}

func TestValues_ReturnsCopy(t *testing.T) {
	s := NewStack([]int{1, 2, 3})
	v := s.Values()
	v[0] = 99
	if s.Values()[0] == 99 {
		t.Fatal("Values should return a copy, not a reference")
	}
}

func TestSwapTopTwo(t *testing.T) {
	s := NewStack([]int{1, 2, 3})
	s.SwapTopTwo()
	if !reflect.DeepEqual(s.Values(), []int{2, 1, 3}) {
		t.Fatalf("expected [2 1 3], got %v", s.Values())
	}
}

func TestSwapTopTwo_LessThanTwo(t *testing.T) {
	s := NewStack([]int{1})
	s.SwapTopTwo() // should be a no-op
	if !reflect.DeepEqual(s.Values(), []int{1}) {
		t.Fatalf("expected [1], got %v", s.Values())
	}
}

func TestRotate(t *testing.T) {
	s := NewStack([]int{1, 2, 3})
	s.Rotate()
	if !reflect.DeepEqual(s.Values(), []int{2, 3, 1}) {
		t.Fatalf("expected [2 3 1], got %v", s.Values())
	}
}

func TestRotate_LessThanTwo(t *testing.T) {
	s := NewStack([]int{1})
	s.Rotate() // should be a no-op
	if !reflect.DeepEqual(s.Values(), []int{1}) {
		t.Fatalf("expected [1], got %v", s.Values())
	}
}

func TestReverseRotate(t *testing.T) {
	s := NewStack([]int{1, 2, 3})
	s.ReverseRotate()
	if !reflect.DeepEqual(s.Values(), []int{3, 1, 2}) {
		t.Fatalf("expected [3 1 2], got %v", s.Values())
	}
}

func TestReverseRotate_LessThanTwo(t *testing.T) {
	s := NewStack([]int{1})
	s.ReverseRotate() // should be a no-op
	if !reflect.DeepEqual(s.Values(), []int{1}) {
		t.Fatalf("expected [1], got %v", s.Values())
	}
}
