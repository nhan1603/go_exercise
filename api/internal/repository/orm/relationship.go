// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Relationship is an object representing the database table.
type Relationship struct {
	ID            int `boil:"id" json:"id" toml:"id" yaml:"id"`
	FirstEmailID  int `boil:"first_email_id" json:"first_email_id" toml:"first_email_id" yaml:"first_email_id"`
	SecondEmailID int `boil:"second_email_id" json:"second_email_id" toml:"second_email_id" yaml:"second_email_id"`
	Status        int `boil:"status" json:"status" toml:"status" yaml:"status"`

	R *relationshipR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L relationshipL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RelationshipColumns = struct {
	ID            string
	FirstEmailID  string
	SecondEmailID string
	Status        string
}{
	ID:            "id",
	FirstEmailID:  "first_email_id",
	SecondEmailID: "second_email_id",
	Status:        "status",
}

var RelationshipTableColumns = struct {
	ID            string
	FirstEmailID  string
	SecondEmailID string
	Status        string
}{
	ID:            "relationship.id",
	FirstEmailID:  "relationship.first_email_id",
	SecondEmailID: "relationship.second_email_id",
	Status:        "relationship.status",
}

// Generated where

var RelationshipWhere = struct {
	ID            whereHelperint
	FirstEmailID  whereHelperint
	SecondEmailID whereHelperint
	Status        whereHelperint
}{
	ID:            whereHelperint{field: "\"relationship\".\"id\""},
	FirstEmailID:  whereHelperint{field: "\"relationship\".\"first_email_id\""},
	SecondEmailID: whereHelperint{field: "\"relationship\".\"second_email_id\""},
	Status:        whereHelperint{field: "\"relationship\".\"status\""},
}

// RelationshipRels is where relationship names are stored.
var RelationshipRels = struct {
}{}

// relationshipR is where relationships are stored.
type relationshipR struct {
}

// NewStruct creates a new relationship struct
func (*relationshipR) NewStruct() *relationshipR {
	return &relationshipR{}
}

// relationshipL is where Load methods for each relationship are stored.
type relationshipL struct{}

var (
	relationshipAllColumns            = []string{"id", "first_email_id", "second_email_id", "status"}
	relationshipColumnsWithoutDefault = []string{"first_email_id", "second_email_id"}
	relationshipColumnsWithDefault    = []string{"id", "status"}
	relationshipPrimaryKeyColumns     = []string{"id"}
	relationshipGeneratedColumns      = []string{}
)

type (
	// RelationshipSlice is an alias for a slice of pointers to Relationship.
	// This should almost always be used instead of []Relationship.
	RelationshipSlice []*Relationship

	relationshipQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	relationshipType                 = reflect.TypeOf(&Relationship{})
	relationshipMapping              = queries.MakeStructMapping(relationshipType)
	relationshipPrimaryKeyMapping, _ = queries.BindMapping(relationshipType, relationshipMapping, relationshipPrimaryKeyColumns)
	relationshipInsertCacheMut       sync.RWMutex
	relationshipInsertCache          = make(map[string]insertCache)
	relationshipUpdateCacheMut       sync.RWMutex
	relationshipUpdateCache          = make(map[string]updateCache)
	relationshipUpsertCacheMut       sync.RWMutex
	relationshipUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single relationship record from the query.
func (q relationshipQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Relationship, error) {
	o := &Relationship{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for relationship")
	}

	return o, nil
}

// All returns all Relationship records from the query.
func (q relationshipQuery) All(ctx context.Context, exec boil.ContextExecutor) (RelationshipSlice, error) {
	var o []*Relationship

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to Relationship slice")
	}

	return o, nil
}

// Count returns the count of all Relationship records in the query.
func (q relationshipQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count relationship rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q relationshipQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if relationship exists")
	}

	return count > 0, nil
}

// Relationships retrieves all the records using an executor.
func Relationships(mods ...qm.QueryMod) relationshipQuery {
	mods = append(mods, qm.From("\"relationship\""))
	return relationshipQuery{NewQuery(mods...)}
}

// FindRelationship retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRelationship(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Relationship, error) {
	relationshipObj := &Relationship{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"relationship\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, relationshipObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from relationship")
	}

	return relationshipObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Relationship) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no relationship provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(relationshipColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	relationshipInsertCacheMut.RLock()
	cache, cached := relationshipInsertCache[key]
	relationshipInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			relationshipAllColumns,
			relationshipColumnsWithDefault,
			relationshipColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(relationshipType, relationshipMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(relationshipType, relationshipMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"relationship\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"relationship\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "orm: unable to insert into relationship")
	}

	if !cached {
		relationshipInsertCacheMut.Lock()
		relationshipInsertCache[key] = cache
		relationshipInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Relationship.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Relationship) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	relationshipUpdateCacheMut.RLock()
	cache, cached := relationshipUpdateCache[key]
	relationshipUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			relationshipAllColumns,
			relationshipPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update relationship, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, relationshipPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(relationshipType, relationshipMapping, append(wl, relationshipPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "orm: unable to update relationship row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for relationship")
	}

	if !cached {
		relationshipUpdateCacheMut.Lock()
		relationshipUpdateCache[key] = cache
		relationshipUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q relationshipQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for relationship")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for relationship")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RelationshipSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("orm: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), relationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"relationship\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, relationshipPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in relationship slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all relationship")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Relationship) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no relationship provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(relationshipColumnsWithDefault, o)

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

	relationshipUpsertCacheMut.RLock()
	cache, cached := relationshipUpsertCache[key]
	relationshipUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			relationshipAllColumns,
			relationshipColumnsWithDefault,
			relationshipColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			relationshipAllColumns,
			relationshipPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("orm: unable to upsert relationship, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(relationshipPrimaryKeyColumns))
			copy(conflict, relationshipPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"relationship\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(relationshipType, relationshipMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(relationshipType, relationshipMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
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
		return errors.Wrap(err, "orm: unable to upsert relationship")
	}

	if !cached {
		relationshipUpsertCacheMut.Lock()
		relationshipUpsertCache[key] = cache
		relationshipUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Relationship record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Relationship) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no Relationship provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), relationshipPrimaryKeyMapping)
	sql := "DELETE FROM \"relationship\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from relationship")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for relationship")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q relationshipQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no relationshipQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from relationship")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for relationship")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RelationshipSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), relationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"relationship\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, relationshipPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from relationship slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for relationship")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Relationship) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRelationship(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RelationshipSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RelationshipSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), relationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"relationship\".* FROM \"relationship\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, relationshipPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in RelationshipSlice")
	}

	*o = slice

	return nil
}

// RelationshipExists checks if the Relationship row exists.
func RelationshipExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"relationship\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if relationship exists")
	}

	return exists, nil
}