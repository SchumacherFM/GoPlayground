# Go Playground

Various weird code examples.

Benchmark Results in package `db`:

```
$ go test -v -run=XX -bench=. -benchmem .
goos: darwin
goarch: amd64
pkg: github.com/SchumacherFM/GoPlayground/db
BenchmarkMapStringScan-4   	    1000	   2215154 ns/op	  422522 B/op	   20912 allocs/op
BenchmarkStrStrScan-4      	     500	   2226420 ns/op	  422418 B/op	   20911 allocs/op
BenchmarkRowMapString-4    	     500	   3354820 ns/op	 1528206 B/op	   36859 allocs/op
BenchmarkSQLx/append-4     	     500	   3956826 ns/op	  535730 B/op	   22677 allocs/op
BenchmarkSQLx/select-4     	     300	   4177540 ns/op	  593676 B/op	   24451 allocs/op
BenchmarkCSFWdbr-4         	     500	   3251117 ns/op	  563784 B/op	   24448 allocs/op
BenchmarkSqlStruct-4       	     300	   5405734 ns/op	  961035 B/op	   29767 allocs/op
BenchmarkSqlBoiler/all-4   	     300	   4359826 ns/op	  847610 B/op	   26239 allocs/op
PASS
ok  	github.com/SchumacherFM/GoPlayground/db	15.808s
```
