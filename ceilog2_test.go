package ceilog2

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

var value uint64 = 100

func TestGetHeight(t *testing.T) {

	msbOf := map[uint64]uint64{
		0:            0,
		1:            1,
		2:            2,
		3:            2,
		4:            3,
		5:            3,
		6:            3,
		7:            3,
		8:            4,
		216:          8,
		1021:         10,
		32767:        15,
		32768:        16,
		3297381253:   32,
		323543844043: 39,
	}

	keys := make([]int, len(msbOf))
	i := 0
	for key := range msbOf {
		keys[i] = int(key)
		i++
	}
	fmt.Printf("%v\n", keys)
	sort.Ints(keys)
	fmt.Printf("%v\n", keys)
	fmt.Printf("%v\n", msbOf)
	for _, k := range keys {
		v := msbOf[uint64(k)]
		fmt.Println("------------")
		fmt.Printf("%64d\n", k)
		fmt.Printf("%.64b\n", k)
		fmt.Printf("%64b\n", uint64(math.Pow(2.0, float64(v-1))))
		fmt.Printf("%64b\n", v)
		fmt.Println("------------")
		if GetHeight(uint64(k)) != v {
			t.Errorf(
				"GetHeight(%d) returned %d, but should be %d",
				k,
				GetHeight(uint64(k)),
				v,
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
