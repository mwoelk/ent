// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/examples/o2mrecur/ent/node"
	"entgo.io/ent/examples/o2mrecur/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeNode = "Node"
)

// NodeMutation represents an operation that mutates the Node nodes in the graph.
type NodeMutation struct {
	config
	op              Op
	typ             string
	id              *int
	value           *int
	addvalue        *int
	clearedFields   map[string]struct{}
	parent          *int
	clearedparent   bool
	children        map[int]struct{}
	removedchildren map[int]struct{}
	clearedchildren bool
	done            bool
	oldValue        func(context.Context) (*Node, error)
	predicates      []predicate.Node
}

var _ ent.Mutation = (*NodeMutation)(nil)

// nodeOption allows management of the mutation configuration using functional options.
type nodeOption func(*NodeMutation)

// newNodeMutation creates new mutation for the Node entity.
func newNodeMutation(c config, op Op, opts ...nodeOption) *NodeMutation {
	m := &NodeMutation{
		config:        c,
		op:            op,
		typ:           TypeNode,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withNodeID sets the ID field of the mutation.
func withNodeID(id int) nodeOption {
	return func(m *NodeMutation) {
		var (
			err   error
			once  sync.Once
			value *Node
		)
		m.oldValue = func(ctx context.Context) (*Node, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Node.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withNode sets the old Node of the mutation.
func withNode(node *Node) nodeOption {
	return func(m *NodeMutation) {
		m.oldValue = func(context.Context) (*Node, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m NodeMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m NodeMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *NodeMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *NodeMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists && len(m.predicates) > 0 {
			m.predicates = append(m.predicates, node.ID(id))
			return m.Client().Node.Query().Where(m.predicates...).IDs(ctx)
		} else if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Node.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetValue sets the "value" field.
func (m *NodeMutation) SetValue(i int) {
	m.value = &i
	m.addvalue = nil
}

// Value returns the value of the "value" field in the mutation.
func (m *NodeMutation) Value() (r int, exists bool) {
	v := m.value
	if v == nil {
		return
	}
	return *v, true
}

// OldValue returns the old "value" field's value of the Node entity.
// If the Node object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *NodeMutation) OldValue(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldValue is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldValue requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldValue: %w", err)
	}
	return oldValue.Value, nil
}

// AddValue adds i to the "value" field.
func (m *NodeMutation) AddValue(i int) {
	if m.addvalue != nil {
		*m.addvalue += i
	} else {
		m.addvalue = &i
	}
}

// AddedValue returns the value that was added to the "value" field in this mutation.
func (m *NodeMutation) AddedValue() (r int, exists bool) {
	v := m.addvalue
	if v == nil {
		return
	}
	return *v, true
}

// ResetValue resets all changes to the "value" field.
func (m *NodeMutation) ResetValue() {
	m.value = nil
	m.addvalue = nil
}

// SetParentID sets the "parent_id" field.
func (m *NodeMutation) SetParentID(i int) {
	m.parent = &i
}

// ParentID returns the value of the "parent_id" field in the mutation.
func (m *NodeMutation) ParentID() (r int, exists bool) {
	v := m.parent
	if v == nil {
		return
	}
	return *v, true
}

// OldParentID returns the old "parent_id" field's value of the Node entity.
// If the Node object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *NodeMutation) OldParentID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldParentID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldParentID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldParentID: %w", err)
	}
	return oldValue.ParentID, nil
}

// ClearParentID clears the value of the "parent_id" field.
func (m *NodeMutation) ClearParentID() {
	m.parent = nil
	m.clearedFields[node.FieldParentID] = struct{}{}
}

// ParentIDCleared returns if the "parent_id" field was cleared in this mutation.
func (m *NodeMutation) ParentIDCleared() bool {
	_, ok := m.clearedFields[node.FieldParentID]
	return ok
}

// ResetParentID resets all changes to the "parent_id" field.
func (m *NodeMutation) ResetParentID() {
	m.parent = nil
	delete(m.clearedFields, node.FieldParentID)
}

// ClearParent clears the "parent" edge to the Node entity.
func (m *NodeMutation) ClearParent() {
	m.clearedparent = true
	m.clearedFields[node.FieldParentID] = struct{}{}
}

// ParentCleared reports if the "parent" edge to the Node entity was cleared.
func (m *NodeMutation) ParentCleared() bool {
	return m.ParentIDCleared() || m.clearedparent
}

// ParentIDs returns the "parent" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ParentID instead. It exists only for internal usage by the builders.
func (m *NodeMutation) ParentIDs() (ids []int) {
	if id := m.parent; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetParent resets all changes to the "parent" edge.
func (m *NodeMutation) ResetParent() {
	m.parent = nil
	m.clearedparent = false
}

// AddChildIDs adds the "children" edge to the Node entity by ids.
func (m *NodeMutation) AddChildIDs(ids ...int) {
	if m.children == nil {
		m.children = make(map[int]struct{})
	}
	for i := range ids {
		m.children[ids[i]] = struct{}{}
	}
}

// ClearChildren clears the "children" edge to the Node entity.
func (m *NodeMutation) ClearChildren() {
	m.clearedchildren = true
}

// ChildrenCleared reports if the "children" edge to the Node entity was cleared.
func (m *NodeMutation) ChildrenCleared() bool {
	return m.clearedchildren
}

// RemoveChildIDs removes the "children" edge to the Node entity by IDs.
func (m *NodeMutation) RemoveChildIDs(ids ...int) {
	if m.removedchildren == nil {
		m.removedchildren = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.children, ids[i])
		m.removedchildren[ids[i]] = struct{}{}
	}
}

// RemovedChildren returns the removed IDs of the "children" edge to the Node entity.
func (m *NodeMutation) RemovedChildrenIDs() (ids []int) {
	for id := range m.removedchildren {
		ids = append(ids, id)
	}
	return
}

// ChildrenIDs returns the "children" edge IDs in the mutation.
func (m *NodeMutation) ChildrenIDs() (ids []int) {
	for id := range m.children {
		ids = append(ids, id)
	}
	return
}

// ResetChildren resets all changes to the "children" edge.
func (m *NodeMutation) ResetChildren() {
	m.children = nil
	m.clearedchildren = false
	m.removedchildren = nil
}

// Where appends a list predicates to the NodeMutation builder.
func (m *NodeMutation) Where(ps ...predicate.Node) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the NodeMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *NodeMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Node, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *NodeMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *NodeMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Node).
func (m *NodeMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *NodeMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.value != nil {
		fields = append(fields, node.FieldValue)
	}
	if m.parent != nil {
		fields = append(fields, node.FieldParentID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *NodeMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case node.FieldValue:
		return m.Value()
	case node.FieldParentID:
		return m.ParentID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *NodeMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case node.FieldValue:
		return m.OldValue(ctx)
	case node.FieldParentID:
		return m.OldParentID(ctx)
	}
	return nil, fmt.Errorf("unknown Node field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *NodeMutation) SetField(name string, value ent.Value) error {
	switch name {
	case node.FieldValue:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetValue(v)
		return nil
	case node.FieldParentID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetParentID(v)
		return nil
	}
	return fmt.Errorf("unknown Node field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *NodeMutation) AddedFields() []string {
	var fields []string
	if m.addvalue != nil {
		fields = append(fields, node.FieldValue)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *NodeMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case node.FieldValue:
		return m.AddedValue()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *NodeMutation) AddField(name string, value ent.Value) error {
	switch name {
	case node.FieldValue:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddValue(v)
		return nil
	}
	return fmt.Errorf("unknown Node numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *NodeMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(node.FieldParentID) {
		fields = append(fields, node.FieldParentID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *NodeMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *NodeMutation) ClearField(name string) error {
	switch name {
	case node.FieldParentID:
		m.ClearParentID()
		return nil
	}
	return fmt.Errorf("unknown Node nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *NodeMutation) ResetField(name string) error {
	switch name {
	case node.FieldValue:
		m.ResetValue()
		return nil
	case node.FieldParentID:
		m.ResetParentID()
		return nil
	}
	return fmt.Errorf("unknown Node field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *NodeMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.parent != nil {
		edges = append(edges, node.EdgeParent)
	}
	if m.children != nil {
		edges = append(edges, node.EdgeChildren)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *NodeMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case node.EdgeParent:
		if id := m.parent; id != nil {
			return []ent.Value{*id}
		}
	case node.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.children))
		for id := range m.children {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *NodeMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedchildren != nil {
		edges = append(edges, node.EdgeChildren)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *NodeMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case node.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.removedchildren))
		for id := range m.removedchildren {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *NodeMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedparent {
		edges = append(edges, node.EdgeParent)
	}
	if m.clearedchildren {
		edges = append(edges, node.EdgeChildren)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *NodeMutation) EdgeCleared(name string) bool {
	switch name {
	case node.EdgeParent:
		return m.clearedparent
	case node.EdgeChildren:
		return m.clearedchildren
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *NodeMutation) ClearEdge(name string) error {
	switch name {
	case node.EdgeParent:
		m.ClearParent()
		return nil
	}
	return fmt.Errorf("unknown Node unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *NodeMutation) ResetEdge(name string) error {
	switch name {
	case node.EdgeParent:
		m.ResetParent()
		return nil
	case node.EdgeChildren:
		m.ResetChildren()
		return nil
	}
	return fmt.Errorf("unknown Node edge %s", name)
}
