// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/entc/integration/gremlin/ent/groupinfo"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Expire holds the value of the "expire" field.
	Expire time.Time `json:"expire,omitempty"`
	// Type holds the value of the "type" field.
	Type *string `json:"type,omitempty"`
	// MaxUsers holds the value of the "max_users" field.
	MaxUsers int `json:"max_users,omitempty"`
	// Name field with multiple validators
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges GroupEdges `json:"edges"`
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Files holds the value of the files edge.
	Files []*File `json:"files,omitempty"`
	// Blocked holds the value of the blocked edge.
	Blocked []*User `json:"blocked,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Info holds the value of the info edge.
	Info *GroupInfo `json:"info,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[0] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// BlockedOrErr returns the Blocked value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) BlockedOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Blocked, nil
	}
	return nil, &NotLoadedError{edge: "blocked"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// InfoOrErr returns the Info value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) InfoOrErr() (*GroupInfo, error) {
	if e.Info != nil {
		return e.Info, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: groupinfo.Label}
	}
	return nil, &NotLoadedError{edge: "info"}
}

// FromResponse scans the gremlin response data into Group.
func (_m *Group) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_m struct {
		ID       string  `json:"id,omitempty"`
		Active   bool    `json:"active,omitempty"`
		Expire   int64   `json:"expire,omitempty"`
		Type     *string `json:"type,omitempty"`
		MaxUsers int     `json:"max_users,omitempty"`
		Name     string  `json:"name,omitempty"`
	}
	if err := vmap.Decode(&scan_m); err != nil {
		return err
	}
	_m.ID = scan_m.ID
	_m.Active = scan_m.Active
	_m.Expire = time.Unix(0, scan_m.Expire)
	_m.Type = scan_m.Type
	_m.MaxUsers = scan_m.MaxUsers
	_m.Name = scan_m.Name
	return nil
}

// QueryFiles queries the "files" edge of the Group entity.
func (_m *Group) QueryFiles() *FileQuery {
	return NewGroupClient(_m.config).QueryFiles(_m)
}

// QueryBlocked queries the "blocked" edge of the Group entity.
func (_m *Group) QueryBlocked() *UserQuery {
	return NewGroupClient(_m.config).QueryBlocked(_m)
}

// QueryUsers queries the "users" edge of the Group entity.
func (_m *Group) QueryUsers() *UserQuery {
	return NewGroupClient(_m.config).QueryUsers(_m)
}

// QueryInfo queries the "info" edge of the Group entity.
func (_m *Group) QueryInfo() *GroupInfoQuery {
	return NewGroupClient(_m.config).QueryInfo(_m)
}

// Update returns a builder for updating this Group.
// Note that you need to call Group.Unwrap() before calling this method if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Group) Update() *GroupUpdateOne {
	return NewGroupClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Group entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Group) Unwrap() *Group {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", _m.Active))
	builder.WriteString(", ")
	builder.WriteString("expire=")
	builder.WriteString(_m.Expire.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := _m.Type; v != nil {
		builder.WriteString("type=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("max_users=")
	builder.WriteString(fmt.Sprintf("%v", _m.MaxUsers))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(_m.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group

// FromResponse scans the gremlin response data into Groups.
func (_m *Groups) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_m []struct {
		ID       string  `json:"id,omitempty"`
		Active   bool    `json:"active,omitempty"`
		Expire   int64   `json:"expire,omitempty"`
		Type     *string `json:"type,omitempty"`
		MaxUsers int     `json:"max_users,omitempty"`
		Name     string  `json:"name,omitempty"`
	}
	if err := vmap.Decode(&scan_m); err != nil {
		return err
	}
	for _, v := range scan_m {
		node := &Group{ID: v.ID}
		node.Active = v.Active
		node.Expire = time.Unix(0, v.Expire)
		node.Type = v.Type
		node.MaxUsers = v.MaxUsers
		node.Name = v.Name
		*_m = append(*_m, node)
	}
	return nil
}
