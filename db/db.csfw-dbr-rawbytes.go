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
	Count   uint
	Init    bool
	Columns []string
	// Alias maps a `key` containing the alias name used in the query to the
	// `value` the original snake case name used in the struct.
	Alias map[string]string
}

func (vs *CPEVRawCollection) ToSQL() (string, []interface{}, error) {
	return TEST_QUERY, nil, nil
}

// RowScan implements dbr.Scanner interface and scans a single row from the
// database query result.
func (vs *CPEVRawCollection) RowScan(r *sql.Rows) error {
	if !vs.Init {
		var err error
		vs.Columns, err = r.Columns()
		if err != nil {
			return err
		}
		vs.scanRaw = make([]*sql.RawBytes, len(vs.Columns))
		vs.scanArg = make([]interface{}, len(vs.Columns))
		for i := range vs.Columns {
			vs.scanRaw[i] = new(sql.RawBytes)
			vs.scanArg[i] = vs.scanRaw[i]
		}
		vs.Init = true
		vs.Count = 0
	}
	if err := r.Scan(vs.scanArg...); err != nil {
		return err
	}

	o := new(CPEV)
	for i, col := range vs.Columns {
		if vs.Alias != nil {
			if orgCol, ok := vs.Alias[col]; ok {
				col = orgCol
			}
		}
		// TODO err checking
		switch col {
		case "value_id":
			o.ValueID, _ = byteconv.ParseIntSQL(vs.scanRaw[i])
		case "entity_type_id":
			uid, _, _ := byteconv.ParseUintSQL(vs.scanRaw[i], 10, 32)
			o.EntityTypeId = uint(uid)
		case "attribute_id":
			uid, _, _ := byteconv.ParseUintSQL(vs.scanRaw[i], 10, 16)
			o.AttributeId = uint16(uid)
		case "store_id":
			uid, _, _ := byteconv.ParseUintSQL(vs.scanRaw[i], 10, 16)
			o.StoreId = uint16(uid)
		case "entity_id":
			uid, _, _ := byteconv.ParseUintSQL(vs.scanRaw[i], 10, 32)
			o.EntityId = uint(uid)
		case "value":
			o.Value = byteconv.ParseNullStringSQL(vs.scanRaw[i])
		}
	}

	//fmt.Printf("%#v \n\n", o)
	//if vs.Count > 5 {
	//	panic("game over")
	//}

	vs.Data = append(vs.Data, o)
	vs.Count++
	return nil
}
