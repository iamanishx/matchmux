// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"ipc/ent/otp"
	"ipc/ent/predicate"
	"ipc/ent/users"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OtpQuery is the builder for querying Otp entities.
type OtpQuery struct {
	config
	ctx        *QueryContext
	order      []otp.OrderOption
	inters     []Interceptor
	predicates []predicate.Otp
	withUsers  *UsersQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OtpQuery builder.
func (oq *OtpQuery) Where(ps ...predicate.Otp) *OtpQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit the number of records to be returned by this query.
func (oq *OtpQuery) Limit(limit int) *OtpQuery {
	oq.ctx.Limit = &limit
	return oq
}

// Offset to start from.
func (oq *OtpQuery) Offset(offset int) *OtpQuery {
	oq.ctx.Offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OtpQuery) Unique(unique bool) *OtpQuery {
	oq.ctx.Unique = &unique
	return oq
}

// Order specifies how the records should be ordered.
func (oq *OtpQuery) Order(o ...otp.OrderOption) *OtpQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryUsers chains the current query on the "users" edge.
func (oq *OtpQuery) QueryUsers() *UsersQuery {
	query := (&UsersClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(otp.Table, otp.FieldID, selector),
			sqlgraph.To(users.Table, users.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, otp.UsersTable, otp.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Otp entity from the query.
// Returns a *NotFoundError when no Otp was found.
func (oq *OtpQuery) First(ctx context.Context) (*Otp, error) {
	nodes, err := oq.Limit(1).All(setContextOp(ctx, oq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{otp.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OtpQuery) FirstX(ctx context.Context) *Otp {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Otp ID from the query.
// Returns a *NotFoundError when no Otp ID was found.
func (oq *OtpQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(1).IDs(setContextOp(ctx, oq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{otp.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OtpQuery) FirstIDX(ctx context.Context) int {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Otp entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Otp entity is found.
// Returns a *NotFoundError when no Otp entities are found.
func (oq *OtpQuery) Only(ctx context.Context) (*Otp, error) {
	nodes, err := oq.Limit(2).All(setContextOp(ctx, oq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{otp.Label}
	default:
		return nil, &NotSingularError{otp.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OtpQuery) OnlyX(ctx context.Context) *Otp {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Otp ID in the query.
// Returns a *NotSingularError when more than one Otp ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *OtpQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(2).IDs(setContextOp(ctx, oq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{otp.Label}
	default:
		err = &NotSingularError{otp.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OtpQuery) OnlyIDX(ctx context.Context) int {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Otps.
func (oq *OtpQuery) All(ctx context.Context) ([]*Otp, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryAll)
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Otp, *OtpQuery]()
	return withInterceptors[[]*Otp](ctx, oq, qr, oq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oq *OtpQuery) AllX(ctx context.Context) []*Otp {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Otp IDs.
func (oq *OtpQuery) IDs(ctx context.Context) (ids []int, err error) {
	if oq.ctx.Unique == nil && oq.path != nil {
		oq.Unique(true)
	}
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryIDs)
	if err = oq.Select(otp.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OtpQuery) IDsX(ctx context.Context) []int {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OtpQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryCount)
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oq, querierCount[*OtpQuery](), oq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OtpQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OtpQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryExist)
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OtpQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OtpQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OtpQuery) Clone() *OtpQuery {
	if oq == nil {
		return nil
	}
	return &OtpQuery{
		config:     oq.config,
		ctx:        oq.ctx.Clone(),
		order:      append([]otp.OrderOption{}, oq.order...),
		inters:     append([]Interceptor{}, oq.inters...),
		predicates: append([]predicate.Otp{}, oq.predicates...),
		withUsers:  oq.withUsers.Clone(),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OtpQuery) WithUsers(opts ...func(*UsersQuery)) *OtpQuery {
	query := (&UsersClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withUsers = query
	return oq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Code string `json:"code,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Otp.Query().
//		GroupBy(otp.FieldCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oq *OtpQuery) GroupBy(field string, fields ...string) *OtpGroupBy {
	oq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OtpGroupBy{build: oq}
	grbuild.flds = &oq.ctx.Fields
	grbuild.label = otp.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Code string `json:"code,omitempty"`
//	}
//
//	client.Otp.Query().
//		Select(otp.FieldCode).
//		Scan(ctx, &v)
func (oq *OtpQuery) Select(fields ...string) *OtpSelect {
	oq.ctx.Fields = append(oq.ctx.Fields, fields...)
	sbuild := &OtpSelect{OtpQuery: oq}
	sbuild.label = otp.Label
	sbuild.flds, sbuild.scan = &oq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OtpSelect configured with the given aggregations.
func (oq *OtpQuery) Aggregate(fns ...AggregateFunc) *OtpSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *OtpQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oq); err != nil {
				return err
			}
		}
	}
	for _, f := range oq.ctx.Fields {
		if !otp.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OtpQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Otp, error) {
	var (
		nodes       = []*Otp{}
		_spec       = oq.querySpec()
		loadedTypes = [1]bool{
			oq.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Otp).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Otp{config: oq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oq.withUsers; query != nil {
		if err := oq.loadUsers(ctx, query, nodes, nil,
			func(n *Otp, e *Users) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oq *OtpQuery) loadUsers(ctx context.Context, query *UsersQuery, nodes []*Otp, init func(*Otp), assign func(*Otp, *Users)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Otp)
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
	query.Where(users.IDIn(ids...))
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

func (oq *OtpQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	_spec.Node.Columns = oq.ctx.Fields
	if len(oq.ctx.Fields) > 0 {
		_spec.Unique = oq.ctx.Unique != nil && *oq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OtpQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(otp.Table, otp.Columns, sqlgraph.NewFieldSpec(otp.FieldID, field.TypeInt))
	_spec.From = oq.sql
	if unique := oq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oq.path != nil {
		_spec.Unique = true
	}
	if fields := oq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, otp.FieldID)
		for i := range fields {
			if fields[i] != otp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if oq.withUsers != nil {
			_spec.Node.AddColumnOnce(otp.FieldUserID)
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OtpQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(otp.Table)
	columns := oq.ctx.Fields
	if len(columns) == 0 {
		columns = otp.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.ctx.Unique != nil && *oq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OtpGroupBy is the group-by builder for Otp entities.
type OtpGroupBy struct {
	selector
	build *OtpQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OtpGroupBy) Aggregate(fns ...AggregateFunc) *OtpGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the selector query and scans the result into the given value.
func (ogb *OtpGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ogb.build.ctx, ent.OpQueryGroupBy)
	if err := ogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OtpQuery, *OtpGroupBy](ctx, ogb.build, ogb, ogb.build.inters, v)
}

func (ogb *OtpGroupBy) sqlScan(ctx context.Context, root *OtpQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ogb.flds)+len(ogb.fns))
		for _, f := range *ogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OtpSelect is the builder for selecting fields of Otp entities.
type OtpSelect struct {
	*OtpQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *OtpSelect) Aggregate(fns ...AggregateFunc) *OtpSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *OtpSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, os.ctx, ent.OpQuerySelect)
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OtpQuery, *OtpSelect](ctx, os.OtpQuery, os, os.inters, v)
}

func (os *OtpSelect) sqlScan(ctx context.Context, root *OtpQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
