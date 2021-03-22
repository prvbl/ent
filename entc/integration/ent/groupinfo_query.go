// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/group"
	"entgo.io/ent/entc/integration/ent/groupinfo"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// GroupInfoQuery is the builder for querying GroupInfo entities.
type GroupInfoQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.GroupInfo
	// eager-loading edges.
	withGroups *GroupQuery
	unique     *bool
	withLock   ent.LockType
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupInfoQuery builder.
func (giq *GroupInfoQuery) Where(ps ...predicate.GroupInfo) *GroupInfoQuery {
	giq.predicates = append(giq.predicates, ps...)
	return giq
}

// Limit adds a limit step to the query.
func (giq *GroupInfoQuery) Limit(limit int) *GroupInfoQuery {
	giq.limit = &limit
	return giq
}

// Offset adds an offset step to the query.
func (giq *GroupInfoQuery) Offset(offset int) *GroupInfoQuery {
	giq.offset = &offset
	return giq
}

// Order adds an order step to the query.
func (giq *GroupInfoQuery) Order(o ...OrderFunc) *GroupInfoQuery {
	giq.order = append(giq.order, o...)
	return giq
}

// QueryGroups chains the current query on the "groups" edge.
func (giq *GroupInfoQuery) QueryGroups() *GroupQuery {
	query := &GroupQuery{config: giq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := giq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := giq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(groupinfo.Table, groupinfo.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, groupinfo.GroupsTable, groupinfo.GroupsColumn),
		)
		fromU = sqlgraph.SetNeighbors(giq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroupInfo entity from the query.
// Returns a *NotFoundError when no GroupInfo was found.
func (giq *GroupInfoQuery) First(ctx context.Context) (*GroupInfo, error) {
	nodes, err := giq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{groupinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (giq *GroupInfoQuery) FirstX(ctx context.Context) *GroupInfo {
	node, err := giq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupInfo ID from the query.
// Returns a *NotFoundError when no GroupInfo ID was found.
func (giq *GroupInfoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = giq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{groupinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (giq *GroupInfoQuery) FirstIDX(ctx context.Context) int {
	id, err := giq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroupInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GroupInfo entity is not found.
// Returns a *NotFoundError when no GroupInfo entities are found.
func (giq *GroupInfoQuery) Only(ctx context.Context) (*GroupInfo, error) {
	nodes, err := giq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{groupinfo.Label}
	default:
		return nil, &NotSingularError{groupinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (giq *GroupInfoQuery) OnlyX(ctx context.Context) *GroupInfo {
	node, err := giq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroupInfo ID in the query.
// Returns a *NotSingularError when exactly one GroupInfo ID is not found.
// Returns a *NotFoundError when no entities are found.
func (giq *GroupInfoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = giq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{groupinfo.Label}
	default:
		err = &NotSingularError{groupinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (giq *GroupInfoQuery) OnlyIDX(ctx context.Context) int {
	id, err := giq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupInfos.
func (giq *GroupInfoQuery) All(ctx context.Context) ([]*GroupInfo, error) {
	if err := giq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return giq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (giq *GroupInfoQuery) AllX(ctx context.Context) []*GroupInfo {
	nodes, err := giq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupInfo IDs.
func (giq *GroupInfoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := giq.Select(groupinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (giq *GroupInfoQuery) IDsX(ctx context.Context) []int {
	ids, err := giq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (giq *GroupInfoQuery) Count(ctx context.Context) (int, error) {
	if err := giq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return giq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (giq *GroupInfoQuery) CountX(ctx context.Context) int {
	count, err := giq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (giq *GroupInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := giq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return giq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (giq *GroupInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := giq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (giq *GroupInfoQuery) Clone() *GroupInfoQuery {
	if giq == nil {
		return nil
	}
	return &GroupInfoQuery{
		config:     giq.config,
		limit:      giq.limit,
		offset:     giq.offset,
		order:      append([]OrderFunc{}, giq.order...),
		predicates: append([]predicate.GroupInfo{}, giq.predicates...),
		withGroups: giq.withGroups.Clone(),
		// clone intermediate query.
		sql:  giq.sql.Clone(),
		path: giq.path,
	}
}

// WithGroups tells the query-builder to eager-load the nodes that are connected to
// the "groups" edge. The optional arguments are used to configure the query builder of the edge.
func (giq *GroupInfoQuery) WithGroups(opts ...func(*GroupQuery)) *GroupInfoQuery {
	query := &GroupQuery{config: giq.config}
	for _, opt := range opts {
		opt(query)
	}
	giq.withGroups = query
	return giq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Desc string `json:"desc,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupInfo.Query().
//		GroupBy(groupinfo.FieldDesc).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (giq *GroupInfoQuery) GroupBy(field string, fields ...string) *GroupInfoGroupBy {
	group := &GroupInfoGroupBy{config: giq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := giq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return giq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Desc string `json:"desc,omitempty"`
//	}
//
//	client.GroupInfo.Query().
//		Select(groupinfo.FieldDesc).
//		Scan(ctx, &v)
//
func (giq *GroupInfoQuery) Select(field string, fields ...string) *GroupInfoSelect {
	giq.fields = append([]string{field}, fields...)
	return &GroupInfoSelect{GroupInfoQuery: giq}
}

func (giq *GroupInfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range giq.fields {
		if !groupinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if giq.path != nil {
		prev, err := giq.path(ctx)
		if err != nil {
			return err
		}
		giq.sql = prev
	}
	return nil
}

// LockForUpdate locks any rows read as if you issued an update for those rows.
func (giq *GroupInfoQuery) LockForUpdate() *GroupInfoQuery {
	giq.withLock = LockForUpdate
	return giq
}

// LockForShare sets a shared mode lock on any rows that are read.
func (giq *GroupInfoQuery) LockForShare() *GroupInfoQuery {
	giq.withLock = LockForShare
	return giq
}

func (giq *GroupInfoQuery) sqlAll(ctx context.Context) ([]*GroupInfo, error) {
	var (
		nodes       = []*GroupInfo{}
		_spec       = giq.querySpec()
		loadedTypes = [1]bool{
			giq.withGroups != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GroupInfo{config: giq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, giq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := giq.withGroups; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*GroupInfo)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Groups = []*Group{}
		}
		query.withFKs = true
		query.Where(predicate.Group(func(s *sql.Selector) {
			s.Where(sql.InValues(groupinfo.GroupsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.group_info
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "group_info" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "group_info" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Groups = append(node.Edges.Groups, n)
		}
	}

	return nodes, nil
}

func (giq *GroupInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := giq.querySpec()
	return sqlgraph.CountNodes(ctx, giq.driver, _spec)
}

func (giq *GroupInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := giq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (giq *GroupInfoQuery) querySpec() *sqlgraph.QuerySpec {
	unique := true
	if giq.unique != nil {
		unique = *giq.unique
	}
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupinfo.Table,
			Columns: groupinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupinfo.FieldID,
			},
		},
		From:     giq.sql,
		Unique:   unique,
		WithLock: giq.withLock,
	}
	if fields := giq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupinfo.FieldID)
		for i := range fields {
			if fields[i] != groupinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := giq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := giq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := giq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := giq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, groupinfo.ValidColumn)
			}
		}
	}
	return _spec
}

func (giq *GroupInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(giq.driver.Dialect())
	t1 := builder.Table(groupinfo.Table)
	selector := builder.Select(t1.Columns(groupinfo.Columns...)...).From(t1)
	if giq.sql != nil {
		selector = giq.sql
		selector.Select(selector.Columns(groupinfo.Columns...)...)
	}
	for _, p := range giq.predicates {
		p(selector)
	}
	for _, p := range giq.order {
		p(selector, groupinfo.ValidColumn)
	}
	if offset := giq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := giq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupInfoGroupBy is the group-by builder for GroupInfo entities.
type GroupInfoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gigb *GroupInfoGroupBy) Aggregate(fns ...AggregateFunc) *GroupInfoGroupBy {
	gigb.fns = append(gigb.fns, fns...)
	return gigb
}

// Scan applies the group-by query and scans the result into the given value.
func (gigb *GroupInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gigb.path(ctx)
	if err != nil {
		return err
	}
	gigb.sql = query
	return gigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gigb *GroupInfoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) StringsX(ctx context.Context) []string {
	v, err := gigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gigb *GroupInfoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupinfo.Label}
	default:
		err = fmt.Errorf("ent: GroupInfoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) StringX(ctx context.Context) string {
	v, err := gigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gigb *GroupInfoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) IntsX(ctx context.Context) []int {
	v, err := gigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gigb *GroupInfoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupinfo.Label}
	default:
		err = fmt.Errorf("ent: GroupInfoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) IntX(ctx context.Context) int {
	v, err := gigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gigb *GroupInfoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gigb *GroupInfoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupinfo.Label}
	default:
		err = fmt.Errorf("ent: GroupInfoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gigb *GroupInfoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gigb *GroupInfoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupinfo.Label}
	default:
		err = fmt.Errorf("ent: GroupInfoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) BoolX(ctx context.Context) bool {
	v, err := gigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gigb *GroupInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gigb.fields {
		if !groupinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gigb *GroupInfoGroupBy) sqlQuery() *sql.Selector {
	selector := gigb.sql
	columns := make([]string, 0, len(gigb.fields)+len(gigb.fns))
	columns = append(columns, gigb.fields...)
	for _, fn := range gigb.fns {
		columns = append(columns, fn(selector, groupinfo.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(gigb.fields...)
}

// GroupInfoSelect is the builder for selecting fields of GroupInfo entities.
type GroupInfoSelect struct {
	*GroupInfoQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gis *GroupInfoSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gis.prepareQuery(ctx); err != nil {
		return err
	}
	gis.sql = gis.GroupInfoQuery.sqlQuery(ctx)
	return gis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gis *GroupInfoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GroupInfoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gis *GroupInfoSelect) StringsX(ctx context.Context) []string {
	v, err := gis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupinfo.Label}
	default:
		err = fmt.Errorf("ent: GroupInfoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gis *GroupInfoSelect) StringX(ctx context.Context) string {
	v, err := gis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GroupInfoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gis *GroupInfoSelect) IntsX(ctx context.Context) []int {
	v, err := gis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupinfo.Label}
	default:
		err = fmt.Errorf("ent: GroupInfoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gis *GroupInfoSelect) IntX(ctx context.Context) int {
	v, err := gis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GroupInfoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gis *GroupInfoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupinfo.Label}
	default:
		err = fmt.Errorf("ent: GroupInfoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gis *GroupInfoSelect) Float64X(ctx context.Context) float64 {
	v, err := gis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GroupInfoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gis *GroupInfoSelect) BoolsX(ctx context.Context) []bool {
	v, err := gis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupinfo.Label}
	default:
		err = fmt.Errorf("ent: GroupInfoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gis *GroupInfoSelect) BoolX(ctx context.Context) bool {
	v, err := gis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gis *GroupInfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gis.sqlQuery().Query()
	if err := gis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gis *GroupInfoSelect) sqlQuery() sql.Querier {
	selector := gis.sql
	selector.Select(selector.Columns(gis.fields...)...)
	return selector
}
