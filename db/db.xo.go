// Package db contains the types for schema 'test'.
package main

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"database/sql/driver"
	"encoding/csv"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// CatalogProductEntityVarchar represents a row from 'test.catalog_product_entity_varchar'.
type CatalogProductEntityVarchar struct {
	ValueID      int            `json:"value_id"`       // value_id
	EntityTypeID uint           `json:"entity_type_id"` // entity_type_id
	AttributeID  int16          `json:"attribute_id"`   // attribute_id
	StoreID      int16          `json:"store_id"`       // store_id
	EntityID     uint           `json:"entity_id"`      // entity_id
	Value        sql.NullString `json:"value"`          // value

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the CatalogProductEntityVarchar exists in the database.
func (cpev *CatalogProductEntityVarchar) Exists() bool {
	return cpev._exists
}

// Deleted provides information if the CatalogProductEntityVarchar has been deleted from the database.
func (cpev *CatalogProductEntityVarchar) Deleted() bool {
	return cpev._deleted
}

// Insert inserts the CatalogProductEntityVarchar to the database.
func (cpev *CatalogProductEntityVarchar) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if cpev._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO test.catalog_product_entity_varchar (` +
		`entity_type_id, attribute_id, store_id, entity_id, value` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, cpev.EntityTypeID, cpev.AttributeID, cpev.StoreID, cpev.EntityID, cpev.Value)
	res, err := db.Exec(sqlstr, cpev.EntityTypeID, cpev.AttributeID, cpev.StoreID, cpev.EntityID, cpev.Value)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	cpev.ValueID = int(id)
	cpev._exists = true

	return nil
}

// Update updates the CatalogProductEntityVarchar in the database.
func (cpev *CatalogProductEntityVarchar) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !cpev._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if cpev._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE test.catalog_product_entity_varchar SET ` +
		`entity_type_id = ?, attribute_id = ?, store_id = ?, entity_id = ?, value = ?` +
		` WHERE value_id = ?`

	// run query
	XOLog(sqlstr, cpev.EntityTypeID, cpev.AttributeID, cpev.StoreID, cpev.EntityID, cpev.Value, cpev.ValueID)
	_, err = db.Exec(sqlstr, cpev.EntityTypeID, cpev.AttributeID, cpev.StoreID, cpev.EntityID, cpev.Value, cpev.ValueID)
	return err
}

// Save saves the CatalogProductEntityVarchar to the database.
func (cpev *CatalogProductEntityVarchar) Save(db XODB) error {
	if cpev.Exists() {
		return cpev.Update(db)
	}

	return cpev.Insert(db)
}

// Delete deletes the CatalogProductEntityVarchar from the database.
func (cpev *CatalogProductEntityVarchar) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !cpev._exists {
		return nil
	}

	// if deleted, bail
	if cpev._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM test.catalog_product_entity_varchar WHERE value_id = ?`

	// run query
	XOLog(sqlstr, cpev.ValueID)
	_, err = db.Exec(sqlstr, cpev.ValueID)
	if err != nil {
		return err
	}

	// set deleted
	cpev._deleted = true

	return nil
}

// CatalogProductEntityVarcharsByAttributeID retrieves a row from 'test.catalog_product_entity_varchar' as a CatalogProductEntityVarchar.
//
// Generated from index 'IDX_CATALOG_PRODUCT_ENTITY_VARCHAR_ATTRIBUTE_ID'.
func CatalogProductEntityVarcharsByAttributeID(db XODB, attributeID int16) ([]*CatalogProductEntityVarchar, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`value_id, entity_type_id, attribute_id, store_id, entity_id, value ` +
		`FROM test.catalog_product_entity_varchar ` +
		`WHERE attribute_id = ?`

	// run query
	XOLog(sqlstr, attributeID)
	q, err := db.Query(sqlstr, attributeID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*CatalogProductEntityVarchar{}
	for q.Next() {
		cpev := CatalogProductEntityVarchar{
			_exists: true,
		}

		// scan
		err = q.Scan(&cpev.ValueID, &cpev.EntityTypeID, &cpev.AttributeID, &cpev.StoreID, &cpev.EntityID, &cpev.Value)
		if err != nil {
			return nil, err
		}

		res = append(res, &cpev)
	}

	return res, nil
}

// CatalogProductEntityVarcharsByEntityID retrieves a row from 'test.catalog_product_entity_varchar' as a CatalogProductEntityVarchar.
//
// Generated from index 'IDX_CATALOG_PRODUCT_ENTITY_VARCHAR_ENTITY_ID'.
func CatalogProductEntityVarcharsByEntityID(db XODB, entityID uint) ([]*CatalogProductEntityVarchar, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`value_id, entity_type_id, attribute_id, store_id, entity_id, value ` +
		`FROM test.catalog_product_entity_varchar ` +
		`WHERE entity_id = ?`

	// run query
	XOLog(sqlstr, entityID)
	q, err := db.Query(sqlstr, entityID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*CatalogProductEntityVarchar{}
	for q.Next() {
		cpev := CatalogProductEntityVarchar{
			_exists: true,
		}

		// scan
		err = q.Scan(&cpev.ValueID, &cpev.EntityTypeID, &cpev.AttributeID, &cpev.StoreID, &cpev.EntityID, &cpev.Value)
		if err != nil {
			return nil, err
		}

		res = append(res, &cpev)
	}

	return res, nil
}

// CatalogProductEntityVarcharsByStoreID retrieves a row from 'test.catalog_product_entity_varchar' as a CatalogProductEntityVarchar.
//
// Generated from index 'IDX_CATALOG_PRODUCT_ENTITY_VARCHAR_STORE_ID'.
func CatalogProductEntityVarcharsByStoreID(db XODB, storeID int16) ([]*CatalogProductEntityVarchar, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`value_id, entity_type_id, attribute_id, store_id, entity_id, value ` +
		`FROM test.catalog_product_entity_varchar ` +
		`WHERE store_id = ?`

	// run query
	XOLog(sqlstr, storeID)
	q, err := db.Query(sqlstr, storeID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*CatalogProductEntityVarchar{}
	for q.Next() {
		cpev := CatalogProductEntityVarchar{
			_exists: true,
		}

		// scan
		err = q.Scan(&cpev.ValueID, &cpev.EntityTypeID, &cpev.AttributeID, &cpev.StoreID, &cpev.EntityID, &cpev.Value)
		if err != nil {
			return nil, err
		}

		res = append(res, &cpev)
	}

	return res, nil
}

// Custom hack
func CatalogProductEntityVarcharsAll(db XODB, res ...*CatalogProductEntityVarchar) ([]*CatalogProductEntityVarchar, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`value_id, entity_type_id, attribute_id, store_id, entity_id, value ` +
		`FROM test.catalog_product_entity_varchar `

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		cpev := CatalogProductEntityVarchar{
			_exists: true,
		}
		// scan
		err = q.Scan(&cpev.ValueID, &cpev.EntityTypeID, &cpev.AttributeID, &cpev.StoreID, &cpev.EntityID, &cpev.Value)
		if err != nil {
			return nil, err
		}
		res = append(res, &cpev)
	}

	return res, nil
}

// CatalogProductEntityVarcharByEntityIDAttributeIDStoreID retrieves a row from 'test.catalog_product_entity_varchar' as a CatalogProductEntityVarchar.
//
// Generated from index 'UNQ_CAT_PRD_ENTT_VCHR_ENTT_ID_ATTR_ID_STORE_ID'.
func CatalogProductEntityVarcharByEntityIDAttributeIDStoreID(db XODB, entityID uint, attributeID int16, storeID int16) (*CatalogProductEntityVarchar, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`value_id, entity_type_id, attribute_id, store_id, entity_id, value ` +
		`FROM test.catalog_product_entity_varchar ` +
		`WHERE entity_id = ? AND attribute_id = ? AND store_id = ?`

	// run query
	XOLog(sqlstr, entityID, attributeID, storeID)
	cpev := CatalogProductEntityVarchar{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, entityID, attributeID, storeID).Scan(&cpev.ValueID, &cpev.EntityTypeID, &cpev.AttributeID, &cpev.StoreID, &cpev.EntityID, &cpev.Value)
	if err != nil {
		return nil, err
	}

	return &cpev, nil
}

// CatalogProductEntityVarcharByValueID retrieves a row from 'test.catalog_product_entity_varchar' as a CatalogProductEntityVarchar.
//
// Generated from index 'catalog_product_entity_varchar_value_id_pkey'.
func CatalogProductEntityVarcharByValueID(db XODB, valueID int) (*CatalogProductEntityVarchar, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`value_id, entity_type_id, attribute_id, store_id, entity_id, value ` +
		`FROM test.catalog_product_entity_varchar ` +
		`WHERE value_id = ?`

	// run query
	XOLog(sqlstr, valueID)
	cpev := CatalogProductEntityVarchar{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, valueID).Scan(&cpev.ValueID, &cpev.EntityTypeID, &cpev.AttributeID, &cpev.StoreID, &cpev.EntityID, &cpev.Value)
	if err != nil {
		return nil, err
	}

	return &cpev, nil
}

// XODB is the common interface for database operations that can be used with
// types from schema 'test'.
//
// This should work with database/sql.DB and database/sql.Tx.
type XODB interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

// XOLog provides the log func used by generated queries.
var XOLog = func(string, ...interface{}) {}

// ScannerValuer is the common interface for types that implement both the
// database/sql.Scanner and sql/driver.Valuer interfaces.
type ScannerValuer interface {
	sql.Scanner
	driver.Valuer
}

// StringSlice is a slice of strings.
type StringSlice []string

// quoteEscapeRegex is the regex to match escaped characters in a string.
var quoteEscapeRegex = regexp.MustCompile(`([^\\]([\\]{2})*)\\"`)

// Scan satisfies the sql.Scanner interface for StringSlice.
func (ss *StringSlice) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid StringSlice")
	}

	// change quote escapes for csv parser
	str := quoteEscapeRegex.ReplaceAllString(string(buf), `$1""`)
	str = strings.Replace(str, `\\`, `\`, -1)

	// remove braces
	str = str[1 : len(str)-1]

	// bail if only one
	if len(str) == 0 {
		*ss = StringSlice([]string{})
		return nil
	}

	// parse with csv reader
	cr := csv.NewReader(strings.NewReader(str))
	slice, err := cr.Read()
	if err != nil {
		fmt.Printf("exiting!: %v\n", err)
		return err
	}

	*ss = StringSlice(slice)

	return nil
}

// Value satisfies the driver.Valuer interface for StringSlice.
func (ss StringSlice) Value() (driver.Value, error) {
	v := make([]string, len(ss))
	for i, s := range ss {
		v[i] = `"` + strings.Replace(strings.Replace(s, `\`, `\\\`, -1), `"`, `\"`, -1) + `"`
	}
	return "{" + strings.Join(v, ",") + "}", nil
}

// Slice is a slice of ScannerValuers.
type Slice []ScannerValuer