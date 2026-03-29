package internal

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrInvalidInteger = errors.New("invalid integer")
	ErrDuplicateValue = errors.New("duplicate value")
)

func ParseArgs(args []string) ([]int, error) {
	if len(args) == 0 {
		return []int{}, nil
	}

	tokens := make([]string, 0)
	for _, arg := range args {
		parts := strings.Fields(arg)
		tokens = append(tokens, parts...)
	}

	values := make([]int, 0, len(tokens))
	seen := make(map[int]struct{}, len(tokens))

	for _, token := range tokens {
		v, err := strconv.Atoi(token)
		if err != nil {
			return nil, ErrInvalidInteger
		}
		if _, exists := seen[v]; exists {
			return nil, ErrDuplicateValue
		}
		seen[v] = struct{}{}
		values = append(values, v)
	}

	return values, nil
}
