// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/oio-network/deeplx-extend/app/deeplx/internal/data/ent/accesslog"
	"github.com/oio-network/deeplx-extend/app/deeplx/internal/data/ent/predicate"
	"github.com/oio-network/deeplx-extend/app/deeplx/internal/data/ent/user"
)

// AccessLogQuery is the builder for querying AccessLog entities.
type AccessLogQuery struct {
	config
	ctx           *QueryContext
	order         []accesslog.OrderOption
	inters        []Interceptor
	predicates    []predicate.AccessLog
	withOwnerUser *UserQuery
	modifiers     []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccessLogQuery builder.
func (alq *AccessLogQuery) Where(ps ...predicate.AccessLog) *AccessLogQuery {
	alq.predicates = append(alq.predicates, ps...)
	return alq
}

// Limit the number of records to be returned by this query.
func (alq *AccessLogQuery) Limit(limit int) *AccessLogQuery {
	alq.ctx.Limit = &limit
	return alq
}

// Offset to start from.
func (alq *AccessLogQuery) Offset(offset int) *AccessLogQuery {
	alq.ctx.Offset = &offset
	return alq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (alq *AccessLogQuery) Unique(unique bool) *AccessLogQuery {
	alq.ctx.Unique = &unique
	return alq
}

// Order specifies how the records should be ordered.
func (alq *AccessLogQuery) Order(o ...accesslog.OrderOption) *AccessLogQuery {
	alq.order = append(alq.order, o...)
	return alq
}

// QueryOwnerUser chains the current query on the "owner_user" edge.
func (alq *AccessLogQuery) QueryOwnerUser() *UserQuery {
	query := (&UserClient{config: alq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := alq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := alq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accesslog.Table, accesslog.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, accesslog.OwnerUserTable, accesslog.OwnerUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(alq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AccessLog entity from the query.
// Returns a *NotFoundError when no AccessLog was found.
func (alq *AccessLogQuery) First(ctx context.Context) (*AccessLog, error) {
	nodes, err := alq.Limit(1).All(setContextOp(ctx, alq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{accesslog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (alq *AccessLogQuery) FirstX(ctx context.Context) *AccessLog {
	node, err := alq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AccessLog ID from the query.
// Returns a *NotFoundError when no AccessLog ID was found.
func (alq *AccessLogQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = alq.Limit(1).IDs(setContextOp(ctx, alq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{accesslog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (alq *AccessLogQuery) FirstIDX(ctx context.Context) int64 {
	id, err := alq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AccessLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AccessLog entity is found.
// Returns a *NotFoundError when no AccessLog entities are found.
func (alq *AccessLogQuery) Only(ctx context.Context) (*AccessLog, error) {
	nodes, err := alq.Limit(2).All(setContextOp(ctx, alq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{accesslog.Label}
	default:
		return nil, &NotSingularError{accesslog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (alq *AccessLogQuery) OnlyX(ctx context.Context) *AccessLog {
	node, err := alq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AccessLog ID in the query.
// Returns a *NotSingularError when more than one AccessLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (alq *AccessLogQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = alq.Limit(2).IDs(setContextOp(ctx, alq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{accesslog.Label}
	default:
		err = &NotSingularError{accesslog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (alq *AccessLogQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := alq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AccessLogs.
func (alq *AccessLogQuery) All(ctx context.Context) ([]*AccessLog, error) {
	ctx = setContextOp(ctx, alq.ctx, "All")
	if err := alq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AccessLog, *AccessLogQuery]()
	return withInterceptors[[]*AccessLog](ctx, alq, qr, alq.inters)
}

// AllX is like All, but panics if an error occurs.
func (alq *AccessLogQuery) AllX(ctx context.Context) []*AccessLog {
	nodes, err := alq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AccessLog IDs.
func (alq *AccessLogQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if alq.ctx.Unique == nil && alq.path != nil {
		alq.Unique(true)
	}
	ctx = setContextOp(ctx, alq.ctx, "IDs")
	if err = alq.Select(accesslog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (alq *AccessLogQuery) IDsX(ctx context.Context) []int64 {
	ids, err := alq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (alq *AccessLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, alq.ctx, "Count")
	if err := alq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, alq, querierCount[*AccessLogQuery](), alq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (alq *AccessLogQuery) CountX(ctx context.Context) int {
	count, err := alq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (alq *AccessLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, alq.ctx, "Exist")
	switch _, err := alq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (alq *AccessLogQuery) ExistX(ctx context.Context) bool {
	exist, err := alq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccessLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (alq *AccessLogQuery) Clone() *AccessLogQuery {
	if alq == nil {
		return nil
	}
	return &AccessLogQuery{
		config:        alq.config,
		ctx:           alq.ctx.Clone(),
		order:         append([]accesslog.OrderOption{}, alq.order...),
		inters:        append([]Interceptor{}, alq.inters...),
		predicates:    append([]predicate.AccessLog{}, alq.predicates...),
		withOwnerUser: alq.withOwnerUser.Clone(),
		// clone intermediate query.
		sql:  alq.sql.Clone(),
		path: alq.path,
	}
}

// WithOwnerUser tells the query-builder to eager-load the nodes that are connected to
// the "owner_user" edge. The optional arguments are used to configure the query builder of the edge.
func (alq *AccessLogQuery) WithOwnerUser(opts ...func(*UserQuery)) *AccessLogQuery {
	query := (&UserClient{config: alq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	alq.withOwnerUser = query
	return alq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AccessLog.Query().
//		GroupBy(accesslog.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (alq *AccessLogQuery) GroupBy(field string, fields ...string) *AccessLogGroupBy {
	alq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AccessLogGroupBy{build: alq}
	grbuild.flds = &alq.ctx.Fields
	grbuild.label = accesslog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.AccessLog.Query().
//		Select(accesslog.FieldCreatedAt).
//		Scan(ctx, &v)
func (alq *AccessLogQuery) Select(fields ...string) *AccessLogSelect {
	alq.ctx.Fields = append(alq.ctx.Fields, fields...)
	sbuild := &AccessLogSelect{AccessLogQuery: alq}
	sbuild.label = accesslog.Label
	sbuild.flds, sbuild.scan = &alq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AccessLogSelect configured with the given aggregations.
func (alq *AccessLogQuery) Aggregate(fns ...AggregateFunc) *AccessLogSelect {
	return alq.Select().Aggregate(fns...)
}

func (alq *AccessLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range alq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, alq); err != nil {
				return err
			}
		}
	}
	for _, f := range alq.ctx.Fields {
		if !accesslog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if alq.path != nil {
		prev, err := alq.path(ctx)
		if err != nil {
			return err
		}
		alq.sql = prev
	}
	return nil
}

func (alq *AccessLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AccessLog, error) {
	var (
		nodes       = []*AccessLog{}
		_spec       = alq.querySpec()
		loadedTypes = [1]bool{
			alq.withOwnerUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AccessLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AccessLog{config: alq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(alq.modifiers) > 0 {
		_spec.Modifiers = alq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, alq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := alq.withOwnerUser; query != nil {
		if err := alq.loadOwnerUser(ctx, query, nodes, nil,
			func(n *AccessLog, e *User) { n.Edges.OwnerUser = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (alq *AccessLogQuery) loadOwnerUser(ctx context.Context, query *UserQuery, nodes []*AccessLog, init func(*AccessLog), assign func(*AccessLog, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*AccessLog)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (alq *AccessLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := alq.querySpec()
	if len(alq.modifiers) > 0 {
		_spec.Modifiers = alq.modifiers
	}
	_spec.Node.Columns = alq.ctx.Fields
	if len(alq.ctx.Fields) > 0 {
		_spec.Unique = alq.ctx.Unique != nil && *alq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, alq.driver, _spec)
}

func (alq *AccessLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(accesslog.Table, accesslog.Columns, sqlgraph.NewFieldSpec(accesslog.FieldID, field.TypeInt64))
	_spec.From = alq.sql
	if unique := alq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if alq.path != nil {
		_spec.Unique = true
	}
	if fields := alq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accesslog.FieldID)
		for i := range fields {
			if fields[i] != accesslog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if alq.withOwnerUser != nil {
			_spec.Node.AddColumnOnce(accesslog.FieldUserID)
		}
	}
	if ps := alq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := alq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := alq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := alq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (alq *AccessLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(alq.driver.Dialect())
	t1 := builder.Table(accesslog.Table)
	columns := alq.ctx.Fields
	if len(columns) == 0 {
		columns = accesslog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if alq.sql != nil {
		selector = alq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if alq.ctx.Unique != nil && *alq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range alq.modifiers {
		m(selector)
	}
	for _, p := range alq.predicates {
		p(selector)
	}
	for _, p := range alq.order {
		p(selector)
	}
	if offset := alq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := alq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (alq *AccessLogQuery) Modify(modifiers ...func(s *sql.Selector)) *AccessLogSelect {
	alq.modifiers = append(alq.modifiers, modifiers...)
	return alq.Select()
}

// AccessLogGroupBy is the group-by builder for AccessLog entities.
type AccessLogGroupBy struct {
	selector
	build *AccessLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (algb *AccessLogGroupBy) Aggregate(fns ...AggregateFunc) *AccessLogGroupBy {
	algb.fns = append(algb.fns, fns...)
	return algb
}

// Scan applies the selector query and scans the result into the given value.
func (algb *AccessLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, algb.build.ctx, "GroupBy")
	if err := algb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccessLogQuery, *AccessLogGroupBy](ctx, algb.build, algb, algb.build.inters, v)
}

func (algb *AccessLogGroupBy) sqlScan(ctx context.Context, root *AccessLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(algb.fns))
	for _, fn := range algb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*algb.flds)+len(algb.fns))
		for _, f := range *algb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*algb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := algb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AccessLogSelect is the builder for selecting fields of AccessLog entities.
type AccessLogSelect struct {
	*AccessLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (als *AccessLogSelect) Aggregate(fns ...AggregateFunc) *AccessLogSelect {
	als.fns = append(als.fns, fns...)
	return als
}

// Scan applies the selector query and scans the result into the given value.
func (als *AccessLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, als.ctx, "Select")
	if err := als.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccessLogQuery, *AccessLogSelect](ctx, als.AccessLogQuery, als, als.inters, v)
}

func (als *AccessLogSelect) sqlScan(ctx context.Context, root *AccessLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(als.fns))
	for _, fn := range als.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*als.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := als.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (als *AccessLogSelect) Modify(modifiers ...func(s *sql.Selector)) *AccessLogSelect {
	als.modifiers = append(als.modifiers, modifiers...)
	return als
}
