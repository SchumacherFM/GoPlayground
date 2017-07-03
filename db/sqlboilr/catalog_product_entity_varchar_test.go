// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package sqlboilr

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testCatalogProductEntityVarchars(t *testing.T) {
	t.Parallel()

	query := CatalogProductEntityVarchars(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCatalogProductEntityVarcharsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = catalogProductEntityVarchar.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := CatalogProductEntityVarchars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCatalogProductEntityVarcharsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = CatalogProductEntityVarchars(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := CatalogProductEntityVarchars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCatalogProductEntityVarcharsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CatalogProductEntityVarcharSlice{catalogProductEntityVarchar}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := CatalogProductEntityVarchars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCatalogProductEntityVarcharsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CatalogProductEntityVarcharExists(tx, catalogProductEntityVarchar.ValueID)
	if err != nil {
		t.Errorf("Unable to check if CatalogProductEntityVarchar exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CatalogProductEntityVarcharExistsG to return true, but got false.")
	}
}
func testCatalogProductEntityVarcharsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx); err != nil {
		t.Error(err)
	}

	catalogProductEntityVarcharFound, err := FindCatalogProductEntityVarchar(tx, catalogProductEntityVarchar.ValueID)
	if err != nil {
		t.Error(err)
	}

	if catalogProductEntityVarcharFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCatalogProductEntityVarcharsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = CatalogProductEntityVarchars(tx).Bind(catalogProductEntityVarchar); err != nil {
		t.Error(err)
	}
}

func testCatalogProductEntityVarcharsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := CatalogProductEntityVarchars(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCatalogProductEntityVarcharsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarcharOne := &CatalogProductEntityVarchar{}
	catalogProductEntityVarcharTwo := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarcharOne, catalogProductEntityVarcharDBTypes, false, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}
	if err = randomize.Struct(seed, catalogProductEntityVarcharTwo, catalogProductEntityVarcharDBTypes, false, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarcharOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = catalogProductEntityVarcharTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := CatalogProductEntityVarchars(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCatalogProductEntityVarcharsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	catalogProductEntityVarcharOne := &CatalogProductEntityVarchar{}
	catalogProductEntityVarcharTwo := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarcharOne, catalogProductEntityVarcharDBTypes, false, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}
	if err = randomize.Struct(seed, catalogProductEntityVarcharTwo, catalogProductEntityVarcharDBTypes, false, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarcharOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = catalogProductEntityVarcharTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CatalogProductEntityVarchars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testCatalogProductEntityVarcharsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CatalogProductEntityVarchars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCatalogProductEntityVarcharsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx, catalogProductEntityVarcharColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := CatalogProductEntityVarchars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCatalogProductEntityVarcharsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = catalogProductEntityVarchar.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCatalogProductEntityVarcharsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CatalogProductEntityVarcharSlice{catalogProductEntityVarchar}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCatalogProductEntityVarcharsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := CatalogProductEntityVarchars(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	catalogProductEntityVarcharDBTypes = map[string]string{`AttributeID`: `smallint`, `EntityID`: `int`, `EntityTypeID`: `int`, `StoreID`: `smallint`, `Value`: `varchar`, `ValueID`: `int`}
	_                                  = bytes.MinRead
)

func testCatalogProductEntityVarcharsUpdate(t *testing.T) {
	t.Parallel()

	if len(catalogProductEntityVarcharColumns) == len(catalogProductEntityVarcharPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CatalogProductEntityVarchars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	if err = catalogProductEntityVarchar.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCatalogProductEntityVarcharsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(catalogProductEntityVarcharColumns) == len(catalogProductEntityVarcharPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	catalogProductEntityVarchar := &CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CatalogProductEntityVarchars(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true, catalogProductEntityVarcharPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(catalogProductEntityVarcharColumns, catalogProductEntityVarcharPrimaryKeyColumns) {
		fields = catalogProductEntityVarcharColumns
	} else {
		fields = strmangle.SetComplement(
			catalogProductEntityVarcharColumns,
			catalogProductEntityVarcharPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(catalogProductEntityVarchar))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CatalogProductEntityVarcharSlice{catalogProductEntityVarchar}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCatalogProductEntityVarcharsUpsert(t *testing.T) {
	t.Parallel()

	if len(catalogProductEntityVarcharColumns) == len(catalogProductEntityVarcharPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	catalogProductEntityVarchar := CatalogProductEntityVarchar{}
	if err = randomize.Struct(seed, &catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catalogProductEntityVarchar.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert CatalogProductEntityVarchar: %s", err)
	}

	count, err := CatalogProductEntityVarchars(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &catalogProductEntityVarchar, catalogProductEntityVarcharDBTypes, false, catalogProductEntityVarcharPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CatalogProductEntityVarchar struct: %s", err)
	}

	if err = catalogProductEntityVarchar.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert CatalogProductEntityVarchar: %s", err)
	}

	count, err = CatalogProductEntityVarchars(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}