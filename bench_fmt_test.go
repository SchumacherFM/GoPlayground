package GoPlayground

import (
	"fmt"
	"testing"
)

type DebugFunc1 func(...struct {
	Key   string
	Value interface{}
})
type DebugFunc2 func(map[string]interface{})

func bmDFStruct(args ...struct {
	Key   string
	Value interface{}
}) {
	_ = fmt.Sprintf("DebugFunc1: %v\n", args)
}

func bmDfMap(arg map[string]interface{}) {
	_ = fmt.Sprintf("DebugFunc2: %v\n", arg)
}

func Benchmark_DebugStruct(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		bmDFStruct([]struct {
			Key   string
			Value interface{}
		}{
			{"key1", 1},
			{"key2", "val2"},
		}...)
	}
}

func Benchmark_DebugMap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		bmDfMap(map[string]interface{}{
			"key1": 1,
			"key2": "val2",
		})
	}
}

const toAppendStr = `Finally, how will we ship and deliver Go 2?

I think the best plan would be to ship the backwards-compatible parts of Go 2 incrementally, feature by feature, as part of the Go 1 release sequence. This has a few important properties. First, it keeps the Go 1 releases on the usual schedule, to continue the timely bug fixes and improvements that users now depend on. Second, it avoids splitting development effort between Go 1 and Go 2. Third, it avoids divergence between Go 1 and Go 2, to ease everyone's eventual migration. Fourth, it allows us to focus on and deliver one change at a time, which should help maintain quality. Fifth, it will encourage us to design features to be backwards-compatible.

We will need time to discuss and plan before any changes start landing in Go 1 releases, but it seems plausible to me that we might start seeing minor changes about a year from now, for Go 1.12 or so. That also gives us time to land package management support first.

Once all the backwards-compatible work is done, say in Go 1.20, then we can make the backwards-incompatible changes in Go 2.0. If there turn out to be no backwards-incompatible changes, maybe we just declare that Go 1.20 is Go 2.0. Either way, at that point we will transition from working on the Go 1.X release sequence to working on the Go 2.X sequence, perhaps with an extended support window for the final Go 1.X release.

This is all a bit speculative, and the specific release numbers I just mentioned are placeholders for ballpark estimates, but I want to make clear that we're not abandoning Go 1, and that in fact we will bring Go 1 along to the greatest extent possible.`

var toAppend = []byte(toAppendStr)

var container = make([]byte, 0, len(toAppend)*2)

func BenchmarkAppend(b *testing.B) {
	b.Run("bytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			container = append(container, toAppend...)
			container = container[:0]
		}
	})
	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			container = append(container, toAppendStr...)
			container = container[:0]
		}
	})
	b.Run("copy", func(b *testing.B) {
		container = container[:len(toAppend)]
		for i := 0; i < b.N; i++ {
			copy(container, toAppend)
			container = container[:0]
		}
	})
}
