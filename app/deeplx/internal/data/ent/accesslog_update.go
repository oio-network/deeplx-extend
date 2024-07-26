// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/oio-network/deeplx-extend/app/deeplx/internal/data/ent/accesslog"
	"github.com/oio-network/deeplx-extend/app/deeplx/internal/data/ent/predicate"
	"github.com/oio-network/deeplx-extend/app/deeplx/internal/data/ent/user"
)

// AccessLogUpdate is the builder for updating AccessLog entities.
type AccessLogUpdate struct {
	config
	hooks     []Hook
	mutation  *AccessLogMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AccessLogUpdate builder.
func (alu *AccessLogUpdate) Where(ps ...predicate.AccessLog) *AccessLogUpdate {
	alu.mutation.Where(ps...)
	return alu
}

// SetUserID sets the "user_id" field.
func (alu *AccessLogUpdate) SetUserID(i int64) *AccessLogUpdate {
	alu.mutation.SetUserID(i)
	return alu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (alu *AccessLogUpdate) SetNillableUserID(i *int64) *AccessLogUpdate {
	if i != nil {
		alu.SetUserID(*i)
	}
	return alu
}

// ClearUserID clears the value of the "user_id" field.
func (alu *AccessLogUpdate) ClearUserID() *AccessLogUpdate {
	alu.mutation.ClearUserID()
	return alu
}

// SetIP sets the "ip" field.
func (alu *AccessLogUpdate) SetIP(s string) *AccessLogUpdate {
	alu.mutation.SetIP(s)
	return alu
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (alu *AccessLogUpdate) SetNillableIP(s *string) *AccessLogUpdate {
	if s != nil {
		alu.SetIP(*s)
	}
	return alu
}

// SetCountryName sets the "country_name" field.
func (alu *AccessLogUpdate) SetCountryName(s string) *AccessLogUpdate {
	alu.mutation.SetCountryName(s)
	return alu
}

// SetNillableCountryName sets the "country_name" field if the given value is not nil.
func (alu *AccessLogUpdate) SetNillableCountryName(s *string) *AccessLogUpdate {
	if s != nil {
		alu.SetCountryName(*s)
	}
	return alu
}

// SetCountryCode sets the "country_code" field.
func (alu *AccessLogUpdate) SetCountryCode(s string) *AccessLogUpdate {
	alu.mutation.SetCountryCode(s)
	return alu
}

// SetNillableCountryCode sets the "country_code" field if the given value is not nil.
func (alu *AccessLogUpdate) SetNillableCountryCode(s *string) *AccessLogUpdate {
	if s != nil {
		alu.SetCountryCode(*s)
	}
	return alu
}

// SetOwnerUserID sets the "owner_user" edge to the User entity by ID.
func (alu *AccessLogUpdate) SetOwnerUserID(id int64) *AccessLogUpdate {
	alu.mutation.SetOwnerUserID(id)
	return alu
}

// SetNillableOwnerUserID sets the "owner_user" edge to the User entity by ID if the given value is not nil.
func (alu *AccessLogUpdate) SetNillableOwnerUserID(id *int64) *AccessLogUpdate {
	if id != nil {
		alu = alu.SetOwnerUserID(*id)
	}
	return alu
}

// SetOwnerUser sets the "owner_user" edge to the User entity.
func (alu *AccessLogUpdate) SetOwnerUser(u *User) *AccessLogUpdate {
	return alu.SetOwnerUserID(u.ID)
}

// Mutation returns the AccessLogMutation object of the builder.
func (alu *AccessLogUpdate) Mutation() *AccessLogMutation {
	return alu.mutation
}

// ClearOwnerUser clears the "owner_user" edge to the User entity.
func (alu *AccessLogUpdate) ClearOwnerUser() *AccessLogUpdate {
	alu.mutation.ClearOwnerUser()
	return alu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (alu *AccessLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, alu.sqlSave, alu.mutation, alu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (alu *AccessLogUpdate) SaveX(ctx context.Context) int {
	affected, err := alu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (alu *AccessLogUpdate) Exec(ctx context.Context) error {
	_, err := alu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alu *AccessLogUpdate) ExecX(ctx context.Context) {
	if err := alu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (alu *AccessLogUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AccessLogUpdate {
	alu.modifiers = append(alu.modifiers, modifiers...)
	return alu
}

func (alu *AccessLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(accesslog.Table, accesslog.Columns, sqlgraph.NewFieldSpec(accesslog.FieldID, field.TypeInt64))
	if ps := alu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := alu.mutation.IP(); ok {
		_spec.SetField(accesslog.FieldIP, field.TypeString, value)
	}
	if value, ok := alu.mutation.CountryName(); ok {
		_spec.SetField(accesslog.FieldCountryName, field.TypeString, value)
	}
	if value, ok := alu.mutation.CountryCode(); ok {
		_spec.SetField(accesslog.FieldCountryCode, field.TypeString, value)
	}
	if alu.mutation.OwnerUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesslog.OwnerUserTable,
			Columns: []string{accesslog.OwnerUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := alu.mutation.OwnerUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesslog.OwnerUserTable,
			Columns: []string{accesslog.OwnerUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(alu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, alu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesslog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	alu.mutation.done = true
	return n, nil
}

// AccessLogUpdateOne is the builder for updating a single AccessLog entity.
type AccessLogUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AccessLogMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUserID sets the "user_id" field.
func (aluo *AccessLogUpdateOne) SetUserID(i int64) *AccessLogUpdateOne {
	aluo.mutation.SetUserID(i)
	return aluo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (aluo *AccessLogUpdateOne) SetNillableUserID(i *int64) *AccessLogUpdateOne {
	if i != nil {
		aluo.SetUserID(*i)
	}
	return aluo
}

// ClearUserID clears the value of the "user_id" field.
func (aluo *AccessLogUpdateOne) ClearUserID() *AccessLogUpdateOne {
	aluo.mutation.ClearUserID()
	return aluo
}

// SetIP sets the "ip" field.
func (aluo *AccessLogUpdateOne) SetIP(s string) *AccessLogUpdateOne {
	aluo.mutation.SetIP(s)
	return aluo
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (aluo *AccessLogUpdateOne) SetNillableIP(s *string) *AccessLogUpdateOne {
	if s != nil {
		aluo.SetIP(*s)
	}
	return aluo
}

// SetCountryName sets the "country_name" field.
func (aluo *AccessLogUpdateOne) SetCountryName(s string) *AccessLogUpdateOne {
	aluo.mutation.SetCountryName(s)
	return aluo
}

// SetNillableCountryName sets the "country_name" field if the given value is not nil.
func (aluo *AccessLogUpdateOne) SetNillableCountryName(s *string) *AccessLogUpdateOne {
	if s != nil {
		aluo.SetCountryName(*s)
	}
	return aluo
}

// SetCountryCode sets the "country_code" field.
func (aluo *AccessLogUpdateOne) SetCountryCode(s string) *AccessLogUpdateOne {
	aluo.mutation.SetCountryCode(s)
	return aluo
}

// SetNillableCountryCode sets the "country_code" field if the given value is not nil.
func (aluo *AccessLogUpdateOne) SetNillableCountryCode(s *string) *AccessLogUpdateOne {
	if s != nil {
		aluo.SetCountryCode(*s)
	}
	return aluo
}

// SetOwnerUserID sets the "owner_user" edge to the User entity by ID.
func (aluo *AccessLogUpdateOne) SetOwnerUserID(id int64) *AccessLogUpdateOne {
	aluo.mutation.SetOwnerUserID(id)
	return aluo
}

// SetNillableOwnerUserID sets the "owner_user" edge to the User entity by ID if the given value is not nil.
func (aluo *AccessLogUpdateOne) SetNillableOwnerUserID(id *int64) *AccessLogUpdateOne {
	if id != nil {
		aluo = aluo.SetOwnerUserID(*id)
	}
	return aluo
}

// SetOwnerUser sets the "owner_user" edge to the User entity.
func (aluo *AccessLogUpdateOne) SetOwnerUser(u *User) *AccessLogUpdateOne {
	return aluo.SetOwnerUserID(u.ID)
}

// Mutation returns the AccessLogMutation object of the builder.
func (aluo *AccessLogUpdateOne) Mutation() *AccessLogMutation {
	return aluo.mutation
}

// ClearOwnerUser clears the "owner_user" edge to the User entity.
func (aluo *AccessLogUpdateOne) ClearOwnerUser() *AccessLogUpdateOne {
	aluo.mutation.ClearOwnerUser()
	return aluo
}

// Where appends a list predicates to the AccessLogUpdate builder.
func (aluo *AccessLogUpdateOne) Where(ps ...predicate.AccessLog) *AccessLogUpdateOne {
	aluo.mutation.Where(ps...)
	return aluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aluo *AccessLogUpdateOne) Select(field string, fields ...string) *AccessLogUpdateOne {
	aluo.fields = append([]string{field}, fields...)
	return aluo
}

// Save executes the query and returns the updated AccessLog entity.
func (aluo *AccessLogUpdateOne) Save(ctx context.Context) (*AccessLog, error) {
	return withHooks(ctx, aluo.sqlSave, aluo.mutation, aluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aluo *AccessLogUpdateOne) SaveX(ctx context.Context) *AccessLog {
	node, err := aluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aluo *AccessLogUpdateOne) Exec(ctx context.Context) error {
	_, err := aluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aluo *AccessLogUpdateOne) ExecX(ctx context.Context) {
	if err := aluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aluo *AccessLogUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AccessLogUpdateOne {
	aluo.modifiers = append(aluo.modifiers, modifiers...)
	return aluo
}

func (aluo *AccessLogUpdateOne) sqlSave(ctx context.Context) (_node *AccessLog, err error) {
	_spec := sqlgraph.NewUpdateSpec(accesslog.Table, accesslog.Columns, sqlgraph.NewFieldSpec(accesslog.FieldID, field.TypeInt64))
	id, ok := aluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AccessLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accesslog.FieldID)
		for _, f := range fields {
			if !accesslog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != accesslog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aluo.mutation.IP(); ok {
		_spec.SetField(accesslog.FieldIP, field.TypeString, value)
	}
	if value, ok := aluo.mutation.CountryName(); ok {
		_spec.SetField(accesslog.FieldCountryName, field.TypeString, value)
	}
	if value, ok := aluo.mutation.CountryCode(); ok {
		_spec.SetField(accesslog.FieldCountryCode, field.TypeString, value)
	}
	if aluo.mutation.OwnerUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesslog.OwnerUserTable,
			Columns: []string{accesslog.OwnerUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aluo.mutation.OwnerUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesslog.OwnerUserTable,
			Columns: []string{accesslog.OwnerUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(aluo.modifiers...)
	_node = &AccessLog{config: aluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesslog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aluo.mutation.done = true
	return _node, nil
}
