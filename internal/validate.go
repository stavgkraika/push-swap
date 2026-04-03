package internal

// IsSortedAscending reports whether the given slice is sorted in ascending order.
func IsSortedAscending(values []int) bool {
	for i := 1; i < len(values); i++ {
		if values[i] < values[i-1] {
			return false
		}
	}
	return true
}
