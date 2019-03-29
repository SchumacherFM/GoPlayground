// run with: $ go test --bench=. -test.benchmem .

package GoPlayground

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/bufio.v1"
)

const fileName = "test.txt"

func BenchmarkFileReading(b *testing.B) {

	b.Run("unbuffered", func(b *testing.B) {
		f, err := os.Open(fileName)
		if err != nil {
			b.Fatal(err)
		}
		defer f.Close()
		for i := 0; i < b.N; i++ {
			if _, err := io.Copy(ioutil.Discard, f); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("buffered 64k", func(b *testing.B) {
		f, err := os.Open(fileName)
		if err != nil {
			b.Fatal(err)
		}
		defer f.Close()
		fb := bufio.NewReaderSize(f, 1024*64)
		for i := 0; i < b.N; i++ {
			if _, err := io.Copy(ioutil.Discard, fb); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("buffered 16k", func(b *testing.B) {
		f, err := os.Open(fileName)
		if err != nil {
			b.Fatal(err)
		}
		defer f.Close()
		fb := bufio.NewReaderSize(f, 1024*16)
		for i := 0; i < b.N; i++ {
			if _, err := io.Copy(ioutil.Discard, fb); err != nil {
				b.Fatal(err)
			}
		}
	})

}
