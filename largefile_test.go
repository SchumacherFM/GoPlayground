// run with: $ go test --bench=. -test.benchmem .

package GoPlayground

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/bufio.v1"
)

// to generate this file:
// $ dd if=/dev/urandom of=test.txt bs=1048576 count=100
// for 100MB
const fileName = "test.txt"

func BenchmarkFileReading(b *testing.B) {

	b.Run("unbuffered", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f, err := os.Open(fileName)
			if err != nil {
				b.Fatal(err)
			}
			if _, err := io.Copy(ioutil.Discard, f); err != nil {
				b.Fatal(err)
			}
			f.Close()
		}
	})
	b.Run("buffered 16k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f, err := os.Open(fileName)
			if err != nil {
				b.Fatal(err)
			}
			fb := bufio.NewReaderSize(f, 1024*16)
			if _, err := io.Copy(ioutil.Discard, fb); err != nil {
				b.Fatal(err)
			}
			f.Close()
		}
	})
	b.Run("buffered 64k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f, err := os.Open(fileName)
			if err != nil {
				b.Fatal(err)
			}
			fb := bufio.NewReaderSize(f, 1024*64)
			if _, err := io.Copy(ioutil.Discard, fb); err != nil {
				b.Fatal(err)
			}
			f.Close()
		}
	})
	b.Run("buffered 128k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f, err := os.Open(fileName)
			if err != nil {
				b.Fatal(err)
			}
			fb := bufio.NewReaderSize(f, 1024*128)
			if _, err := io.Copy(ioutil.Discard, fb); err != nil {
				b.Fatal(err)
			}
			f.Close()
		}
	})
}
