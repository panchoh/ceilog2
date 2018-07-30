package ceilog2

import "testing"

var value uint64 = 100

func TestGetHeight2(t *testing.T) {
	for i := uint64(0); i < 65536; i++ {
		if GetHeight(i) != GetHeight2(i) {
			t.Errorf(
				"GetHeight2(%d) returned %d, but should be %d",
				i,
				GetHeight2(i),
				GetHeight(i),
			)
		}
	}
}

func BenchmarkGetHeight(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		GetHeight(value)
	}
}

func BenchmarkGetHeight2(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		GetHeight2(value)
	}
}
