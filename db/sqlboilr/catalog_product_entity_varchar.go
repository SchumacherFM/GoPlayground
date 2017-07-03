// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package sqlboilr

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v6"
)

// CatalogProductEntityVarchar is an object representing the database table.
type CatalogProductEntityVarchar struct {
	ValueID      int         `boil:"value_id" json:"value_id" toml:"value_id" yaml:"value_id"`
	EntityTypeID uint        `boil:"entity_type_id" json:"entity_type_id" toml:"entity_type_id" yaml:"entity_type_id"`
	AttributeID  uint16      `boil:"attribute_id" json:"attribute_id" toml:"attribute_id" yaml:"attribute_id"`
	StoreID      uint16      `boil:"store_id" json:"store_id" toml:"store_id" yaml:"store_id"`
	EntityID     uint        `boil:"entity_id" json:"entity_id" toml:"entity_id" yaml:"entity_id"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *catalogProductEntityVarcharR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L catalogProductEntityVarcharL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CatalogProductEntityVarcharColumns = struct {
	ValueID      string
	EntityTypeID string
	AttributeID  string
	StoreID      string
	EntityID     string
	Value        string
}{
	ValueID:      "value_id",
	EntityTypeID: "entity_type_id",
	AttributeID:  "attribute_id",
	StoreID:      "store_id",
	EntityID:     "entity_id",
	Value:        "value",
}

// catalogProductEntityVarcharR is where relationships are stored.
type catalogProductEntityVarcharR struct {
}

// catalogProductEntityVarcharL is where Load methods for each relationship are stored.
type catalogProductEntityVarcharL struct{}

var (
	catalogProductEntityVarcharColumns               = []string{"value_id", "entity_type_id", "attribute_id", "store_id", "entity_id", "value"}
	catalogProductEntityVarcharColumnsWithoutDefault = []string{"entity_type_id", "attribute_id", "store_id", "entity_id", "value"}
	catalogProductEntityVarcharColumnsWithDefault    = []string{"value_id"}
	catalogProductEntityVarcharPrimaryKeyColumns     = []string{"value_id"}
)

type (
	// CatalogProductEntityVarcharSlice is an alias for a slice of pointers to CatalogProductEntityVarchar.
	// This should generally be used opposed to []CatalogProductEntityVarchar.
	CatalogProductEntityVarcharSlice []*CatalogProductEntityVarchar

	catalogProductEntityVarcharQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	catalogProductEntityVarcharType                 = reflect.TypeOf(&CatalogProductEntityVarchar{})
	catalogProductEntityVarcharMapping              = queries.MakeStructMapping(catalogProductEntityVarcharType)
	catalogProductEntityVarcharPrimaryKeyMapping, _ = queries.BindMapping(catalogProductEntityVarcharType, catalogProductEntityVarcharMapping, catalogProductEntityVarcharPrimaryKeyColumns)
	catalogProductEntityVarcharInsertCacheMut       sync.RWMutex
	catalogProductEntityVarcharInsertCache          = make(map[string]insertCache)
	catalogProductEntityVarcharUpdateCacheMut       sync.RWMutex
	catalogProductEntityVarcharUpdateCache          = make(map[string]updateCache)
	catalogProductEntityVarcharUpsertCacheMut       sync.RWMutex
	catalogProductEntityVarcharUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single catalogProductEntityVarchar record from the query, and panics on error.
func (q catalogProductEntityVarcharQuery) OneP() *CatalogProductEntityVarchar {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single catalogProductEntityVarchar record from the query.
func (q catalogProductEntityVarcharQuery) One() (*CatalogProductEntityVarchar, error) {
	o := &CatalogProductEntityVarchar{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboilr: failed to execute a one query for catalog_product_entity_varchar")
	}

	return o, nil
}

// AllP returns all CatalogProductEntityVarchar records from the query, and panics on error.
func (q catalogProductEntityVarcharQuery) AllP() CatalogProductEntityVarcharSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all CatalogProductEntityVarchar records from the query.
func (q catalogProductEntityVarcharQuery) All(o ...*CatalogProductEntityVarchar) (CatalogProductEntityVarcharSlice, error) {
	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlboilr: failed to assign all query results to CatalogProductEntityVarchar slice")
	}

	return o, nil
}

// CountP returns the count of all CatalogProductEntityVarchar records in the query, and panics on error.
func (q catalogProductEntityVarcharQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all CatalogProductEntityVarchar records in the query.
func (q catalogProductEntityVarcharQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilr: failed to count catalog_product_entity_varchar rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q catalogProductEntityVarcharQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q catalogProductEntityVarcharQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlboilr: failed to check if catalog_product_entity_varchar exists")
	}

	return count > 0, nil
}

// CatalogProductEntityVarcharsG retrieves all records.
func CatalogProductEntityVarcharsG(mods ...qm.QueryMod) catalogProductEntityVarcharQuery {
	return CatalogProductEntityVarchars(boil.GetDB(), mods...)
}

// CatalogProductEntityVarchars retrieves all the records using an executor.
func CatalogProductEntityVarchars(exec boil.Executor, mods ...qm.QueryMod) catalogProductEntityVarcharQuery {
	mods = append(mods, qm.From("`catalog_product_entity_varchar`"))
	return catalogProductEntityVarcharQuery{NewQuery(exec, mods...)}
}

// FindCatalogProductEntityVarcharG retrieves a single record by ID.
func FindCatalogProductEntityVarcharG(valueID int, selectCols ...string) (*CatalogProductEntityVarchar, error) {
	return FindCatalogProductEntityVarchar(boil.GetDB(), valueID, selectCols...)
}

// FindCatalogProductEntityVarcharGP retrieves a single record by ID, and panics on error.
func FindCatalogProductEntityVarcharGP(valueID int, selectCols ...string) *CatalogProductEntityVarchar {
	retobj, err := FindCatalogProductEntityVarchar(boil.GetDB(), valueID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCatalogProductEntityVarchar retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCatalogProductEntityVarchar(exec boil.Executor, valueID int, selectCols ...string) (*CatalogProductEntityVarchar, error) {
	catalogProductEntityVarcharObj := &CatalogProductEntityVarchar{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `catalog_product_entity_varchar` where `value_id`=?", sel,
	)

	q := queries.Raw(exec, query, valueID)

	err := q.Bind(catalogProductEntityVarcharObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboilr: unable to select from catalog_product_entity_varchar")
	}

	return catalogProductEntityVarcharObj, nil
}

// FindCatalogProductEntityVarcharP retrieves a single record by ID with an executor, and panics on error.
func FindCatalogProductEntityVarcharP(exec boil.Executor, valueID int, selectCols ...string) *CatalogProductEntityVarchar {
	retobj, err := FindCatalogProductEntityVarchar(exec, valueID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *CatalogProductEntityVarchar) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *CatalogProductEntityVarchar) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *CatalogProductEntityVarchar) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *CatalogProductEntityVarchar) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("sqlboilr: no catalog_product_entity_varchar provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(catalogProductEntityVarcharColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	catalogProductEntityVarcharInsertCacheMut.RLock()
	cache, cached := catalogProductEntityVarcharInsertCache[key]
	catalogProductEntityVarcharInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			catalogProductEntityVarcharColumns,
			catalogProductEntityVarcharColumnsWithDefault,
			catalogProductEntityVarcharColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(catalogProductEntityVarcharType, catalogProductEntityVarcharMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(catalogProductEntityVarcharType, catalogProductEntityVarcharMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `catalog_product_entity_varchar` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `catalog_product_entity_varchar` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `catalog_product_entity_varchar` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, catalogProductEntityVarcharPrimaryKeyColumns))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "sqlboilr: unable to insert into catalog_product_entity_varchar")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ValueID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == catalogProductEntityVarcharMapping["ValueID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ValueID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "sqlboilr: unable to populate default values for catalog_product_entity_varchar")
	}

CacheNoHooks:
	if !cached {
		catalogProductEntityVarcharInsertCacheMut.Lock()
		catalogProductEntityVarcharInsertCache[key] = cache
		catalogProductEntityVarcharInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single CatalogProductEntityVarchar record. See Update for
// whitelist behavior description.
func (o *CatalogProductEntityVarchar) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single CatalogProductEntityVarchar record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *CatalogProductEntityVarchar) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the CatalogProductEntityVarchar, and panics on error.
// See Update for whitelist behavior description.
func (o *CatalogProductEntityVarchar) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the CatalogProductEntityVarchar.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *CatalogProductEntityVarchar) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	catalogProductEntityVarcharUpdateCacheMut.RLock()
	cache, cached := catalogProductEntityVarcharUpdateCache[key]
	catalogProductEntityVarcharUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			catalogProductEntityVarcharColumns,
			catalogProductEntityVarcharPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("sqlboilr: unable to update catalog_product_entity_varchar, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `catalog_product_entity_varchar` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, catalogProductEntityVarcharPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(catalogProductEntityVarcharType, catalogProductEntityVarcharMapping, append(wl, catalogProductEntityVarcharPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "sqlboilr: unable to update catalog_product_entity_varchar row")
	}

	if !cached {
		catalogProductEntityVarcharUpdateCacheMut.Lock()
		catalogProductEntityVarcharUpdateCache[key] = cache
		catalogProductEntityVarcharUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q catalogProductEntityVarcharQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q catalogProductEntityVarcharQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "sqlboilr: unable to update all for catalog_product_entity_varchar")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CatalogProductEntityVarcharSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CatalogProductEntityVarcharSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CatalogProductEntityVarcharSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CatalogProductEntityVarcharSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("sqlboilr: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), catalogProductEntityVarcharPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `catalog_product_entity_varchar` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, catalogProductEntityVarcharPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "sqlboilr: unable to update all in catalogProductEntityVarchar slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *CatalogProductEntityVarchar) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *CatalogProductEntityVarchar) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *CatalogProductEntityVarchar) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *CatalogProductEntityVarchar) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("sqlboilr: no catalog_product_entity_varchar provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(catalogProductEntityVarcharColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	catalogProductEntityVarcharUpsertCacheMut.RLock()
	cache, cached := catalogProductEntityVarcharUpsertCache[key]
	catalogProductEntityVarcharUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			catalogProductEntityVarcharColumns,
			catalogProductEntityVarcharColumnsWithDefault,
			catalogProductEntityVarcharColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			catalogProductEntityVarcharColumns,
			catalogProductEntityVarcharPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("sqlboilr: unable to upsert catalog_product_entity_varchar, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "catalog_product_entity_varchar", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `catalog_product_entity_varchar` WHERE `value_id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(catalogProductEntityVarcharType, catalogProductEntityVarcharMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(catalogProductEntityVarcharType, catalogProductEntityVarcharMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "sqlboilr: unable to upsert for catalog_product_entity_varchar")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ValueID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == catalogProductEntityVarcharMapping["ValueID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ValueID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "sqlboilr: unable to populate default values for catalog_product_entity_varchar")
	}

CacheNoHooks:
	if !cached {
		catalogProductEntityVarcharUpsertCacheMut.Lock()
		catalogProductEntityVarcharUpsertCache[key] = cache
		catalogProductEntityVarcharUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single CatalogProductEntityVarchar record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CatalogProductEntityVarchar) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single CatalogProductEntityVarchar record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *CatalogProductEntityVarchar) DeleteG() error {
	if o == nil {
		return errors.New("sqlboilr: no CatalogProductEntityVarchar provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single CatalogProductEntityVarchar record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CatalogProductEntityVarchar) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single CatalogProductEntityVarchar record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CatalogProductEntityVarchar) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("sqlboilr: no CatalogProductEntityVarchar provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), catalogProductEntityVarcharPrimaryKeyMapping)
	sql := "DELETE FROM `catalog_product_entity_varchar` WHERE `value_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "sqlboilr: unable to delete from catalog_product_entity_varchar")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q catalogProductEntityVarcharQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q catalogProductEntityVarcharQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("sqlboilr: no catalogProductEntityVarcharQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "sqlboilr: unable to delete all from catalog_product_entity_varchar")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CatalogProductEntityVarcharSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CatalogProductEntityVarcharSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("sqlboilr: no CatalogProductEntityVarchar slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CatalogProductEntityVarcharSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CatalogProductEntityVarcharSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("sqlboilr: no CatalogProductEntityVarchar slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), catalogProductEntityVarcharPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `catalog_product_entity_varchar` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, catalogProductEntityVarcharPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "sqlboilr: unable to delete all from catalogProductEntityVarchar slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *CatalogProductEntityVarchar) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *CatalogProductEntityVarchar) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *CatalogProductEntityVarchar) ReloadG() error {
	if o == nil {
		return errors.New("sqlboilr: no CatalogProductEntityVarchar provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CatalogProductEntityVarchar) Reload(exec boil.Executor) error {
	ret, err := FindCatalogProductEntityVarchar(exec, o.ValueID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CatalogProductEntityVarcharSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CatalogProductEntityVarcharSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CatalogProductEntityVarcharSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("sqlboilr: empty CatalogProductEntityVarcharSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CatalogProductEntityVarcharSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	catalogProductEntityVarchars := CatalogProductEntityVarcharSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), catalogProductEntityVarcharPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `catalog_product_entity_varchar`.* FROM `catalog_product_entity_varchar` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, catalogProductEntityVarcharPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&catalogProductEntityVarchars)
	if err != nil {
		return errors.Wrap(err, "sqlboilr: unable to reload all in CatalogProductEntityVarcharSlice")
	}

	*o = catalogProductEntityVarchars

	return nil
}

// CatalogProductEntityVarcharExists checks if the CatalogProductEntityVarchar row exists.
func CatalogProductEntityVarcharExists(exec boil.Executor, valueID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `catalog_product_entity_varchar` where `value_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, valueID)
	}

	row := exec.QueryRow(sql, valueID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlboilr: unable to check if catalog_product_entity_varchar exists")
	}

	return exists, nil
}

// CatalogProductEntityVarcharExistsG checks if the CatalogProductEntityVarchar row exists.
func CatalogProductEntityVarcharExistsG(valueID int) (bool, error) {
	return CatalogProductEntityVarcharExists(boil.GetDB(), valueID)
}

// CatalogProductEntityVarcharExistsGP checks if the CatalogProductEntityVarchar row exists. Panics on error.
func CatalogProductEntityVarcharExistsGP(valueID int) bool {
	e, err := CatalogProductEntityVarcharExists(boil.GetDB(), valueID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CatalogProductEntityVarcharExistsP checks if the CatalogProductEntityVarchar row exists. Panics on error.
func CatalogProductEntityVarcharExistsP(exec boil.Executor, valueID int) bool {
	e, err := CatalogProductEntityVarcharExists(exec, valueID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
