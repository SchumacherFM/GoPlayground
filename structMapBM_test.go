package GoPlayground

import (
	"html/template"
	"io/ioutil"
	"testing"
)

const tpl = "{{.Title}} {{.Firstname}} {{.Lastname}} {{.Rank}} {{.Married}}"

var btmpl *template.Template

// Regarding to http://talks.golang.org/2015/tricks.slide#6
// Benchmark results via Go1.5
var s = []struct {
	i int
	s string
}{
	{6 * 9, "Question"},

	{42, "Answer"},
}

var t = []struct {
	i int
	s string
}{
	{6 * 9, "Question"},
	{42, "Answer"},
}

func init() {
	var err error
	btmpl, err = template.New("test").Parse(tpl)
	if err != nil {
		panic(err)
	}
}

// BenchmarkTplMapStringIf-4	  200000	     11249 ns/op	    1760 B/op	      48 allocs/op
func BenchmarkTplMapStringIf(b *testing.B) {
	data := map[string]interface{}{
		"Title":     "Dr",
		"Firstname": "Carl",
		"Lastname":  "Sagan",
		"Rank":      7,
		"Married":   true,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := btmpl.Execute(ioutil.Discard, data); err != nil {
			b.Error(err)
		}
	}
}

// BenchmarkTpl_AnonStruct-4	  200000	     11007 ns/op	    1744 B/op	      44 allocs/op
func BenchmarkTpl_AnonStruct(b *testing.B) {
	data := struct {
		Title               string
		Firstname, Lastname string
		Rank                int
		Married             bool
	}{
		"Dr", "Carl", "Sagan", 7, true,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := btmpl.Execute(ioutil.Discard, data); err != nil {
			b.Error(err)
		}
	}
}
