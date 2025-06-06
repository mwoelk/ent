// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/customid/ent/revision"
)

// Revision is the model entity for the Revision schema.
type Revision struct {
	config
	// ID of the ent.
	ID           string `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Revision) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case revision.FieldID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Revision fields.
func (_m *Revision) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case revision.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				_m.ID = value.String
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Revision.
// This includes values selected through modifiers, order, etc.
func (_m *Revision) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// Update returns a builder for updating this Revision.
// Note that you need to call Revision.Unwrap() before calling this method if this Revision
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Revision) Update() *RevisionUpdateOne {
	return NewRevisionClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Revision entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Revision) Unwrap() *Revision {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Revision is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Revision) String() string {
	var builder strings.Builder
	builder.WriteString("Revision(")
	builder.WriteString(fmt.Sprintf("id=%v", _m.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Revisions is a parsable slice of Revision.
type Revisions []*Revision
