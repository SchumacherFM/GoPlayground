package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// 1745 rows
	// columns are: value_id, entity_type_id, attribute_id, store_id, entity_id, value
	TEST_QUERY = `SELECT * FROM catalog_product_entity_varchar`
)

func main() {
	db, err := sql.Open("mysql", "magento-1-8:magento-1-8@tcp(:3306)/magento-1-8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(TEST_QUERY)
	fck(err)
	defer rows.Close()
	columnNames, err := rows.Columns()
	fck(err)
	rc := NewMapStringScan(columnNames)
	for rows.Next() {
		//        cv, err := rowMapString(columnNames, rows)
		//        fck(err)
		err := rc.Update(rows)
		fck(err)
		cv := rc.Get()
		log.Printf("%#v\n\n", cv)
	}
	fck(rows.Err())
}

/**
  using a map
*/
type mapStringScan struct {
	// cp are the column pointers
	cp []interface{}
	// row contains the final result
	row      map[string]string
	colCount int
	colNames []string
}

func NewMapStringScan(columnNames []string) *mapStringScan {
	lenCN := len(columnNames)
	s := &mapStringScan{
		cp:       make([]interface{}, lenCN),
		row:      make(map[string]string, lenCN),
		colCount: lenCN,
		colNames: columnNames,
	}
	for i := 0; i < lenCN; i++ {
		s.cp[i] = new(sql.RawBytes)
	}
	return s
}

func (s *mapStringScan) Update(rows *sql.Rows) error {
	if err := rows.Scan(s.cp...); err != nil {
		return err
	}

	for i := 0; i < s.colCount; i++ {
		if rb, ok := s.cp[i].(*sql.RawBytes); ok {
			s.row[s.colNames[i]] = string(*rb)
			*rb = nil // reset pointer to discard current value to avoid a bug
		} else {
			return fmt.Errorf("Cannot convert index %d column %s to type *sql.RawBytes", i, s.colNames[i])
		}
	}
	return nil
}

func (s *mapStringScan) Get() map[string]string {
	return s.row
}

/**
  using a string slice
*/
type stringStringScan struct {
	// cp are the column pointers
	cp []interface{}
	// row contains the final result
	row      []string
	colCount int
	colNames []string
}

func NewStringStringScan(columnNames []string) *stringStringScan {
	lenCN := len(columnNames)
	s := &stringStringScan{
		cp:       make([]interface{}, lenCN),
		row:      make([]string, lenCN*2),
		colCount: lenCN,
		colNames: columnNames,
	}
	j := 0
	for i := 0; i < lenCN; i++ {
		s.cp[i] = new(sql.RawBytes)
		s.row[j] = s.colNames[i]
		j = j + 2
	}
	return s
}

func (s *stringStringScan) Update(rows *sql.Rows) error {
	if err := rows.Scan(s.cp...); err != nil {
		return err
	}
	j := 0
	for i := 0; i < s.colCount; i++ {
		if rb, ok := s.cp[i].(*sql.RawBytes); ok {
			s.row[j+1] = string(*rb)
			*rb = nil // reset pointer to discard current value to avoid a bug
		} else {
			return fmt.Errorf("Cannot convert index %d column %s to type *sql.RawBytes", i, s.colNames[i])
		}
		j = j + 2
	}
	return nil
}

func (s *stringStringScan) Get() []string {
	return s.row
}

// rowMapString was the first implementation but it creates for each row a new
// map and pointers and is considered as slow. see benchmark
func rowMapString(columnNames []string, rows *sql.Rows) (map[string]string, error) {
	lenCN := len(columnNames)
	ret := make(map[string]string, lenCN)

	columnPointers := make([]interface{}, lenCN)
	for i := 0; i < lenCN; i++ {
		columnPointers[i] = new(sql.RawBytes)
	}

	if err := rows.Scan(columnPointers...); err != nil {
		return nil, err
	}

	for i := 0; i < lenCN; i++ {
		if rb, ok := columnPointers[i].(*sql.RawBytes); ok {
			ret[columnNames[i]] = string(*rb)
		} else {
			return nil, fmt.Errorf("Cannot convert index %d column %s to type *sql.RawBytes", i, columnNames[i])
		}
	}

	return ret, nil
}

func fck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type CPEV struct {
	/*
	  `value_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Value ID',
	  `entity_type_id` int(10) unsigned NOT NULL COMMENT 'Entity Type ID',
	  `attribute_id` smallint(5) unsigned NOT NULL COMMENT 'Attribute ID',
	  `store_id` smallint(5) unsigned NOT NULL COMMENT 'Store ID',
	  `entity_id` int(10) unsigned NOT NULL COMMENT 'Entity ID',
	  `value` varchar(255) DEFAULT NULL COMMENT 'Value',
	*/
	ValueID      int64          `sql:"value_id" db:"value_id"`
	EntityTypeId uint           `sql:"entity_type_id" db:"entity_type_id"`
	AttributeId  uint16         `sql:"attribute_id" db:"attribute_id"`
	StoreId      uint16         `sql:"store_id" db:"store_id"`
	EntityId     uint           `sql:"entity_id" db:"entity_id"`
	Value        sql.NullString `sql:"value" db:"value"`
}

type CPEVCollection struct {
	Data    []*CPEV
	scanArg [6]interface{}
	dto     CPEV
	first   bool
}

func (vs *CPEVCollection) ToSQL() (string, []interface{}, error) {
	return TEST_QUERY, nil, nil
}

// RowScan implements dbr.Scanner interface and scans a single row from the
// database query result.
func (vs *CPEVCollection) RowScan(idx int64, _ []string, scan func(...interface{}) error) error {
	if !vs.first {
		vs.scanArg = [6]interface{}{
			&vs.dto.ValueID, &vs.dto.EntityTypeId, &vs.dto.AttributeId,
			&vs.dto.StoreId, &vs.dto.EntityId, &vs.dto.Value,
		}
		vs.first = true
	}
	if err := scan(vs.scanArg[:]...); err != nil {
		return err
	}
	o := vs.dto
	vs.Data = append(vs.Data, &o)
	return nil
}
