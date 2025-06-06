// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/item"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks     []Hook
	mutation  *ItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ItemUpdate builder.
func (_u *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetText sets the "text" field.
func (_u *ItemUpdate) SetText(v string) *ItemUpdate {
	_u.mutation.SetText(v)
	return _u
}

// SetNillableText sets the "text" field if the given value is not nil.
func (_u *ItemUpdate) SetNillableText(v *string) *ItemUpdate {
	if v != nil {
		_u.SetText(*v)
	}
	return _u
}

// ClearText clears the value of the "text" field.
func (_u *ItemUpdate) ClearText() *ItemUpdate {
	_u.mutation.ClearText()
	return _u
}

// Mutation returns the ItemMutation object of the builder.
func (_u *ItemUpdate) Mutation() *ItemMutation {
	return _u.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *ItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *ItemUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *ItemUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *ItemUpdate) check() error {
	if v, ok := _u.mutation.Text(); ok {
		if err := item.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Item.text": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (_u *ItemUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ItemUpdate {
	_u.modifiers = append(_u.modifiers, modifiers...)
	return _u
}

func (_u *ItemUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeString))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Text(); ok {
		_spec.SetField(item.FieldText, field.TypeString, value)
	}
	if _u.mutation.TextCleared() {
		_spec.ClearField(item.FieldText, field.TypeString)
	}
	_spec.AddModifiers(_u.modifiers...)
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetText sets the "text" field.
func (_u *ItemUpdateOne) SetText(v string) *ItemUpdateOne {
	_u.mutation.SetText(v)
	return _u
}

// SetNillableText sets the "text" field if the given value is not nil.
func (_u *ItemUpdateOne) SetNillableText(v *string) *ItemUpdateOne {
	if v != nil {
		_u.SetText(*v)
	}
	return _u
}

// ClearText clears the value of the "text" field.
func (_u *ItemUpdateOne) ClearText() *ItemUpdateOne {
	_u.mutation.ClearText()
	return _u
}

// Mutation returns the ItemMutation object of the builder.
func (_u *ItemUpdateOne) Mutation() *ItemMutation {
	return _u.mutation
}

// Where appends a list predicates to the ItemUpdate builder.
func (_u *ItemUpdateOne) Where(ps ...predicate.Item) *ItemUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *ItemUpdateOne) Select(field string, fields ...string) *ItemUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Item entity.
func (_u *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *ItemUpdateOne) check() error {
	if v, ok := _u.mutation.Text(); ok {
		if err := item.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Item.text": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (_u *ItemUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ItemUpdateOne {
	_u.modifiers = append(_u.modifiers, modifiers...)
	return _u
}

func (_u *ItemUpdateOne) sqlSave(ctx context.Context) (_node *Item, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeString))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Item.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for _, f := range fields {
			if !item.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Text(); ok {
		_spec.SetField(item.FieldText, field.TypeString, value)
	}
	if _u.mutation.TextCleared() {
		_spec.ClearField(item.FieldText, field.TypeString)
	}
	_spec.AddModifiers(_u.modifiers...)
	_node = &Item{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
