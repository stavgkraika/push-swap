// Package internal provides core utilities for the push-swap program,
// including argument parsing and stack operations.
package internal

import (
	"errors"
	"strconv"
	"strings"
)

// Sentinel errors returned by ParseArgs so callers can distinguish failure
// modes without inspecting error strings.
var (
	// ErrInvalidInteger is returned when a token cannot be parsed as a
	// base-10 integer (e.g. "abc", "3.14").
	ErrInvalidInteger = errors.New("invalid integer")

	// ErrDuplicateValue is returned when the same integer appears more
	// than once across all provided arguments.
	ErrDuplicateValue = errors.New("duplicate value")
)

// ParseArgs converts raw command-line arguments into an ordered slice of
// unique integers.
//
// Each element of args may contain multiple space-separated tokens, so both
// individual values ("3", "1") and a single quoted string ("3 1") work.
//
// Returns an empty (non-nil) slice when args is empty.
// Returns ErrInvalidInteger if any token is not a valid integer.
// Returns ErrDuplicateValue if the same integer appears more than once.
func ParseArgs(args []string) ([]int, error) {
	// Fast-path: nothing to parse.
	if len(args) == 0 {
		return []int{}, nil
	}

	// Flatten all arguments into individual tokens so that both
	// "./push-swap 3 1 2" and "./push-swap '3 1 2'" are handled uniformly.
	tokens := make([]string, 0)
	for _, arg := range args {
		parts := strings.Fields(arg) // split on any whitespace
		tokens = append(tokens, parts...)
	}

	// Pre-allocate with expected capacity to avoid reallocations.
	values := make([]int, 0, len(tokens))
	seen := make(map[int]struct{}, len(tokens)) // tracks already-encountered values

	for _, token := range tokens {
		// Attempt to parse the token as a base-10 integer.
		v, err := strconv.Atoi(token)
		if err != nil {
			return nil, ErrInvalidInteger
		}

		// Reject duplicates to preserve the uniqueness invariant.
		if _, exists := seen[v]; exists {
			return nil, ErrDuplicateValue
		}

		seen[v] = struct{}{} // empty struct uses no memory
		values = append(values, v)
	}

	return values, nil
}
