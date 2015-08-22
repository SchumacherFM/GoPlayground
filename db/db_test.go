package main

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

//$ go test -v -bench=. -benchmem .
//testing: warning: no tests to run
//PASS
//BenchmarkMapStringScan	 300	   4388248 ns/op	  416755 B/op	   20602 allocs/op
//BenchmarkStrStrScan	     300	   3990020 ns/op	  416629 B/op	   20602 allocs/op
//BenchmarkRowMapString	     200	   7937005 ns/op	 1505452 B/op	   36304 allocs/op
func BenchmarkMapStringScan(b *testing.B) {

	db, err := sql.Open("mysql", "magento-1-8:magento-1-8@tcp(:3306)/magento-1-8")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows, err := db.Query(TEST_QUERY)
		if err != nil {
			b.Fatal(err)
		}
		defer rows.Close()
		columnNames, err := rows.Columns()
		if err != nil {
			b.Fatal(err)
		}
		rc := NewMapStringScan(columnNames)
		for rows.Next() {
			err := rc.Update(rows)
			if err != nil {
				b.Fatal(err)
			}
			_ = rc.Get()
		}
		if err = rows.Err(); err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkStrStrScan(b *testing.B) {

	db, err := sql.Open("mysql", "magento-1-8:magento-1-8@tcp(:3306)/magento-1-8")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows, err := db.Query(TEST_QUERY)
		if err != nil {
			b.Fatal(err)
		}
		defer rows.Close()
		columnNames, err := rows.Columns()
		if err != nil {
			b.Fatal(err)
		}
		rc := NewStringStringScan(columnNames)
		for rows.Next() {
			err := rc.Update(rows)
			if err != nil {
				b.Fatal(err)
			}
			_ = rc.Get()
		}
		if err = rows.Err(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRowMapString(b *testing.B) {

	db, err := sql.Open("mysql", "magento-1-8:magento-1-8@tcp(:3306)/magento-1-8")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows, err := db.Query(TEST_QUERY)
		if err != nil {
			b.Fatal(err)
		}
		defer rows.Close()
		columnNames, err := rows.Columns()
		if err != nil {
			b.Fatal(err)
		}

		for rows.Next() {
			_, err := rowMapString(columnNames, rows)
			if err != nil {
				b.Fatal(err)
			}
		}
		if err = rows.Err(); err != nil {
			b.Fatal(err)
		}
	}
}
