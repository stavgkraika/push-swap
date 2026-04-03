package internal

import (
	"errors"
	"testing"
)

func TestParseArgs_Empty(t *testing.T) {
	got, err := ParseArgs([]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %v", got)
	}
}

func TestParseArgs_SingleArg(t *testing.T) {
	got, err := ParseArgs([]string{"3"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || got[0] != 3 {
		t.Fatalf("expected [3], got %v", got)
	}
}

func TestParseArgs_MultipleArgs(t *testing.T) {
	got, err := ParseArgs([]string{"3", "1", "2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []int{3, 1, 2}
	for i, v := range want {
		if got[i] != v {
			t.Fatalf("expected %v, got %v", want, got)
		}
	}
}

func TestParseArgs_SpaceSeparatedString(t *testing.T) {
	got, err := ParseArgs([]string{"3 1 2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 3 {
		t.Fatalf("expected 3 values, got %v", got)
	}
}

func TestParseArgs_NegativeNumbers(t *testing.T) {
	got, err := ParseArgs([]string{"-1", "-2", "3"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got[0] != -1 || got[1] != -2 || got[2] != 3 {
		t.Fatalf("unexpected values: %v", got)
	}
}

func TestParseArgs_InvalidInteger(t *testing.T) {
	_, err := ParseArgs([]string{"abc"})
	if !errors.Is(err, ErrInvalidInteger) {
		t.Fatalf("expected ErrInvalidInteger, got %v", err)
	}
}

func TestParseArgs_FloatInvalid(t *testing.T) {
	_, err := ParseArgs([]string{"3.14"})
	if !errors.Is(err, ErrInvalidInteger) {
		t.Fatalf("expected ErrInvalidInteger, got %v", err)
	}
}

func TestParseArgs_Duplicate(t *testing.T) {
	_, err := ParseArgs([]string{"1", "2", "1"})
	if !errors.Is(err, ErrDuplicateValue) {
		t.Fatalf("expected ErrDuplicateValue, got %v", err)
	}
}

func TestParseArgs_DuplicateInSameArg(t *testing.T) {
	_, err := ParseArgs([]string{"1 1"})
	if !errors.Is(err, ErrDuplicateValue) {
		t.Fatalf("expected ErrDuplicateValue, got %v", err)
	}
}
