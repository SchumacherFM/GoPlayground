// run with: $ go test -v  --bench=BenchmarkNan -test.benchmem .
// http://research.swtch.com/randhash
package GoPlayground

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func build(n int) map[float64]int {
	m := map[float64]int{}
	for i := 0; i < n; i++ {
		m[math.NaN()] = 1
	}
	return m
}

func BenchmarkNan(b *testing.B) {

	for n := 0; n < b.N; n++ {

		n := 1
		for i := 0; i < 20; i++ {
			t := time.Now()
			build(n)
			//            b.Logf("%6d %10.6f\n", n, time.Since(t).Seconds())
			fmt.Printf("%6d %10.6f\n", n, time.Since(t).Seconds())
			n *= 2
		}
	}
}
