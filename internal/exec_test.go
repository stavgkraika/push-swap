package internal

import (
	"testing"
)

func TestExecuteInstruction_AllValid(t *testing.T) {
	instructions := []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}
	for _, inst := range instructions {
		s := newState([]int{1, 2, 3}, []int{4, 5, 6})
		if err := ExecuteInstruction(s, inst); err != nil {
			t.Errorf("instruction %q returned unexpected error: %v", inst, err)
		}
	}
}

func TestExecuteInstruction_Invalid(t *testing.T) {
	s := newState([]int{1, 2, 3}, []int{})
	if err := ExecuteInstruction(s, "xyz"); err == nil {
		t.Fatal("expected error for invalid instruction, got nil")
	}
}
