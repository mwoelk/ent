// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/friendship"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/schema/field"
)

// FriendshipCreate is the builder for creating a Friendship entity.
type FriendshipCreate struct {
	config
	mutation *FriendshipMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetWeight sets the "weight" field.
func (_c *FriendshipCreate) SetWeight(v int) *FriendshipCreate {
	_c.mutation.SetWeight(v)
	return _c
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (_c *FriendshipCreate) SetNillableWeight(v *int) *FriendshipCreate {
	if v != nil {
		_c.SetWeight(*v)
	}
	return _c
}

// SetCreatedAt sets the "created_at" field.
func (_c *FriendshipCreate) SetCreatedAt(v time.Time) *FriendshipCreate {
	_c.mutation.SetCreatedAt(v)
	return _c
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (_c *FriendshipCreate) SetNillableCreatedAt(v *time.Time) *FriendshipCreate {
	if v != nil {
		_c.SetCreatedAt(*v)
	}
	return _c
}

// SetUserID sets the "user_id" field.
func (_c *FriendshipCreate) SetUserID(v int) *FriendshipCreate {
	_c.mutation.SetUserID(v)
	return _c
}

// SetFriendID sets the "friend_id" field.
func (_c *FriendshipCreate) SetFriendID(v int) *FriendshipCreate {
	_c.mutation.SetFriendID(v)
	return _c
}

// SetUser sets the "user" edge to the User entity.
func (_c *FriendshipCreate) SetUser(v *User) *FriendshipCreate {
	return _c.SetUserID(v.ID)
}

// SetFriend sets the "friend" edge to the User entity.
func (_c *FriendshipCreate) SetFriend(v *User) *FriendshipCreate {
	return _c.SetFriendID(v.ID)
}

// Mutation returns the FriendshipMutation object of the builder.
func (_c *FriendshipCreate) Mutation() *FriendshipMutation {
	return _c.mutation
}

// Save creates the Friendship in the database.
func (_c *FriendshipCreate) Save(ctx context.Context) (*Friendship, error) {
	_c.defaults()
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *FriendshipCreate) SaveX(ctx context.Context) *Friendship {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *FriendshipCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *FriendshipCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *FriendshipCreate) defaults() {
	if _, ok := _c.mutation.Weight(); !ok {
		v := friendship.DefaultWeight
		_c.mutation.SetWeight(v)
	}
	if _, ok := _c.mutation.CreatedAt(); !ok {
		v := friendship.DefaultCreatedAt()
		_c.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *FriendshipCreate) check() error {
	if _, ok := _c.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "Friendship.weight"`)}
	}
	if _, ok := _c.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Friendship.created_at"`)}
	}
	if _, ok := _c.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Friendship.user_id"`)}
	}
	if _, ok := _c.mutation.FriendID(); !ok {
		return &ValidationError{Name: "friend_id", err: errors.New(`ent: missing required field "Friendship.friend_id"`)}
	}
	if len(_c.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Friendship.user"`)}
	}
	if len(_c.mutation.FriendIDs()) == 0 {
		return &ValidationError{Name: "friend", err: errors.New(`ent: missing required edge "Friendship.friend"`)}
	}
	return nil
}

func (_c *FriendshipCreate) sqlSave(ctx context.Context) (*Friendship, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	_node, _spec := _c.createSpec()
	if err := sqlgraph.CreateNode(ctx, _c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *FriendshipCreate) createSpec() (*Friendship, *sqlgraph.CreateSpec) {
	var (
		_node = &Friendship{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(friendship.Table, sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt))
	)
	_spec.OnConflict = _c.conflict
	if value, ok := _c.mutation.Weight(); ok {
		_spec.SetField(friendship.FieldWeight, field.TypeInt, value)
		_node.Weight = value
	}
	if value, ok := _c.mutation.CreatedAt(); ok {
		_spec.SetField(friendship.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := _c.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.UserTable,
			Columns: []string{friendship.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.FriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.FriendTable,
			Columns: []string{friendship.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FriendID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Friendship.Create().
//		SetWeight(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FriendshipUpsert) {
//			SetWeight(v+v).
//		}).
//		Exec(ctx)
func (_c *FriendshipCreate) OnConflict(opts ...sql.ConflictOption) *FriendshipUpsertOne {
	_c.conflict = opts
	return &FriendshipUpsertOne{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *FriendshipCreate) OnConflictColumns(columns ...string) *FriendshipUpsertOne {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &FriendshipUpsertOne{
		create: _c,
	}
}

type (
	// FriendshipUpsertOne is the builder for "upsert"-ing
	//  one Friendship node.
	FriendshipUpsertOne struct {
		create *FriendshipCreate
	}

	// FriendshipUpsert is the "OnConflict" setter.
	FriendshipUpsert struct {
		*sql.UpdateSet
	}
)

// SetWeight sets the "weight" field.
func (u *FriendshipUpsert) SetWeight(v int) *FriendshipUpsert {
	u.Set(friendship.FieldWeight, v)
	return u
}

// UpdateWeight sets the "weight" field to the value that was provided on create.
func (u *FriendshipUpsert) UpdateWeight() *FriendshipUpsert {
	u.SetExcluded(friendship.FieldWeight)
	return u
}

// AddWeight adds v to the "weight" field.
func (u *FriendshipUpsert) AddWeight(v int) *FriendshipUpsert {
	u.Add(friendship.FieldWeight, v)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FriendshipUpsert) SetCreatedAt(v time.Time) *FriendshipUpsert {
	u.Set(friendship.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FriendshipUpsert) UpdateCreatedAt() *FriendshipUpsert {
	u.SetExcluded(friendship.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FriendshipUpsertOne) UpdateNewValues() *FriendshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.UserID(); exists {
			s.SetIgnore(friendship.FieldUserID)
		}
		if _, exists := u.create.mutation.FriendID(); exists {
			s.SetIgnore(friendship.FieldFriendID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Friendship.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FriendshipUpsertOne) Ignore() *FriendshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FriendshipUpsertOne) DoNothing() *FriendshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FriendshipCreate.OnConflict
// documentation for more info.
func (u *FriendshipUpsertOne) Update(set func(*FriendshipUpsert)) *FriendshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FriendshipUpsert{UpdateSet: update})
	}))
	return u
}

// SetWeight sets the "weight" field.
func (u *FriendshipUpsertOne) SetWeight(v int) *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetWeight(v)
	})
}

// AddWeight adds v to the "weight" field.
func (u *FriendshipUpsertOne) AddWeight(v int) *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.AddWeight(v)
	})
}

// UpdateWeight sets the "weight" field to the value that was provided on create.
func (u *FriendshipUpsertOne) UpdateWeight() *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateWeight()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FriendshipUpsertOne) SetCreatedAt(v time.Time) *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FriendshipUpsertOne) UpdateCreatedAt() *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FriendshipUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FriendshipCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FriendshipUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FriendshipUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FriendshipUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FriendshipCreateBulk is the builder for creating many Friendship entities in bulk.
type FriendshipCreateBulk struct {
	config
	err      error
	builders []*FriendshipCreate
	conflict []sql.ConflictOption
}

// Save creates the Friendship entities in the database.
func (_c *FriendshipCreateBulk) Save(ctx context.Context) ([]*Friendship, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Friendship, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FriendshipMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, _c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = _c.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, _c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, _c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (_c *FriendshipCreateBulk) SaveX(ctx context.Context) []*Friendship {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *FriendshipCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *FriendshipCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Friendship.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FriendshipUpsert) {
//			SetWeight(v+v).
//		}).
//		Exec(ctx)
func (_c *FriendshipCreateBulk) OnConflict(opts ...sql.ConflictOption) *FriendshipUpsertBulk {
	_c.conflict = opts
	return &FriendshipUpsertBulk{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *FriendshipCreateBulk) OnConflictColumns(columns ...string) *FriendshipUpsertBulk {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &FriendshipUpsertBulk{
		create: _c,
	}
}

// FriendshipUpsertBulk is the builder for "upsert"-ing
// a bulk of Friendship nodes.
type FriendshipUpsertBulk struct {
	create *FriendshipCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FriendshipUpsertBulk) UpdateNewValues() *FriendshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.UserID(); exists {
				s.SetIgnore(friendship.FieldUserID)
			}
			if _, exists := b.mutation.FriendID(); exists {
				s.SetIgnore(friendship.FieldFriendID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FriendshipUpsertBulk) Ignore() *FriendshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FriendshipUpsertBulk) DoNothing() *FriendshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FriendshipCreateBulk.OnConflict
// documentation for more info.
func (u *FriendshipUpsertBulk) Update(set func(*FriendshipUpsert)) *FriendshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FriendshipUpsert{UpdateSet: update})
	}))
	return u
}

// SetWeight sets the "weight" field.
func (u *FriendshipUpsertBulk) SetWeight(v int) *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetWeight(v)
	})
}

// AddWeight adds v to the "weight" field.
func (u *FriendshipUpsertBulk) AddWeight(v int) *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.AddWeight(v)
	})
}

// UpdateWeight sets the "weight" field to the value that was provided on create.
func (u *FriendshipUpsertBulk) UpdateWeight() *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateWeight()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FriendshipUpsertBulk) SetCreatedAt(v time.Time) *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FriendshipUpsertBulk) UpdateCreatedAt() *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FriendshipUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FriendshipCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FriendshipCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FriendshipUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
