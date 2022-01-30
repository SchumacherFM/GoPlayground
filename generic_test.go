package GoPlayground

import (
	"fmt"
	"testing"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type int65 int32

func TestGenerics001(t *testing.T) {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	int65s := map[string]int65{
		"first":  34,
		"second": 13,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	t.Logf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats),
	)
	floatSlice := maps.Values(floats)
	slices.SortFunc(floatSlice, func(a, b float64) bool {
		return a < b
	})
	//_ = slices.Sort[[]int]

	t.Logf("Generic Sums: %v and %v and %v\n%#v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
		SumIntsOrFloats(int65s),
		floatSlice,
	)
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V ~int64 | ~int32 | float64](m map[K]V) V {
	var s V
	fmt.Printf("%T\n", m)
	for _, v := range m {
		s += v
	}
	return s
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
