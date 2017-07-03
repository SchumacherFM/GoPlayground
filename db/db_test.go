package main

import (
	"database/sql"
	"testing"

	"context"
	"github.com/SchumacherFM/GoPlayground/db/sqlboilr"
	"github.com/corestoreio/csfw/storage/dbr"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kisielk/sqlstruct"
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
		if err := rows.Close(); err != nil {
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
		if err := rows.Close(); err != nil {
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
		if err := rows.Close(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSQLx(b *testing.B) {
	db, err := sqlx.Connect("mysql", "magento-1-8:magento-1-8@tcp(:3306)/magento-1-8")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()
	b.Run("append", func(b *testing.B) {
		var sqlxCPEV []*CPEV
		for i := 0; i < b.N; i++ {
			rows, err := db.Queryx(TEST_QUERY)
			if err != nil {
				b.Fatal(err)
			}
			for rows.Next() {
				cpev := new(CPEV)
				if err := rows.StructScan(cpev); err != nil {
					b.Fatal(err)
				}
				sqlxCPEV = append(sqlxCPEV, cpev)
			}
			sqlxCPEV = sqlxCPEV[:0]
		}
	})

	b.Run("select", func(b *testing.B) {
		cpev := []CPEV{}
		for i := 0; i < b.N; i++ {
			if err := db.Select(&cpev, TEST_QUERY); err != nil {
				b.Fatal(err)
			}
			cpev = cpev[:0]
		}
	})
}

func BenchmarkCSFWdbr(b *testing.B) {

	dbc := dbr.MustConnectAndVerify(dbr.WithDSN("magento-1-8:magento-1-8@tcp(:3306)/magento-1-8"))
	ctx := context.Background()

	defer dbc.Close()
	b.ResetTimer()

	cpc := new(CPEVCollection)
	for i := 0; i < b.N; i++ {
		if _, err := dbr.Load(ctx, dbc.DB, cpc, cpc); err != nil {
			b.Fatal(err)
		}
		cpc.Data = cpc.Data[:0]
	}
}

func BenchmarkSqlStruct(b *testing.B) {

	db, err := sql.Open("mysql", "magento-1-8:magento-1-8@tcp(:3306)/magento-1-8")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()
	b.ResetTimer()

	var sqlxCPEV []*CPEV
	for i := 0; i < b.N; i++ {

		rows, err := db.Query(TEST_QUERY)
		if err != nil {
			b.Fatal(err)
		}

		for rows.Next() {
			cpev := new(CPEV)
			sqlstruct.Scan(cpev, rows)
			if err != nil {
				b.Fatal(err)
			}
			sqlxCPEV = append(sqlxCPEV, cpev)
		}
		if err = rows.Err(); err != nil {
			b.Fatal(err)
		}
		if err := rows.Close(); err != nil {
			b.Fatal(err)
		}
		sqlxCPEV = sqlxCPEV[:0]
	}
}

func BenchmarkSqlBoiler(b *testing.B) {
	db, err := sql.Open("mysql", "magento-1-8:magento-1-8@tcp(:3306)/test")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()
	b.ResetTimer()

	b.Run("all", func(b *testing.B) {
		cpev := sqlboilr.CatalogProductEntityVarchars(db)
		for i := 0; i < b.N; i++ {
			res, err := cpev.All()
			if err != nil {
				b.Fatal(err)
			}
			res = res[:0]
		}
	})

}
