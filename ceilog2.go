package ceilog2

import "math"

// The log base 2 of an integer is the same as the position of the highest bit
// set (or most significant bit set, MSB).

func GetHeight(version uint64) uint64 {
	return uint64(
		math.Ceil(
			math.Log2(
				float64(
					version + 1,
				),
			),
		),
	)
}

// Return ceil(log_2(x))
func GetHeight2(version uint64) uint64 {
	var i, pow uint64
	pow = 1
	for pow <= version {
		pow *= 2
		i++
	}
	return i
}
