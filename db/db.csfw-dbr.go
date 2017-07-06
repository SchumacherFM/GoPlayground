package main

import (
	"database/sql"
)

//reform:catalog_product_entity_varchar
type CPEV struct {
	/*
	  `value_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Value ID',
	  `entity_type_id` int(10) unsigned NOT NULL COMMENT 'Entity Type ID',
	  `attribute_id` smallint(5) unsigned NOT NULL COMMENT 'Attribute ID',
	  `store_id` smallint(5) unsigned NOT NULL COMMENT 'Store ID',
	  `entity_id` int(10) unsigned NOT NULL COMMENT 'Entity ID',
	  `value` varchar(255) DEFAULT NULL COMMENT 'Value',
	*/
	ValueID      int64          `sql:"value_id" db:"value_id" reform:"value_id,pk"`
	EntityTypeId uint           `sql:"entity_type_id" db:"entity_type_id" reform:"entity_type_id"`
	AttributeId  uint16         `sql:"attribute_id" db:"attribute_id" reform:"attribute_id"`
	StoreId      uint16         `sql:"store_id" db:"store_id" reform:"store_id"`
	EntityId     uint           `sql:"entity_id" db:"entity_id" reform:"entity_id"`
	Value        sql.NullString `sql:"value" db:"value" reform:"value"`
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
func (vs *CPEVCollection) RowScan(r *sql.Rows) error {
	var o CPEV
	if !vs.first {
		vs.scanArg = [6]interface{}{
			&vs.dto.ValueID, &vs.dto.EntityTypeId, &vs.dto.AttributeId,
			&vs.dto.StoreId, &vs.dto.EntityId, &vs.dto.Value,
		}
		vs.first = true
	}
	if err := r.Scan(vs.scanArg[:]...); err != nil {
		return err
	}
	o = vs.dto
	vs.Data = append(vs.Data, &o)
	return nil
}
