// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AuthAcknowledgmentLog is an object representing the database table.
type AuthAcknowledgmentLog struct { // primary key
	AuthAcknowledgmentLogID int64 `boil:"auth_acknowledgment_log_id" json:"auth_acknowledgment_log_id" toml:"auth_acknowledgment_log_id" yaml:"auth_acknowledgment_log_id"`
	// gateway hash
	AuthAcknowledgmentLogGatewayHash string `boil:"auth_acknowledgment_log_gateway_hash" json:"auth_acknowledgment_log_gateway_hash" toml:"auth_acknowledgment_log_gateway_hash" yaml:"auth_acknowledgment_log_gateway_hash"`
	// raw payload string in gateway ack request
	AuthAcknowledgmentLogRawPayload string `boil:"auth_acknowledgment_log_raw_payload" json:"auth_acknowledgment_log_raw_payload" toml:"auth_acknowledgment_log_raw_payload" yaml:"auth_acknowledgment_log_raw_payload"`

	R *authAcknowledgmentLogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authAcknowledgmentLogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AuthAcknowledgmentLogColumns = struct {
	AuthAcknowledgmentLogID          string
	AuthAcknowledgmentLogGatewayHash string
	AuthAcknowledgmentLogRawPayload  string
}{
	AuthAcknowledgmentLogID:          "auth_acknowledgment_log_id",
	AuthAcknowledgmentLogGatewayHash: "auth_acknowledgment_log_gateway_hash",
	AuthAcknowledgmentLogRawPayload:  "auth_acknowledgment_log_raw_payload",
}

var AuthAcknowledgmentLogTableColumns = struct {
	AuthAcknowledgmentLogID          string
	AuthAcknowledgmentLogGatewayHash string
	AuthAcknowledgmentLogRawPayload  string
}{
	AuthAcknowledgmentLogID:          "auth_acknowledgment_log.auth_acknowledgment_log_id",
	AuthAcknowledgmentLogGatewayHash: "auth_acknowledgment_log.auth_acknowledgment_log_gateway_hash",
	AuthAcknowledgmentLogRawPayload:  "auth_acknowledgment_log.auth_acknowledgment_log_raw_payload",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var AuthAcknowledgmentLogWhere = struct {
	AuthAcknowledgmentLogID          whereHelperint64
	AuthAcknowledgmentLogGatewayHash whereHelperstring
	AuthAcknowledgmentLogRawPayload  whereHelperstring
}{
	AuthAcknowledgmentLogID:          whereHelperint64{field: "\"auth_acknowledgment_log\".\"auth_acknowledgment_log_id\""},
	AuthAcknowledgmentLogGatewayHash: whereHelperstring{field: "\"auth_acknowledgment_log\".\"auth_acknowledgment_log_gateway_hash\""},
	AuthAcknowledgmentLogRawPayload:  whereHelperstring{field: "\"auth_acknowledgment_log\".\"auth_acknowledgment_log_raw_payload\""},
}

// AuthAcknowledgmentLogRels is where relationship names are stored.
var AuthAcknowledgmentLogRels = struct {
}{}

// authAcknowledgmentLogR is where relationships are stored.
type authAcknowledgmentLogR struct {
}

// NewStruct creates a new relationship struct
func (*authAcknowledgmentLogR) NewStruct() *authAcknowledgmentLogR {
	return &authAcknowledgmentLogR{}
}

// authAcknowledgmentLogL is where Load methods for each relationship are stored.
type authAcknowledgmentLogL struct{}

var (
	authAcknowledgmentLogAllColumns            = []string{"auth_acknowledgment_log_id", "auth_acknowledgment_log_gateway_hash", "auth_acknowledgment_log_raw_payload"}
	authAcknowledgmentLogColumnsWithoutDefault = []string{"auth_acknowledgment_log_id", "auth_acknowledgment_log_gateway_hash", "auth_acknowledgment_log_raw_payload"}
	authAcknowledgmentLogColumnsWithDefault    = []string{}
	authAcknowledgmentLogPrimaryKeyColumns     = []string{"auth_acknowledgment_log_id"}
	authAcknowledgmentLogGeneratedColumns      = []string{}
)

type (
	// AuthAcknowledgmentLogSlice is an alias for a slice of pointers to AuthAcknowledgmentLog.
	// This should almost always be used instead of []AuthAcknowledgmentLog.
	AuthAcknowledgmentLogSlice []*AuthAcknowledgmentLog
	// AuthAcknowledgmentLogHook is the signature for custom AuthAcknowledgmentLog hook methods
	AuthAcknowledgmentLogHook func(context.Context, boil.ContextExecutor, *AuthAcknowledgmentLog) error

	authAcknowledgmentLogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authAcknowledgmentLogType                 = reflect.TypeOf(&AuthAcknowledgmentLog{})
	authAcknowledgmentLogMapping              = queries.MakeStructMapping(authAcknowledgmentLogType)
	authAcknowledgmentLogPrimaryKeyMapping, _ = queries.BindMapping(authAcknowledgmentLogType, authAcknowledgmentLogMapping, authAcknowledgmentLogPrimaryKeyColumns)
	authAcknowledgmentLogInsertCacheMut       sync.RWMutex
	authAcknowledgmentLogInsertCache          = make(map[string]insertCache)
	authAcknowledgmentLogUpdateCacheMut       sync.RWMutex
	authAcknowledgmentLogUpdateCache          = make(map[string]updateCache)
	authAcknowledgmentLogUpsertCacheMut       sync.RWMutex
	authAcknowledgmentLogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var authAcknowledgmentLogAfterSelectHooks []AuthAcknowledgmentLogHook

var authAcknowledgmentLogBeforeInsertHooks []AuthAcknowledgmentLogHook
var authAcknowledgmentLogAfterInsertHooks []AuthAcknowledgmentLogHook

var authAcknowledgmentLogBeforeUpdateHooks []AuthAcknowledgmentLogHook
var authAcknowledgmentLogAfterUpdateHooks []AuthAcknowledgmentLogHook

var authAcknowledgmentLogBeforeDeleteHooks []AuthAcknowledgmentLogHook
var authAcknowledgmentLogAfterDeleteHooks []AuthAcknowledgmentLogHook

var authAcknowledgmentLogBeforeUpsertHooks []AuthAcknowledgmentLogHook
var authAcknowledgmentLogAfterUpsertHooks []AuthAcknowledgmentLogHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthAcknowledgmentLog) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authAcknowledgmentLogAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthAcknowledgmentLog) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authAcknowledgmentLogBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthAcknowledgmentLog) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authAcknowledgmentLogAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthAcknowledgmentLog) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authAcknowledgmentLogBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthAcknowledgmentLog) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authAcknowledgmentLogAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthAcknowledgmentLog) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authAcknowledgmentLogBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthAcknowledgmentLog) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authAcknowledgmentLogAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthAcknowledgmentLog) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authAcknowledgmentLogBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthAcknowledgmentLog) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authAcknowledgmentLogAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthAcknowledgmentLogHook registers your hook function for all future operations.
func AddAuthAcknowledgmentLogHook(hookPoint boil.HookPoint, authAcknowledgmentLogHook AuthAcknowledgmentLogHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		authAcknowledgmentLogAfterSelectHooks = append(authAcknowledgmentLogAfterSelectHooks, authAcknowledgmentLogHook)
	case boil.BeforeInsertHook:
		authAcknowledgmentLogBeforeInsertHooks = append(authAcknowledgmentLogBeforeInsertHooks, authAcknowledgmentLogHook)
	case boil.AfterInsertHook:
		authAcknowledgmentLogAfterInsertHooks = append(authAcknowledgmentLogAfterInsertHooks, authAcknowledgmentLogHook)
	case boil.BeforeUpdateHook:
		authAcknowledgmentLogBeforeUpdateHooks = append(authAcknowledgmentLogBeforeUpdateHooks, authAcknowledgmentLogHook)
	case boil.AfterUpdateHook:
		authAcknowledgmentLogAfterUpdateHooks = append(authAcknowledgmentLogAfterUpdateHooks, authAcknowledgmentLogHook)
	case boil.BeforeDeleteHook:
		authAcknowledgmentLogBeforeDeleteHooks = append(authAcknowledgmentLogBeforeDeleteHooks, authAcknowledgmentLogHook)
	case boil.AfterDeleteHook:
		authAcknowledgmentLogAfterDeleteHooks = append(authAcknowledgmentLogAfterDeleteHooks, authAcknowledgmentLogHook)
	case boil.BeforeUpsertHook:
		authAcknowledgmentLogBeforeUpsertHooks = append(authAcknowledgmentLogBeforeUpsertHooks, authAcknowledgmentLogHook)
	case boil.AfterUpsertHook:
		authAcknowledgmentLogAfterUpsertHooks = append(authAcknowledgmentLogAfterUpsertHooks, authAcknowledgmentLogHook)
	}
}

// One returns a single authAcknowledgmentLog record from the query.
func (q authAcknowledgmentLogQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AuthAcknowledgmentLog, error) {
	o := &AuthAcknowledgmentLog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for auth_acknowledgment_log")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AuthAcknowledgmentLog records from the query.
func (q authAcknowledgmentLogQuery) All(ctx context.Context, exec boil.ContextExecutor) (AuthAcknowledgmentLogSlice, error) {
	var o []*AuthAcknowledgmentLog

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AuthAcknowledgmentLog slice")
	}

	if len(authAcknowledgmentLogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AuthAcknowledgmentLog records in the query.
func (q authAcknowledgmentLogQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count auth_acknowledgment_log rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q authAcknowledgmentLogQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if auth_acknowledgment_log exists")
	}

	return count > 0, nil
}

// AuthAcknowledgmentLogs retrieves all the records using an executor.
func AuthAcknowledgmentLogs(mods ...qm.QueryMod) authAcknowledgmentLogQuery {
	mods = append(mods, qm.From("\"auth_acknowledgment_log\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"auth_acknowledgment_log\".*"})
	}

	return authAcknowledgmentLogQuery{q}
}

// FindAuthAcknowledgmentLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthAcknowledgmentLog(ctx context.Context, exec boil.ContextExecutor, authAcknowledgmentLogID int64, selectCols ...string) (*AuthAcknowledgmentLog, error) {
	authAcknowledgmentLogObj := &AuthAcknowledgmentLog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auth_acknowledgment_log\" where \"auth_acknowledgment_log_id\"=$1", sel,
	)

	q := queries.Raw(query, authAcknowledgmentLogID)

	err := q.Bind(ctx, exec, authAcknowledgmentLogObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from auth_acknowledgment_log")
	}

	if err = authAcknowledgmentLogObj.doAfterSelectHooks(ctx, exec); err != nil {
		return authAcknowledgmentLogObj, err
	}

	return authAcknowledgmentLogObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AuthAcknowledgmentLog) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no auth_acknowledgment_log provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authAcknowledgmentLogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	authAcknowledgmentLogInsertCacheMut.RLock()
	cache, cached := authAcknowledgmentLogInsertCache[key]
	authAcknowledgmentLogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			authAcknowledgmentLogAllColumns,
			authAcknowledgmentLogColumnsWithDefault,
			authAcknowledgmentLogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(authAcknowledgmentLogType, authAcknowledgmentLogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authAcknowledgmentLogType, authAcknowledgmentLogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"auth_acknowledgment_log\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"auth_acknowledgment_log\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into auth_acknowledgment_log")
	}

	if !cached {
		authAcknowledgmentLogInsertCacheMut.Lock()
		authAcknowledgmentLogInsertCache[key] = cache
		authAcknowledgmentLogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AuthAcknowledgmentLog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AuthAcknowledgmentLog) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	authAcknowledgmentLogUpdateCacheMut.RLock()
	cache, cached := authAcknowledgmentLogUpdateCache[key]
	authAcknowledgmentLogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			authAcknowledgmentLogAllColumns,
			authAcknowledgmentLogPrimaryKeyColumns,
		)
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update auth_acknowledgment_log, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auth_acknowledgment_log\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authAcknowledgmentLogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authAcknowledgmentLogType, authAcknowledgmentLogMapping, append(wl, authAcknowledgmentLogPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update auth_acknowledgment_log row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for auth_acknowledgment_log")
	}

	if !cached {
		authAcknowledgmentLogUpdateCacheMut.Lock()
		authAcknowledgmentLogUpdateCache[key] = cache
		authAcknowledgmentLogUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q authAcknowledgmentLogQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for auth_acknowledgment_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for auth_acknowledgment_log")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthAcknowledgmentLogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authAcknowledgmentLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"auth_acknowledgment_log\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, authAcknowledgmentLogPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in authAcknowledgmentLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all authAcknowledgmentLog")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AuthAcknowledgmentLog) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no auth_acknowledgment_log provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authAcknowledgmentLogColumnsWithDefault, o)

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

	authAcknowledgmentLogUpsertCacheMut.RLock()
	cache, cached := authAcknowledgmentLogUpsertCache[key]
	authAcknowledgmentLogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			authAcknowledgmentLogAllColumns,
			authAcknowledgmentLogColumnsWithDefault,
			authAcknowledgmentLogColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			authAcknowledgmentLogAllColumns,
			authAcknowledgmentLogPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert auth_acknowledgment_log, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authAcknowledgmentLogPrimaryKeyColumns))
			copy(conflict, authAcknowledgmentLogPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"auth_acknowledgment_log\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(authAcknowledgmentLogType, authAcknowledgmentLogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authAcknowledgmentLogType, authAcknowledgmentLogMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert auth_acknowledgment_log")
	}

	if !cached {
		authAcknowledgmentLogUpsertCacheMut.Lock()
		authAcknowledgmentLogUpsertCache[key] = cache
		authAcknowledgmentLogUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AuthAcknowledgmentLog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthAcknowledgmentLog) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AuthAcknowledgmentLog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authAcknowledgmentLogPrimaryKeyMapping)
	sql := "DELETE FROM \"auth_acknowledgment_log\" WHERE \"auth_acknowledgment_log_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from auth_acknowledgment_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for auth_acknowledgment_log")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q authAcknowledgmentLogQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no authAcknowledgmentLogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from auth_acknowledgment_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for auth_acknowledgment_log")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthAcknowledgmentLogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(authAcknowledgmentLogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authAcknowledgmentLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"auth_acknowledgment_log\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, authAcknowledgmentLogPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from authAcknowledgmentLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for auth_acknowledgment_log")
	}

	if len(authAcknowledgmentLogAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AuthAcknowledgmentLog) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAuthAcknowledgmentLog(ctx, exec, o.AuthAcknowledgmentLogID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthAcknowledgmentLogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AuthAcknowledgmentLogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authAcknowledgmentLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"auth_acknowledgment_log\".* FROM \"auth_acknowledgment_log\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, authAcknowledgmentLogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuthAcknowledgmentLogSlice")
	}

	*o = slice

	return nil
}

// AuthAcknowledgmentLogExists checks if the AuthAcknowledgmentLog row exists.
func AuthAcknowledgmentLogExists(ctx context.Context, exec boil.ContextExecutor, authAcknowledgmentLogID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"auth_acknowledgment_log\" where \"auth_acknowledgment_log_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, authAcknowledgmentLogID)
	}
	row := exec.QueryRowContext(ctx, sql, authAcknowledgmentLogID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if auth_acknowledgment_log exists")
	}

	return exists, nil
}

// Exists checks if the AuthAcknowledgmentLog row exists.
func (o *AuthAcknowledgmentLog) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AuthAcknowledgmentLogExists(ctx, exec, o.AuthAcknowledgmentLogID)
}

func chunkAuthAcknowledgmentLogs(o []*AuthAcknowledgmentLog, chuckSize int) [][]*AuthAcknowledgmentLog {
	chunks := make([][]*AuthAcknowledgmentLog, 0, (len(o)+chuckSize-1)/chuckSize)

	for chuckSize < len(o) {
		o, chunks = o[chuckSize:], append(chunks, o[0:chuckSize:chuckSize])
	}
	chunks = append(chunks, o)

	return chunks
}

// BulkInsertAuthAcknowledgmentLog Ref) https://stackoverflow.com/a/25192138/8979550
func BulkInsertAuthAcknowledgmentLog(unsavedRows []*AuthAcknowledgmentLog, exec boil.ContextExecutor) error {
	if len(unsavedRows) == 0 {
		return nil
	}

	for _, authAcknowledgmentLogs := range chunkAuthAcknowledgmentLogs(unsavedRows, 32767) {
		valueStrings := make([]string, 0, len(unsavedRows))
		valueArgs := make([]interface{}, 0, len(unsavedRows)*2)

		for _, authAcknowledgmentLog := range authAcknowledgmentLogs {
			valueStrings = append(valueStrings, "(?,?)")

			valueArgs = append(valueArgs, authAcknowledgmentLog.AuthAcknowledgmentLogGatewayHash)
			valueArgs = append(valueArgs, authAcknowledgmentLog.AuthAcknowledgmentLogRawPayload)
		}
		stmt := fmt.Sprintf(
			"INSERT INTO "+
				TableNames.AuthAcknowledgmentLog+
				"("+
				AuthAcknowledgmentLogColumns.AuthAcknowledgmentLogGatewayHash+", "+
				AuthAcknowledgmentLogColumns.AuthAcknowledgmentLogRawPayload+
				") VALUES %s",
			strings.Join(valueStrings, ","))
		if _, err := exec.Exec(stmt, valueArgs...); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}