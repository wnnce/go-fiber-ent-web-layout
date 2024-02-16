// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-fiber-ent-web-layout/ent/example"
	"go-fiber-ent-web-layout/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExampleUpdate is the builder for updating Example entities.
type ExampleUpdate struct {
	config
	hooks    []Hook
	mutation *ExampleMutation
}

// Where appends a list predicates to the ExampleUpdate builder.
func (eu *ExampleUpdate) Where(ps ...predicate.Example) *ExampleUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *ExampleUpdate) SetName(s string) *ExampleUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (eu *ExampleUpdate) SetNillableName(s *string) *ExampleUpdate {
	if s != nil {
		eu.SetName(*s)
	}
	return eu
}

// SetSummary sets the "summary" field.
func (eu *ExampleUpdate) SetSummary(s string) *ExampleUpdate {
	eu.mutation.SetSummary(s)
	return eu
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (eu *ExampleUpdate) SetNillableSummary(s *string) *ExampleUpdate {
	if s != nil {
		eu.SetSummary(*s)
	}
	return eu
}

// SetPrice sets the "price" field.
func (eu *ExampleUpdate) SetPrice(f float64) *ExampleUpdate {
	eu.mutation.ResetPrice()
	eu.mutation.SetPrice(f)
	return eu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (eu *ExampleUpdate) SetNillablePrice(f *float64) *ExampleUpdate {
	if f != nil {
		eu.SetPrice(*f)
	}
	return eu
}

// AddPrice adds f to the "price" field.
func (eu *ExampleUpdate) AddPrice(f float64) *ExampleUpdate {
	eu.mutation.AddPrice(f)
	return eu
}

// SetCreateTime sets the "create_time" field.
func (eu *ExampleUpdate) SetCreateTime(t time.Time) *ExampleUpdate {
	eu.mutation.SetCreateTime(t)
	return eu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (eu *ExampleUpdate) SetNillableCreateTime(t *time.Time) *ExampleUpdate {
	if t != nil {
		eu.SetCreateTime(*t)
	}
	return eu
}

// ClearCreateTime clears the value of the "create_time" field.
func (eu *ExampleUpdate) ClearCreateTime() *ExampleUpdate {
	eu.mutation.ClearCreateTime()
	return eu
}

// SetUpdateTime sets the "update_time" field.
func (eu *ExampleUpdate) SetUpdateTime(t time.Time) *ExampleUpdate {
	eu.mutation.SetUpdateTime(t)
	return eu
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (eu *ExampleUpdate) SetNillableUpdateTime(t *time.Time) *ExampleUpdate {
	if t != nil {
		eu.SetUpdateTime(*t)
	}
	return eu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (eu *ExampleUpdate) ClearUpdateTime() *ExampleUpdate {
	eu.mutation.ClearUpdateTime()
	return eu
}

// Mutation returns the ExampleMutation object of the builder.
func (eu *ExampleUpdate) Mutation() *ExampleMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *ExampleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *ExampleUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *ExampleUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *ExampleUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *ExampleUpdate) check() error {
	if v, ok := eu.mutation.Name(); ok {
		if err := example.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Example.name": %w`, err)}
		}
	}
	if v, ok := eu.mutation.Price(); ok {
		if err := example.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "Example.price": %w`, err)}
		}
	}
	return nil
}

func (eu *ExampleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(example.Table, example.Columns, sqlgraph.NewFieldSpec(example.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(example.FieldName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Summary(); ok {
		_spec.SetField(example.FieldSummary, field.TypeString, value)
	}
	if value, ok := eu.mutation.Price(); ok {
		_spec.SetField(example.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := eu.mutation.AddedPrice(); ok {
		_spec.AddField(example.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := eu.mutation.CreateTime(); ok {
		_spec.SetField(example.FieldCreateTime, field.TypeTime, value)
	}
	if eu.mutation.CreateTimeCleared() {
		_spec.ClearField(example.FieldCreateTime, field.TypeTime)
	}
	if value, ok := eu.mutation.UpdateTime(); ok {
		_spec.SetField(example.FieldUpdateTime, field.TypeTime, value)
	}
	if eu.mutation.UpdateTimeCleared() {
		_spec.ClearField(example.FieldUpdateTime, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{example.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// ExampleUpdateOne is the builder for updating a single Example entity.
type ExampleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExampleMutation
}

// SetName sets the "name" field.
func (euo *ExampleUpdateOne) SetName(s string) *ExampleUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (euo *ExampleUpdateOne) SetNillableName(s *string) *ExampleUpdateOne {
	if s != nil {
		euo.SetName(*s)
	}
	return euo
}

// SetSummary sets the "summary" field.
func (euo *ExampleUpdateOne) SetSummary(s string) *ExampleUpdateOne {
	euo.mutation.SetSummary(s)
	return euo
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (euo *ExampleUpdateOne) SetNillableSummary(s *string) *ExampleUpdateOne {
	if s != nil {
		euo.SetSummary(*s)
	}
	return euo
}

// SetPrice sets the "price" field.
func (euo *ExampleUpdateOne) SetPrice(f float64) *ExampleUpdateOne {
	euo.mutation.ResetPrice()
	euo.mutation.SetPrice(f)
	return euo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (euo *ExampleUpdateOne) SetNillablePrice(f *float64) *ExampleUpdateOne {
	if f != nil {
		euo.SetPrice(*f)
	}
	return euo
}

// AddPrice adds f to the "price" field.
func (euo *ExampleUpdateOne) AddPrice(f float64) *ExampleUpdateOne {
	euo.mutation.AddPrice(f)
	return euo
}

// SetCreateTime sets the "create_time" field.
func (euo *ExampleUpdateOne) SetCreateTime(t time.Time) *ExampleUpdateOne {
	euo.mutation.SetCreateTime(t)
	return euo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (euo *ExampleUpdateOne) SetNillableCreateTime(t *time.Time) *ExampleUpdateOne {
	if t != nil {
		euo.SetCreateTime(*t)
	}
	return euo
}

// ClearCreateTime clears the value of the "create_time" field.
func (euo *ExampleUpdateOne) ClearCreateTime() *ExampleUpdateOne {
	euo.mutation.ClearCreateTime()
	return euo
}

// SetUpdateTime sets the "update_time" field.
func (euo *ExampleUpdateOne) SetUpdateTime(t time.Time) *ExampleUpdateOne {
	euo.mutation.SetUpdateTime(t)
	return euo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (euo *ExampleUpdateOne) SetNillableUpdateTime(t *time.Time) *ExampleUpdateOne {
	if t != nil {
		euo.SetUpdateTime(*t)
	}
	return euo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (euo *ExampleUpdateOne) ClearUpdateTime() *ExampleUpdateOne {
	euo.mutation.ClearUpdateTime()
	return euo
}

// Mutation returns the ExampleMutation object of the builder.
func (euo *ExampleUpdateOne) Mutation() *ExampleMutation {
	return euo.mutation
}

// Where appends a list predicates to the ExampleUpdate builder.
func (euo *ExampleUpdateOne) Where(ps ...predicate.Example) *ExampleUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *ExampleUpdateOne) Select(field string, fields ...string) *ExampleUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Example entity.
func (euo *ExampleUpdateOne) Save(ctx context.Context) (*Example, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *ExampleUpdateOne) SaveX(ctx context.Context) *Example {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *ExampleUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *ExampleUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *ExampleUpdateOne) check() error {
	if v, ok := euo.mutation.Name(); ok {
		if err := example.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Example.name": %w`, err)}
		}
	}
	if v, ok := euo.mutation.Price(); ok {
		if err := example.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "Example.price": %w`, err)}
		}
	}
	return nil
}

func (euo *ExampleUpdateOne) sqlSave(ctx context.Context) (_node *Example, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(example.Table, example.Columns, sqlgraph.NewFieldSpec(example.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Example.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, example.FieldID)
		for _, f := range fields {
			if !example.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != example.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(example.FieldName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Summary(); ok {
		_spec.SetField(example.FieldSummary, field.TypeString, value)
	}
	if value, ok := euo.mutation.Price(); ok {
		_spec.SetField(example.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := euo.mutation.AddedPrice(); ok {
		_spec.AddField(example.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := euo.mutation.CreateTime(); ok {
		_spec.SetField(example.FieldCreateTime, field.TypeTime, value)
	}
	if euo.mutation.CreateTimeCleared() {
		_spec.ClearField(example.FieldCreateTime, field.TypeTime)
	}
	if value, ok := euo.mutation.UpdateTime(); ok {
		_spec.SetField(example.FieldUpdateTime, field.TypeTime, value)
	}
	if euo.mutation.UpdateTimeCleared() {
		_spec.ClearField(example.FieldUpdateTime, field.TypeTime)
	}
	_node = &Example{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{example.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}