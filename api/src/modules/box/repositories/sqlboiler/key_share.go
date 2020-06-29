// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboiler

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// KeyShare is an object representing the database table.
type KeyShare struct {
	InvitationHash string    `boil:"invitation_hash" json:"invitation_hash" toml:"invitation_hash" yaml:"invitation_hash"`
	Share          string    `boil:"share" json:"share" toml:"share" yaml:"share"`
	CreatedAt      time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *keyShareR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L keyShareL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var KeyShareColumns = struct {
	InvitationHash string
	Share          string
	CreatedAt      string
}{
	InvitationHash: "invitation_hash",
	Share:          "share",
	CreatedAt:      "created_at",
}

// Generated where

var KeyShareWhere = struct {
	InvitationHash whereHelperstring
	Share          whereHelperstring
	CreatedAt      whereHelpertime_Time
}{
	InvitationHash: whereHelperstring{field: "\"key_share\".\"invitation_hash\""},
	Share:          whereHelperstring{field: "\"key_share\".\"share\""},
	CreatedAt:      whereHelpertime_Time{field: "\"key_share\".\"created_at\""},
}

// KeyShareRels is where relationship names are stored.
var KeyShareRels = struct {
}{}

// keyShareR is where relationships are stored.
type keyShareR struct {
}

// NewStruct creates a new relationship struct
func (*keyShareR) NewStruct() *keyShareR {
	return &keyShareR{}
}

// keyShareL is where Load methods for each relationship are stored.
type keyShareL struct{}

var (
	keyShareAllColumns            = []string{"invitation_hash", "share", "created_at"}
	keyShareColumnsWithoutDefault = []string{"invitation_hash", "share", "created_at"}
	keyShareColumnsWithDefault    = []string{}
	keySharePrimaryKeyColumns     = []string{"invitation_hash"}
)

type (
	// KeyShareSlice is an alias for a slice of pointers to KeyShare.
	// This should generally be used opposed to []KeyShare.
	KeyShareSlice []*KeyShare

	keyShareQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	keyShareType                 = reflect.TypeOf(&KeyShare{})
	keyShareMapping              = queries.MakeStructMapping(keyShareType)
	keySharePrimaryKeyMapping, _ = queries.BindMapping(keyShareType, keyShareMapping, keySharePrimaryKeyColumns)
	keyShareInsertCacheMut       sync.RWMutex
	keyShareInsertCache          = make(map[string]insertCache)
	keyShareUpdateCacheMut       sync.RWMutex
	keyShareUpdateCache          = make(map[string]updateCache)
	keyShareUpsertCacheMut       sync.RWMutex
	keyShareUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single keyShare record from the query.
func (q keyShareQuery) One(ctx context.Context, exec boil.ContextExecutor) (*KeyShare, error) {
	o := &KeyShare{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: failed to execute a one query for key_share")
	}

	return o, nil
}

// All returns all KeyShare records from the query.
func (q keyShareQuery) All(ctx context.Context, exec boil.ContextExecutor) (KeyShareSlice, error) {
	var o []*KeyShare

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlboiler: failed to assign all query results to KeyShare slice")
	}

	return o, nil
}

// Count returns the count of all KeyShare records in the query.
func (q keyShareQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to count key_share rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q keyShareQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: failed to check if key_share exists")
	}

	return count > 0, nil
}

// KeyShares retrieves all the records using an executor.
func KeyShares(mods ...qm.QueryMod) keyShareQuery {
	mods = append(mods, qm.From("\"key_share\""))
	return keyShareQuery{NewQuery(mods...)}
}

// FindKeyShare retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindKeyShare(ctx context.Context, exec boil.ContextExecutor, invitationHash string, selectCols ...string) (*KeyShare, error) {
	keyShareObj := &KeyShare{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"key_share\" where \"invitation_hash\"=$1", sel,
	)

	q := queries.Raw(query, invitationHash)

	err := q.Bind(ctx, exec, keyShareObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: unable to select from key_share")
	}

	return keyShareObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *KeyShare) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboiler: no key_share provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(keyShareColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	keyShareInsertCacheMut.RLock()
	cache, cached := keyShareInsertCache[key]
	keyShareInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			keyShareAllColumns,
			keyShareColumnsWithDefault,
			keyShareColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(keyShareType, keyShareMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(keyShareType, keyShareMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"key_share\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"key_share\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to insert into key_share")
	}

	if !cached {
		keyShareInsertCacheMut.Lock()
		keyShareInsertCache[key] = cache
		keyShareInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the KeyShare.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *KeyShare) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	keyShareUpdateCacheMut.RLock()
	cache, cached := keyShareUpdateCache[key]
	keyShareUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			keyShareAllColumns,
			keySharePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("sqlboiler: unable to update key_share, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"key_share\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, keySharePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(keyShareType, keyShareMapping, append(wl, keySharePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update key_share row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by update for key_share")
	}

	if !cached {
		keyShareUpdateCacheMut.Lock()
		keyShareUpdateCache[key] = cache
		keyShareUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q keyShareQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update all for key_share")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to retrieve rows affected for key_share")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o KeyShareSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("sqlboiler: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), keySharePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"key_share\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, keySharePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update all in keyShare slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to retrieve rows affected all in update all keyShare")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *KeyShare) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboiler: no key_share provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(keyShareColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	keyShareUpsertCacheMut.RLock()
	cache, cached := keyShareUpsertCache[key]
	keyShareUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			keyShareAllColumns,
			keyShareColumnsWithDefault,
			keyShareColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			keyShareAllColumns,
			keySharePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("sqlboiler: unable to upsert key_share, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(keySharePrimaryKeyColumns))
			copy(conflict, keySharePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"key_share\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(keyShareType, keyShareMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(keyShareType, keyShareMapping, ret)
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to upsert key_share")
	}

	if !cached {
		keyShareUpsertCacheMut.Lock()
		keyShareUpsertCache[key] = cache
		keyShareUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single KeyShare record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *KeyShare) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sqlboiler: no KeyShare provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), keySharePrimaryKeyMapping)
	sql := "DELETE FROM \"key_share\" WHERE \"invitation_hash\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete from key_share")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by delete for key_share")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q keyShareQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sqlboiler: no keyShareQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete all from key_share")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by deleteall for key_share")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o KeyShareSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), keySharePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"key_share\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, keySharePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete all from keyShare slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by deleteall for key_share")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *KeyShare) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindKeyShare(ctx, exec, o.InvitationHash)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *KeyShareSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := KeyShareSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), keySharePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"key_share\".* FROM \"key_share\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, keySharePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to reload all in KeyShareSlice")
	}

	*o = slice

	return nil
}

// KeyShareExists checks if the KeyShare row exists.
func KeyShareExists(ctx context.Context, exec boil.ContextExecutor, invitationHash string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"key_share\" where \"invitation_hash\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, invitationHash)
	}
	row := exec.QueryRowContext(ctx, sql, invitationHash)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: unable to check if key_share exists")
	}

	return exists, nil
}