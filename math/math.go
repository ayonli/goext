package math

// Returns the maximal value from the given values.
func Max(first float64, rest ...float64) float64 {
	for _, v := range rest {
		if v > first {
			first = v
		}
	}

	return first
}

// Returns the minimal value from the given values.
func Min(first float64, rest ...float64) float64 {
	for _, v := range rest {
		if v < first {
			first = v
		}
	}

	return first
}

// Returns the sum value of the given values.
func Sum(values ...float64) float64 {
	sum := 0.0

	for _, v := range values {
		sum += v
	}

	return sum
}

// Returns a the product value multiplied by the given values.
func Product(first float64, rest ...float64) float64 {
	for _, v := range rest {
		first *= v
	}

	return first
}
