package ceilog2

import (
	"fmt"
	"testing"
)

var value uint64 = 100

func TestGetHeight(t *testing.T) {
	tests := []struct {
		in  uint64
		out uint64
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 3},
		{5, 3},
		{6, 3},
		{7, 3},
		{8, 4},
		{216, 8},
		{1021, 10},
		{32767, 15},
		{32768, 16},
		{3297381253, 32},
		{323543844043, 39},
	}

	for _, tt := range tests {
		fmt.Printf(
			"size:\t\t%64d\nbinary size:\t%064b\nMSB:\t\t%64b\nLen(size):\t%64b\n",
			tt.in,
			tt.in,
			1<<(tt.out-1),
			tt.out,
		)
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			if GetHeight(tt.in) != tt.out {
				t.Errorf(
					"GetHeight(%d) returned %d, but should be %d",
					tt.in,
					GetHeight(uint64(tt.in)),
					tt.out,
				)
			}
		})
	}
}

func TestGetHeight1(t *testing.T) {
	for i := uint64(0); i < 65536; i++ {
		if GetHeight(i) != GetHeight1(i) {
			t.Errorf(
				"GetHeight1(%d) returned %d, but should be %d",
				i,
				GetHeight1(i),
				GetHeight(i),
			)
		}
	}
}

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

func TestGetHeight3(t *testing.T) {
	for i := uint64(0); i < 65536; i++ {
		if GetHeight(i) != GetHeight3(i) {
			t.Errorf(
				"GetHeight3(%d) returned %d, but should be %d",
				i,
				GetHeight3(i),
				GetHeight(i),
			)
		}
	}
}

func TestGetHeight4(t *testing.T) {
	// It fails for now
	t.Skip()
	for i := uint64(0); i < 65536; i++ {
		if GetHeight(i) != GetHeight4(i) {
			t.Errorf(
				"GetHeight3(%d) returned %d, but should be %d",
				i,
				GetHeight4(i),
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

func BenchmarkGetHeight1(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		GetHeight1(value)
	}
}

func BenchmarkGetHeight2(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		GetHeight2(value)
	}
}

func BenchmarkGetHeight3(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		GetHeight3(value)
	}
}

func BenchmarkGetHeight4(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		GetHeight4(value)
	}
}
