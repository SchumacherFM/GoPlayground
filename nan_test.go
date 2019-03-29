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

type TestE struct {
	Field1 string
}

type TestC struct {
	Data []*TestE
}

var globalField string

func (m *TestC) WithIF() {
	if len(m.Data) > 0 {
		for _, msg := range m.Data {
			globalField = msg.Field1
		}
	}
}

func (m *TestC) WithoutIF() {
	for _, msg := range m.Data {
		globalField = msg.Field1
	}
}

func BenchmarkIfOrFor(b *testing.B) {
	b.Run("WithIF", func(b *testing.B) {
		c := new(TestC)
		for i := 0; i < b.N; i++ {
			c.WithIF()
		}
	})
	b.Run("WithoutIF", func(b *testing.B) {
		c := new(TestC)
		for i := 0; i < b.N; i++ {
			c.WithoutIF()
		}
	})
}
