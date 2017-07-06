package main

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type cPEVTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *cPEVTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("catalog_product_entity_varchar").
func (v *cPEVTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *cPEVTableType) Columns() []string {
	return []string{"value_id", "entity_type_id", "attribute_id", "store_id", "entity_id", "value"}
}

// NewStruct makes a new struct for that view or table.
func (v *cPEVTableType) NewStruct() reform.Struct {
	return new(CPEV)
}

// NewRecord makes a new record for that table.
func (v *cPEVTableType) NewRecord() reform.Record {
	return new(CPEV)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *cPEVTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// CPEVTable represents catalog_product_entity_varchar view or table in SQL database.
var CPEVTable = &cPEVTableType{
	s: parse.StructInfo{Type: "CPEV", SQLSchema: "", SQLName: "catalog_product_entity_varchar", Fields: []parse.FieldInfo{{Name: "ValueID", PKType: "int64", Column: "value_id"}, {Name: "EntityTypeId", PKType: "", Column: "entity_type_id"}, {Name: "AttributeId", PKType: "", Column: "attribute_id"}, {Name: "StoreId", PKType: "", Column: "store_id"}, {Name: "EntityId", PKType: "", Column: "entity_id"}, {Name: "Value", PKType: "", Column: "value"}}, PKFieldIndex: 0},
	z: new(CPEV).Values(),
}

// String returns a string representation of this struct or record.
func (s CPEV) String() string {
	res := make([]string, 6)
	res[0] = "ValueID: " + reform.Inspect(s.ValueID, true)
	res[1] = "EntityTypeId: " + reform.Inspect(s.EntityTypeId, true)
	res[2] = "AttributeId: " + reform.Inspect(s.AttributeId, true)
	res[3] = "StoreId: " + reform.Inspect(s.StoreId, true)
	res[4] = "EntityId: " + reform.Inspect(s.EntityId, true)
	res[5] = "Value: " + reform.Inspect(s.Value, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *CPEV) Values() []interface{} {
	return []interface{}{
		s.ValueID,
		s.EntityTypeId,
		s.AttributeId,
		s.StoreId,
		s.EntityId,
		s.Value,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *CPEV) Pointers() []interface{} {
	return []interface{}{
		&s.ValueID,
		&s.EntityTypeId,
		&s.AttributeId,
		&s.StoreId,
		&s.EntityId,
		&s.Value,
	}
}

// View returns View object for that struct.
func (s *CPEV) View() reform.View {
	return CPEVTable
}

// Table returns Table object for that record.
func (s *CPEV) Table() reform.Table {
	return CPEVTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *CPEV) PKValue() interface{} {
	return s.ValueID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *CPEV) PKPointer() interface{} {
	return &s.ValueID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *CPEV) HasPK() bool {
	return s.ValueID != CPEVTable.z[CPEVTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *CPEV) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ValueID = int64(i64)
	} else {
		s.ValueID = pk.(int64)
	}
}

// check interfaces
var (
	_ reform.View   = CPEVTable
	_ reform.Struct = new(CPEV)
	_ reform.Table  = CPEVTable
	_ reform.Record = new(CPEV)
	_ fmt.Stringer  = new(CPEV)
)

func init() {
	parse.AssertUpToDate(&CPEVTable.s, new(CPEV))
}
