package main

import (
	"database/sql"

	"github.com/corestoreio/csfw/util/byteconv"
)

//type CPEV struct {
//	/*
//	  `value_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Value ID',
//	  `entity_type_id` int(10) unsigned NOT NULL COMMENT 'Entity Type ID',
//	  `attribute_id` smallint(5) unsigned NOT NULL COMMENT 'Attribute ID',
//	  `store_id` smallint(5) unsigned NOT NULL COMMENT 'Store ID',
//	  `entity_id` int(10) unsigned NOT NULL COMMENT 'Entity ID',
//	  `value` varchar(255) DEFAULT NULL COMMENT 'Value',
//	*/
//	ValueID      int64          `sql:"value_id" db:"value_id"`
//	EntityTypeId uint           `sql:"entity_type_id" db:"entity_type_id"`
//	AttributeId  uint16         `sql:"attribute_id" db:"attribute_id"`
//	StoreId      uint16         `sql:"store_id" db:"store_id"`
//	EntityId     uint           `sql:"entity_id" db:"entity_id"`
//	Value        sql.NullString `sql:"value" db:"value"`
//}

type CPEVRawCollection struct {
	Data    []*CPEV
	scanRaw []*sql.RawBytes
	scanArg []interface{}
	first   bool
	c       uint
}

func (vs *CPEVRawCollection) ToSQL() (string, []interface{}, error) {
	return TEST_QUERY, nil, nil
}

// RowScan implements dbr.Scanner interface and scans a single row from the
// database query result.
func (vs *CPEVRawCollection) RowScan(r *sql.Rows) error {
	if !vs.first {
		cols, err := r.Columns()
		if err != nil {
			return err
		}
		vs.scanRaw = make([]*sql.RawBytes, len(cols))
		vs.scanArg = make([]interface{}, len(cols))
		for i := range vs.scanArg {
			vs.scanRaw[i] = new(sql.RawBytes)
			vs.scanArg[i] = vs.scanRaw[i]
		}
		vs.first = true
	}
	if err := r.Scan(vs.scanArg...); err != nil {
		return err
	}

	var o CPEV

	// TODO err checking
	o.ValueID, _, _ = byteconv.ParseIntSQL(vs.scanRaw[0])
	uid, _, _ := byteconv.ParseUintSQL(vs.scanRaw[1], 10, 32)
	o.EntityTypeId = uint(uid)
	uid, _, _ = byteconv.ParseUintSQL(vs.scanRaw[2], 10, 16)
	o.AttributeId = uint16(uid)
	uid, _, _ = byteconv.ParseUintSQL(vs.scanRaw[3], 10, 16)
	o.StoreId = uint16(uid)
	uid, _, _ = byteconv.ParseUintSQL(vs.scanRaw[4], 10, 32)
	o.EntityId = uint(uid)

	o.Value = byteconv.ParseStringSQL(vs.scanRaw[5])

	//fmt.Printf("%#v => %s\n\n", *(vs.scanRaw[2]), *(vs.scanRaw[2]))
	//if vs.c > 5 {
	//	panic("game over")
	//}

	vs.Data = append(vs.Data, &o)
	vs.c++
	return nil
}
