package GoPlayground

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
)

var bigSlice = make([]int64, 1000)

func init() {
	for i := range bigSlice {
		bigSlice[i] = rand.Int63()
	}
}

/*
Benchmark_ForRange-4   	   10000	    187468 ns/op	   58805 B/op	    1000 allocs/op
Benchmark_ForRange-4   	   10000	    162055 ns/op	    8000 B/op	    1000 allocs/op
Benchmark_ForRange-4   	   10000	    162687 ns/op	    8000 B/op	    1000 allocs/op
Benchmark_ForRange-4   	   10000	    164929 ns/op	    8000 B/op	    1000 allocs/op
Benchmark_ForRange-4   	   10000	    163521 ns/op	    8000 B/op	    1000 allocs/op
Benchmark_ForLen-4     	   10000	    159295 ns/op	    8000 B/op	    1000 allocs/op
Benchmark_ForLen-4     	   10000	    159210 ns/op	    8000 B/op	    1000 allocs/op
Benchmark_ForLen-4     	   10000	    157978 ns/op	    8000 B/op	    1000 allocs/op
Benchmark_ForLen-4     	   10000	    164814 ns/op	    8000 B/op	    1000 allocs/op
Benchmark_ForLen-4     	   10000	    163690 ns/op	    8000 B/op	    1000 allocs/op
*/

var forBuf = bytes.NewBuffer(make([]byte, 0, 1000000))

func Benchmark_ForRange(b *testing.B) {
	forBuf.Reset()

	for i := 0; i < b.N; i++ {
		for _, s := range bigSlice {
			if _, err := fmt.Fprintf(forBuf, "%d-", s); err != nil {
				b.Fatal(err)
			}
		}
	}
}

func Benchmark_ForLen(b *testing.B) {
	forBuf.Reset()

	for i := 0; i < b.N; i++ {
		var err error
		for j := 0; j < len(bigSlice) && err == nil; j++ {
			_, err = fmt.Fprintf(forBuf, "%d-", bigSlice[j])
		}
		if err != nil {
			b.Fatal(err)
		}
	}
}
