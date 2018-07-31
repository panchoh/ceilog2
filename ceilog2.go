package ceilog2

import (
	"math"
	"math/bits"
)

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
func GetHeight1(x uint64) uint64 {
	return uint64(bits.Len64(x))
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

// Return ceil(log_2(x))
func GetHeight3(x uint64) uint64 {
	r := uint64(0)
	for ; x != 0; x >>= 1 {
		r++
	}
	return r
}

var t [6]uint64 = [6]uint64{0xFFFFFFFF00000000, 0xFFFF0000, 0xFF00, 0xF0, 0xC, 0x2}

func GetHeight4(x uint64) uint64 {
	var i, j, k, y uint64
	j = 32

	if (x & (x - 1)) == 0 {
		y = 0
	} else {
		y = 1
	}

	for i = 0; i < 6; i++ {
		if (x & t[i]) == 0 {
			k = 0
		} else {
			k = j
		}
		y += k
		x >>= k
		j >>= 1
	}

	return y
}

// int ceil_log2(unsigned long long x)
// {
//   static const unsigned long long t[6] = {
//     0xFFFFFFFF00000000ull,
//     0x00000000FFFF0000ull,
//     0x000000000000FF00ull,
//     0x00000000000000F0ull,
//     0x000000000000000Cull,
//     0x0000000000000002ull
//   };

//   int y = (((x & (x - 1)) == 0) ? 0 : 1);
//   int j = 32;
//   int i;

//   for (i = 0; i < 6; i++) {
//     int k = (((x & t[i]) == 0) ? 0 : j);
//     y += k;
//     x >>= k;
//     j >>= 1;
//   }

//   return y;
// }
