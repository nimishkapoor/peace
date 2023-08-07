// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kapt/kapt/gql/ent/predicate"
	"kapt/kapt/gql/ent/ticketmodel"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TicketModelQuery is the builder for querying TicketModel entities.
type TicketModelQuery struct {
	config
	ctx        *QueryContext
	order      []ticketmodel.OrderOption
	inters     []Interceptor
	predicates []predicate.TicketModel
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TicketModelQuery builder.
func (tmq *TicketModelQuery) Where(ps ...predicate.TicketModel) *TicketModelQuery {
	tmq.predicates = append(tmq.predicates, ps...)
	return tmq
}

// Limit the number of records to be returned by this query.
func (tmq *TicketModelQuery) Limit(limit int) *TicketModelQuery {
	tmq.ctx.Limit = &limit
	return tmq
}

// Offset to start from.
func (tmq *TicketModelQuery) Offset(offset int) *TicketModelQuery {
	tmq.ctx.Offset = &offset
	return tmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tmq *TicketModelQuery) Unique(unique bool) *TicketModelQuery {
	tmq.ctx.Unique = &unique
	return tmq
}

// Order specifies how the records should be ordered.
func (tmq *TicketModelQuery) Order(o ...ticketmodel.OrderOption) *TicketModelQuery {
	tmq.order = append(tmq.order, o...)
	return tmq
}

// First returns the first TicketModel entity from the query.
// Returns a *NotFoundError when no TicketModel was found.
func (tmq *TicketModelQuery) First(ctx context.Context) (*TicketModel, error) {
	nodes, err := tmq.Limit(1).All(setContextOp(ctx, tmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ticketmodel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tmq *TicketModelQuery) FirstX(ctx context.Context) *TicketModel {
	node, err := tmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TicketModel ID from the query.
// Returns a *NotFoundError when no TicketModel ID was found.
func (tmq *TicketModelQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tmq.Limit(1).IDs(setContextOp(ctx, tmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ticketmodel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tmq *TicketModelQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TicketModel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TicketModel entity is found.
// Returns a *NotFoundError when no TicketModel entities are found.
func (tmq *TicketModelQuery) Only(ctx context.Context) (*TicketModel, error) {
	nodes, err := tmq.Limit(2).All(setContextOp(ctx, tmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ticketmodel.Label}
	default:
		return nil, &NotSingularError{ticketmodel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tmq *TicketModelQuery) OnlyX(ctx context.Context) *TicketModel {
	node, err := tmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TicketModel ID in the query.
// Returns a *NotSingularError when more than one TicketModel ID is found.
// Returns a *NotFoundError when no entities are found.
func (tmq *TicketModelQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tmq.Limit(2).IDs(setContextOp(ctx, tmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ticketmodel.Label}
	default:
		err = &NotSingularError{ticketmodel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tmq *TicketModelQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TicketModels.
func (tmq *TicketModelQuery) All(ctx context.Context) ([]*TicketModel, error) {
	ctx = setContextOp(ctx, tmq.ctx, "All")
	if err := tmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TicketModel, *TicketModelQuery]()
	return withInterceptors[[]*TicketModel](ctx, tmq, qr, tmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tmq *TicketModelQuery) AllX(ctx context.Context) []*TicketModel {
	nodes, err := tmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TicketModel IDs.
func (tmq *TicketModelQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tmq.ctx.Unique == nil && tmq.path != nil {
		tmq.Unique(true)
	}
	ctx = setContextOp(ctx, tmq.ctx, "IDs")
	if err = tmq.Select(ticketmodel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tmq *TicketModelQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tmq *TicketModelQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tmq.ctx, "Count")
	if err := tmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tmq, querierCount[*TicketModelQuery](), tmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tmq *TicketModelQuery) CountX(ctx context.Context) int {
	count, err := tmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tmq *TicketModelQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tmq.ctx, "Exist")
	switch _, err := tmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tmq *TicketModelQuery) ExistX(ctx context.Context) bool {
	exist, err := tmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TicketModelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tmq *TicketModelQuery) Clone() *TicketModelQuery {
	if tmq == nil {
		return nil
	}
	return &TicketModelQuery{
		config:     tmq.config,
		ctx:        tmq.ctx.Clone(),
		order:      append([]ticketmodel.OrderOption{}, tmq.order...),
		inters:     append([]Interceptor{}, tmq.inters...),
		predicates: append([]predicate.TicketModel{}, tmq.predicates...),
		// clone intermediate query.
		sql:  tmq.sql.Clone(),
		path: tmq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TicketModel.Query().
//		GroupBy(ticketmodel.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tmq *TicketModelQuery) GroupBy(field string, fields ...string) *TicketModelGroupBy {
	tmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TicketModelGroupBy{build: tmq}
	grbuild.flds = &tmq.ctx.Fields
	grbuild.label = ticketmodel.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//	}
//
//	client.TicketModel.Query().
//		Select(ticketmodel.FieldUserID).
//		Scan(ctx, &v)
func (tmq *TicketModelQuery) Select(fields ...string) *TicketModelSelect {
	tmq.ctx.Fields = append(tmq.ctx.Fields, fields...)
	sbuild := &TicketModelSelect{TicketModelQuery: tmq}
	sbuild.label = ticketmodel.Label
	sbuild.flds, sbuild.scan = &tmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TicketModelSelect configured with the given aggregations.
func (tmq *TicketModelQuery) Aggregate(fns ...AggregateFunc) *TicketModelSelect {
	return tmq.Select().Aggregate(fns...)
}

func (tmq *TicketModelQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tmq); err != nil {
				return err
			}
		}
	}
	for _, f := range tmq.ctx.Fields {
		if !ticketmodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tmq.path != nil {
		prev, err := tmq.path(ctx)
		if err != nil {
			return err
		}
		tmq.sql = prev
	}
	return nil
}

func (tmq *TicketModelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TicketModel, error) {
	var (
		nodes = []*TicketModel{}
		_spec = tmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TicketModel).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TicketModel{config: tmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tmq *TicketModelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tmq.querySpec()
	_spec.Node.Columns = tmq.ctx.Fields
	if len(tmq.ctx.Fields) > 0 {
		_spec.Unique = tmq.ctx.Unique != nil && *tmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tmq.driver, _spec)
}

func (tmq *TicketModelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(ticketmodel.Table, ticketmodel.Columns, sqlgraph.NewFieldSpec(ticketmodel.FieldID, field.TypeUUID))
	_spec.From = tmq.sql
	if unique := tmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tmq.path != nil {
		_spec.Unique = true
	}
	if fields := tmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ticketmodel.FieldID)
		for i := range fields {
			if fields[i] != ticketmodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tmq *TicketModelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tmq.driver.Dialect())
	t1 := builder.Table(ticketmodel.Table)
	columns := tmq.ctx.Fields
	if len(columns) == 0 {
		columns = ticketmodel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tmq.sql != nil {
		selector = tmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tmq.ctx.Unique != nil && *tmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tmq.predicates {
		p(selector)
	}
	for _, p := range tmq.order {
		p(selector)
	}
	if offset := tmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TicketModelGroupBy is the group-by builder for TicketModel entities.
type TicketModelGroupBy struct {
	selector
	build *TicketModelQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tmgb *TicketModelGroupBy) Aggregate(fns ...AggregateFunc) *TicketModelGroupBy {
	tmgb.fns = append(tmgb.fns, fns...)
	return tmgb
}

// Scan applies the selector query and scans the result into the given value.
func (tmgb *TicketModelGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tmgb.build.ctx, "GroupBy")
	if err := tmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TicketModelQuery, *TicketModelGroupBy](ctx, tmgb.build, tmgb, tmgb.build.inters, v)
}

func (tmgb *TicketModelGroupBy) sqlScan(ctx context.Context, root *TicketModelQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tmgb.fns))
	for _, fn := range tmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tmgb.flds)+len(tmgb.fns))
		for _, f := range *tmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TicketModelSelect is the builder for selecting fields of TicketModel entities.
type TicketModelSelect struct {
	*TicketModelQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tms *TicketModelSelect) Aggregate(fns ...AggregateFunc) *TicketModelSelect {
	tms.fns = append(tms.fns, fns...)
	return tms
}

// Scan applies the selector query and scans the result into the given value.
func (tms *TicketModelSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tms.ctx, "Select")
	if err := tms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TicketModelQuery, *TicketModelSelect](ctx, tms.TicketModelQuery, tms, tms.inters, v)
}

func (tms *TicketModelSelect) sqlScan(ctx context.Context, root *TicketModelQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tms.fns))
	for _, fn := range tms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
