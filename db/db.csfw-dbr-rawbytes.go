package main

import (
	"database/sql"

	"github.com/corestoreio/csfw/storage/dbr"
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
	Base           dbr.Base
	Data           []*CPEV
	EventAfterScan []func(*CPEV)
}

func (vs *CPEVRawCollection) ToSQL() (string, []interface{}, error) {
	return TEST_QUERY, nil, nil
}

// RowScan implements dbr.Scanner interface and scans a single row from the
// database query result.
func (vs *CPEVRawCollection) RowScan(r *sql.Rows) error {
	if err := vs.Base.Scan(r); err != nil {
		return err
	}

	o := new(CPEV)
	for i, col := range vs.Base.Columns {
		if vs.Base.Alias != nil {
			if orgCol, ok := vs.Base.Alias[col]; ok {
				col = orgCol
			}
		}
		b := vs.Base.Index(i)
		var err error
		switch col {
		case "value_id":
			o.ValueID, err = b.Int64()
		case "entity_type_id":
			o.EntityTypeId, err = b.Uint()
		case "attribute_id":
			o.AttributeId, err = b.Uint16()
		case "store_id":
			o.StoreId, err = b.Uint16()
		case "entity_id":
			o.EntityId, err = b.Uint()
		case "value":
			o.Value, err = b.NullString()
		}
		if err != nil {
			return err
		}
	}

	for _, fn := range vs.EventAfterScan {
		fn(o)
	}

	vs.Data = append(vs.Data, o)
	return nil
}
