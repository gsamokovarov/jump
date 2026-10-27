package fuzzy

// Length computes the LCS length of two string using an altered (and memoized)
// version of the algorithm described at http://bit.ly/1ICdGss.
func Length(left, right string) uint64 {
	var shorter, longer string

	if len(left) < len(right) {
		shorter, longer = left, right
	} else {
		longer, shorter = left, right
	}

	shorterLen, longerLen := len(shorter), len(longer)

	current := make([]uint64, shorterLen+1)
	previous := make([]uint64, shorterLen+1)

	for i := range longerLen {
		for j := range shorterLen {
			if shorter[j] == longer[i] {
				current[j+1] = previous[j] + 1
			} else {
				current[j+1] = max(current[j], previous[j+1])
			}
		}

		for j := range shorterLen {
			previous[j+1] = current[j+1]
		}
	}

	return current[shorterLen]
}
