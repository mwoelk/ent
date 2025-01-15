// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package entsql

import (
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema"
)

// Annotation is a builtin schema annotation for attaching
// SQL metadata to schema objects for both codegen and runtime.
type Annotation struct {
	// The Schema option allows setting the schema which the table belongs to.
	// Note, this option is no-op for Ent default migration engine. However, schema
	// extensions (like Atlas) can accept this option and implement it accordingly.
	//
	//	entsql.Annotation{
	//		Schema: "public",
	//	}
	//
	Schema string `json:"schema,omitempty"`

	// The Table option allows overriding the default table
	// name that is generated by ent. For example:
	//
	//	entsql.Annotation{
	//		Table: "Users",
	//	}
	//
	Table string `json:"table,omitempty"`

	// Charset defines the character-set of the table. For example:
	//
	//	entsql.Annotation{
	//		Charset: "utf8mb4",
	//	}
	//
	Charset string `json:"charset,omitempty"`

	// Collation defines the collation of the table (a set of rules for comparing
	// characters in a character set). For example:
	//
	//	entsql.Annotation{
	//		Collation: "utf8mb4_bin",
	//	}
	//
	Collation string `json:"collation,omitempty"`

	// Default specifies a literal default value of a column. Note that using
	// this option overrides the default behavior of the code-generation.
	//
	//	entsql.Annotation{
	//		Default: `{"key":"value"}`,
	//	}
	//
	Default string `json:"default,omitempty"`

	// DefaultExpr specifies an expression default value of a column. Using this option,
	// users can define custom expressions to be set as database default values. Note that
	// using this option overrides the default behavior of the code-generation.
	//
	//	entsql.Annotation{
	//		DefaultExpr: "CURRENT_TIMESTAMP",
	//	}
	//
	//	entsql.Annotation{
	//		DefaultExpr: "uuid_generate_v4()",
	//	}
	//
	//	entsql.Annotation{
	//		DefaultExpr: "(a + b)",
	//	}
	//
	DefaultExpr string `json:"default_expr,omitempty"`

	// DefaultExpr specifies an expression default value of a column per dialect.
	// See, DefaultExpr for full doc.
	//
	//	entsql.Annotation{
	//		DefaultExprs: map[string]string{
	//			dialect.MySQL:    "uuid()",
	//			dialect.Postgres: "uuid_generate_v4",
	//		}
	//
	DefaultExprs map[string]string `json:"default_exprs,omitempty"`

	// Options defines the additional table options. For example:
	//
	//	entsql.Annotation{
	//		Options: "ENGINE = INNODB",
	//	}
	//
	Options string `json:"options,omitempty"`

	// Size defines the column size in the generated schema. For example:
	//
	//	entsql.Annotation{
	//		Size: 128,
	//	}
	//
	Size int64 `json:"size,omitempty"`

	// WithComments specifies whether fields' comments should
	// be stored in the database schema as column comments.
	//
	//  withCommentsEnabled := true
	//	entsql.WithComments{
	//		WithComments: &withCommentsEnabled,
	//	}
	//
	WithComments *bool `json:"with_comments,omitempty"`

	// Incremental defines the auto-incremental behavior of a column. For example:
	//
	//  incrementalEnabled := true
	//  entsql.Annotation{
	//      Incremental: &incrementalEnabled,
	//  }
	//
	// By default, this value is nil defaulting to whatever best fits each scenario.
	//
	Incremental *bool `json:"incremental,omitempty"`

	// IncrementStart defines the auto-incremental start value of a column. For example:
	//
	//  incrementStart := 100
	//  entsql.Annotation{
	//      IncrementStart: &incrementStart,
	//  }
	//
	// By default, this value is nil defaulting to whatever the database settings are.
	//
	IncrementStart *int64 `json:"increment_start,omitempty"`

	// OnDelete specifies a custom referential action for DELETE operations on parent
	// table that has matching rows in the child table.
	//
	// For example, in order to delete rows from the parent table and automatically delete
	// their matching rows in the child table, pass the following annotation:
	//
	//	entsql.Annotation{
	//		OnDelete: entsql.Cascade,
	//	}
	//
	OnDelete ReferenceOption `json:"on_delete,omitempty"`

	// Check allows injecting custom "DDL" for setting an unnamed "CHECK" clause in "CREATE TABLE".
	//
	//	entsql.Annotation{
	//		Check: "age < 10",
	//	}
	//
	Check string `json:"check,omitempty"`

	// Checks allows injecting custom "DDL" for setting named "CHECK" clauses in "CREATE TABLE".
	//
	//	entsql.Annotation{
	//		Checks: map[string]string{
	//			"valid_discount": "price > discount_price",
	//		},
	//	}
	//
	Checks map[string]string `json:"checks,omitempty"`

	// Skip indicates that the field or the schema is skipped/ignored during
	// migration (e.g., defined externally).
	//
	//	entsql.Annotation{
	//		Skip: true,
	//	}
	//
	Skip bool `json:"skip,omitempty"`

	// ViewAs allows defining a view for the schema. For example:
	//
	//	entsql.Annotation{
	//		View: "SELECT name FROM users",
	//	}
	ViewAs string `json:"view_as,omitempty"`

	// ViewFor allows defining a view for the schema per dialect. For example:
	//
	//	entsql.Annotation{
	//		ViewFor: map[string]string{
	//			dialect.MySQL:    "...",
	//			dialect.Postgres: "...",
	//		},
	//	}
	ViewFor map[string]string `json:"view_for,omitempty"`

	// error occurs during annotation build. This field is not
	// serialized to JSON and used only by the codegen loader.
	err error
}

// Name describes the annotation name.
func (Annotation) Name() string {
	return "EntSQL"
}

// The Schema option allows setting the schema which the table belongs to.
// Note, this option is no-op for Ent default migration engine. However, schema
// extensions (like Atlas) can accept this option and implement it accordingly.
//
//	func (T) Annotations() []schema.Annotation {
//		return []schema.Annotation{
//			entsql.Schema("public"),
//		}
//	}
func Schema(s string) *Annotation {
	return &Annotation{
		Schema: s,
	}
}

// The Table option allows overriding the default table
// name that is generated by ent. For example:
//
//	func (T) Annotations() []schema.Annotation {
//		return []schema.Annotation{
//			entsql.Table("Users"),
//		}
//	}
func Table(t string) *Annotation {
	return &Annotation{
		Table: t,
	}
}

// SchemaTable allows setting both schema and table name in one annotation.
func SchemaTable(s, t string) *Annotation {
	return &Annotation{
		Schema: s,
		Table:  t,
	}
}

// Check allows injecting custom "DDL" for setting an unnamed "CHECK" clause in "CREATE TABLE".
//
//	entsql.Annotation{
//		Check: "(`age` < 10)",
//	}
func Check(c string) *Annotation {
	return &Annotation{
		Check: c,
	}
}

// Checks allows injecting custom "DDL" for setting named "CHECK" clauses in "CREATE TABLE".
//
//	entsql.Annotation{
//		Checks: map[string]string{
//			"valid_discount": "price > discount_price",
//		},
//	}
func Checks(c map[string]string) *Annotation {
	return &Annotation{
		Checks: c,
	}
}

// Skip indicates that the field or the schema is skipped/ignored during
// migration (e.g., defined externally).
func Skip() *Annotation {
	return &Annotation{Skip: true}
}

// View specifies the definition of a view.
func View(as string) *Annotation {
	return &Annotation{ViewAs: as}
}

// ViewFor specifies the definition of a view.
func ViewFor(dialect string, as func(*sql.Selector)) *Annotation {
	b := sql.Dialect(dialect).Select()
	as(b)
	switch q, args := b.Query(); {
	case len(args) > 0:
		return &Annotation{
			err: fmt.Errorf("entsql: view query should not contain arguments. got: %d", len(args)),
		}
	case q == "":
		return &Annotation{
			err: errors.New("entsql: view query is empty"),
		}
	case b.Err() != nil:
		return &Annotation{
			err: b.Err(),
		}
	default:
		return &Annotation{
			ViewFor: map[string]string{dialect: q},
		}
	}
}

// Default specifies a literal default value of a column. Note that using
// this option overrides the default behavior of the code-generation.
//
//	entsql.Annotation{
//		Default: `{"key":"value"}`,
//	}
func Default(literal string) *Annotation {
	return &Annotation{
		Default: literal,
	}
}

// DefaultExpr specifies an expression default value for the annotated column.
// Using this option, users can define custom expressions to be set as database
// default values.Note that using this option overrides the default behavior of
// the code-generation.
//
//	field.UUID("id", uuid.Nil).
//		Default(uuid.New).
//		Annotations(
//			entsql.DefaultExpr("uuid_generate_v4()"),
//		)
func DefaultExpr(expr string) *Annotation {
	return &Annotation{
		DefaultExpr: expr,
	}
}

// DefaultExprs specifies an expression default value for the annotated
// column per dialect. See, DefaultExpr for full doc.
//
//	field.UUID("id", uuid.Nil).
//		Default(uuid.New).
//		Annotations(
//			entsql.DefaultExprs(map[string]string{
//				dialect.MySQL:    "uuid()",
//				dialect.Postgres: "uuid_generate_v4()",
//			}),
//		)
func DefaultExprs(exprs map[string]string) *Annotation {
	return &Annotation{
		DefaultExprs: exprs,
	}
}

// WithComments specifies whether fields' comments should
// be stored in the database schema as column comments.
//
//	func (T) Annotations() []schema.Annotation {
//		return []schema.Annotation{
//			entsql.WithComments(true),
//		}
//	}
func WithComments(b bool) *Annotation {
	return &Annotation{
		WithComments: &b,
	}
}

// OnDelete specifies a custom referential action for DELETE operations on parent
// table that has matching rows in the child table.
//
// For example, in order to delete rows from the parent table and automatically delete
// their matching rows in the child table, pass the following annotation:
//
//	func (T) Annotations() []schema.Annotation {
//		return []schema.Annotation{
//			entsql.OnDelete(entsql.Cascade),
//		}
//	}
func OnDelete(opt ReferenceOption) *Annotation {
	return &Annotation{
		OnDelete: opt,
	}
}

// IncrementStart specifies the starting value for auto-increment columns.
//
// For example, in order to define the starting value for auto-increment to be 100:
//
//	func (T) Annotations() []schema.Annotation {
//		return []schema.Annotation{
//			entsql.IncrementStart(100),
//		}
//	}
func IncrementStart(i int64) *Annotation {
	return &Annotation{
		IncrementStart: &i,
	}
}

// Merge implements the schema.Merger interface.
func (a Annotation) Merge(other schema.Annotation) schema.Annotation {
	var ant Annotation
	switch other := other.(type) {
	case Annotation:
		ant = other
	case *Annotation:
		if other != nil {
			ant = *other
		}
	default:
		return a
	}
	if s := ant.Schema; s != "" {
		a.Schema = s
	}
	if t := ant.Table; t != "" {
		a.Table = t
	}
	if c := ant.Charset; c != "" {
		a.Charset = c
	}
	if c := ant.Collation; c != "" {
		a.Collation = c
	}
	if d := ant.Default; d != "" {
		a.Default = d
	}
	if d := ant.DefaultExpr; d != "" {
		a.DefaultExpr = d
	}
	if d := ant.DefaultExprs; d != nil {
		if a.DefaultExprs == nil {
			a.DefaultExprs = make(map[string]string)
		}
		for dialect, x := range d {
			a.DefaultExprs[dialect] = x
		}
	}
	if o := ant.Options; o != "" {
		a.Options = o
	}
	if s := ant.Size; s != 0 {
		a.Size = s
	}
	if b := ant.WithComments; b != nil {
		a.WithComments = b
	}
	if i := ant.Incremental; i != nil {
		a.Incremental = i
	}
	if i := ant.IncrementStart; i != nil {
		a.IncrementStart = i
	}
	if od := ant.OnDelete; od != "" {
		a.OnDelete = od
	}
	if c := ant.Check; c != "" {
		a.Check = c
	}
	if checks := ant.Checks; len(checks) > 0 {
		if a.Checks == nil {
			a.Checks = make(map[string]string)
		}
		for name, check := range checks {
			a.Checks[name] = check
		}
	}
	if ant.Skip {
		a.Skip = true
	}
	if v := ant.ViewAs; v != "" {
		a.ViewAs = v
	}
	if vf := ant.ViewFor; len(vf) > 0 {
		if a.ViewFor == nil {
			a.ViewFor = make(map[string]string)
		}
		for dialect, view := range vf {
			a.ViewFor[dialect] = view
		}
	}
	if ant.err != nil {
		a.err = errors.Join(a.err, ant.err)
	}
	return a
}

// Err returns the error that occurred during annotation build, if any.
func (a Annotation) Err() error {
	return a.err
}

var _ interface {
	schema.Annotation
	schema.Merger
} = (*Annotation)(nil)

// ReferenceOption for constraint actions.
type ReferenceOption string

// Reference options (actions) specified by ON UPDATE and ON DELETE
// subclauses of the FOREIGN KEY clause.
const (
	NoAction   ReferenceOption = "NO ACTION"
	Restrict   ReferenceOption = "RESTRICT"
	Cascade    ReferenceOption = "CASCADE"
	SetNull    ReferenceOption = "SET NULL"
	SetDefault ReferenceOption = "SET DEFAULT"
)

// IndexAnnotation is a builtin schema annotation for attaching
// SQL metadata to schema indexes for both codegen and runtime.
type IndexAnnotation struct {
	// Prefix defines a column prefix for a single string column index.
	// In MySQL, the following annotation maps to:
	//
	//	index.Fields("column").
	//		Annotation(entsql.Prefix(100))
	//
	//	CREATE INDEX `table_column` ON `table`(`column`(100))
	//
	Prefix uint

	// PrefixColumns defines column prefixes for a multi-column index.
	// In MySQL, the following annotation maps to:
	//
	//	index.Fields("c1", "c2", "c3").
	//		Annotation(
	//			entsql.PrefixColumn("c1", 100),
	//			entsql.PrefixColumn("c2", 200),
	//		)
	//
	//	CREATE INDEX `table_c1_c2_c3` ON `table`(`c1`(100), `c2`(200), `c3`)
	//
	PrefixColumns map[string]uint

	// Desc defines the DESC clause for a single column index.
	// In MySQL, the following annotation maps to:
	//
	//	index.Fields("column").
	//		Annotation(entsql.Desc())
	//
	//	CREATE INDEX `table_column` ON `table`(`column` DESC)
	//
	Desc bool

	// DescColumns defines the DESC clause for columns in multi-column index.
	// In MySQL, the following annotation maps to:
	//
	//	index.Fields("c1", "c2", "c3").
	//		Annotation(
	//			entsql.DescColumns("c1", "c2"),
	//		)
	//
	//	CREATE INDEX `table_c1_c2_c3` ON `table`(`c1` DESC, `c2` DESC, `c3`)
	//
	DescColumns map[string]bool

	// IncludeColumns defines the INCLUDE clause for the index.
	// Works only in Postgres and its definition is as follows:
	//
	//	index.Fields("c1").
	//		Annotation(
	//			entsql.IncludeColumns("c2"),
	//		)
	//
	//	CREATE INDEX "table_column" ON "table"("c1") INCLUDE ("c2")
	//
	IncludeColumns []string

	// Type defines the type of the index.
	// In MySQL, the following annotation maps to:
	//
	//	index.Fields("c1").
	//		Annotation(
	//			entsql.IndexType("FULLTEXT"),
	//		)
	//
	//	CREATE FULLTEXT INDEX `table_c1` ON `table`(`c1`)
	//
	Type string

	// Types is like the Type option but allows mapping an index-type per dialect.
	//
	//	index.Fields("c1").
	//		Annotation(
	//			entsql.IndexTypes(map[string]string{
	//				dialect.MySQL:		"FULLTEXT",
	//				dialect.Postgres:	"GIN",
	//			}),
	//		)
	//
	Types map[string]string

	// OpClass defines the operator class for a single string column index.
	// In PostgreSQL, the following annotation maps to:
	//
	//	index.Fields("column").
	//		Annotation(
	//			entsql.IndexType("BRIN"),
	//			entsql.OpClass("int8_bloom_ops"),
	//		)
	//
	//	CREATE INDEX "table_column" ON "table" USING BRIN ("column" int8_bloom_ops)
	//
	OpClass string

	// OpClassColumns defines operator-classes for a multi-column index.
	// In PostgreSQL, the following annotation maps to:
	//
	//	index.Fields("c1", "c2", "c3").
	//		Annotation(
	//			entsql.IndexType("BRIN"),
	//			entsql.OpClassColumn("c1", "int8_bloom_ops"),
	//			entsql.OpClassColumn("c2", "int8_minmax_multi_ops(values_per_range=8)"),
	//		)
	//
	//	CREATE INDEX "table_column" ON "table" USING BRIN ("c1" int8_bloom_ops, "c2" int8_minmax_multi_ops(values_per_range=8), "c3")
	//
	OpClassColumns map[string]string

	// IndexWhere allows configuring partial indexes in SQLite and PostgreSQL.
	// Read more: https://postgresql.org/docs/current/indexes-partial.html.
	//
	// Note that the `WHERE` clause should be defined exactly like it is
	// stored in the database (i.e. normal form). Read more about this on
	// the Atlas website: https://atlasgo.io/concepts/dev-database#diffing.
	//
	//	index.Fields("a").
	//		Annotations(
	//			entsql.IndexWhere("b AND c > 0"),
	//		)
	//	CREATE INDEX "table_a" ON "table"("a") WHERE (b AND c > 0)
	Where string
}

// Prefix returns a new index annotation with a single string column index.
// In MySQL, the following annotation maps to:
//
//	index.Fields("column").
//		Annotation(entsql.Prefix(100))
//
//	CREATE INDEX `table_column` ON `table`(`column`(100))
func Prefix(prefix uint) *IndexAnnotation {
	return &IndexAnnotation{
		Prefix: prefix,
	}
}

// PrefixColumn returns a new index annotation with column prefix for
// multi-column indexes. In MySQL, the following annotation maps to:
//
//	index.Fields("c1", "c2", "c3").
//		Annotation(
//			entsql.PrefixColumn("c1", 100),
//			entsql.PrefixColumn("c2", 200),
//		)
//
//	CREATE INDEX `table_c1_c2_c3` ON `table`(`c1`(100), `c2`(200), `c3`)
func PrefixColumn(name string, prefix uint) *IndexAnnotation {
	return &IndexAnnotation{
		PrefixColumns: map[string]uint{
			name: prefix,
		},
	}
}

// OpClass defines the operator class for a single string column index.
// In PostgreSQL, the following annotation maps to:
//
//	index.Fields("column").
//		Annotation(
//			entsql.IndexType("BRIN"),
//			entsql.OpClass("int8_bloom_ops"),
//		)
//
//	CREATE INDEX "table_column" ON "table" USING BRIN ("column" int8_bloom_ops)
func OpClass(op string) *IndexAnnotation {
	return &IndexAnnotation{
		OpClass: op,
	}
}

// OpClassColumn returns a new index annotation with column operator
// class for multi-column indexes. In PostgreSQL, the following annotation maps to:
//
//	index.Fields("c1", "c2", "c3").
//		Annotation(
//			entsql.IndexType("BRIN"),
//			entsql.OpClassColumn("c1", "int8_bloom_ops"),
//			entsql.OpClassColumn("c2", "int8_minmax_multi_ops(values_per_range=8)"),
//		)
//
//	CREATE INDEX "table_column" ON "table" USING BRIN ("c1" int8_bloom_ops, "c2" int8_minmax_multi_ops(values_per_range=8), "c3")
func OpClassColumn(name, op string) *IndexAnnotation {
	return &IndexAnnotation{
		OpClassColumns: map[string]string{
			name: op,
		},
	}
}

// Desc returns a new index annotation with the DESC clause for a
// single column index. In MySQL, the following annotation maps to:
//
//	index.Fields("column").
//		Annotation(entsql.Desc())
//
//	CREATE INDEX `table_column` ON `table`(`column` DESC)
func Desc() *IndexAnnotation {
	return &IndexAnnotation{
		Desc: true,
	}
}

// DescColumns returns a new index annotation with the DESC clause attached to
// the columns in the index. In MySQL, the following annotation maps to:
//
//	index.Fields("c1", "c2", "c3").
//		Annotation(
//			entsql.DescColumns("c1", "c2"),
//		)
//
//	CREATE INDEX `table_c1_c2_c3` ON `table`(`c1` DESC, `c2` DESC, `c3`)
func DescColumns(names ...string) *IndexAnnotation {
	ant := &IndexAnnotation{
		DescColumns: make(map[string]bool, len(names)),
	}
	for i := range names {
		ant.DescColumns[names[i]] = true
	}
	return ant
}

// IncludeColumns defines the INCLUDE clause for the index.
// Works only in Postgres and its definition is as follows:
//
//	index.Fields("c1").
//		Annotation(
//			entsql.IncludeColumns("c2"),
//		)
//
//	CREATE INDEX "table_column" ON "table"("c1") INCLUDE ("c2")
func IncludeColumns(names ...string) *IndexAnnotation {
	return &IndexAnnotation{IncludeColumns: names}
}

// IndexType defines the type of the index.
// In MySQL, the following annotation maps to:
//
//	index.Fields("c1").
//		Annotation(
//			entsql.IndexType("FULLTEXT"),
//		)
//
//	CREATE FULLTEXT INDEX `table_c1` ON `table`(`c1`)
func IndexType(t string) *IndexAnnotation {
	return &IndexAnnotation{Type: t}
}

// IndexTypes is like the Type option but allows mapping an index-type per dialect.
//
//	index.Fields("c1").
//		Annotations(
//			entsql.IndexTypes(map[string]string{
//				dialect.MySQL:    "FULLTEXT",
//				dialect.Postgres: "GIN",
//			}),
//		)
func IndexTypes(types map[string]string) *IndexAnnotation {
	return &IndexAnnotation{Types: types}
}

// IndexWhere allows configuring partial indexes in SQLite and PostgreSQL.
// Read more: https://postgresql.org/docs/current/indexes-partial.html.
//
// Note that the `WHERE` clause should be defined exactly like it is
// stored in the database (i.e. normal form). Read more about this on the
// Atlas website: https://atlasgo.io/concepts/dev-database#diffing.
//
//	index.Fields("a").
//		Annotations(
//			entsql.IndexWhere("b AND c > 0"),
//		)
//	CREATE INDEX "table_a" ON "table"("a") WHERE (b AND c > 0)
func IndexWhere(pred string) *IndexAnnotation {
	return &IndexAnnotation{Where: pred}
}

// Name describes the annotation name.
func (IndexAnnotation) Name() string {
	return "EntSQLIndexes"
}

// Merge implements the schema.Merger interface.
func (a IndexAnnotation) Merge(other schema.Annotation) schema.Annotation {
	var ant IndexAnnotation
	switch other := other.(type) {
	case IndexAnnotation:
		ant = other
	case *IndexAnnotation:
		if other != nil {
			ant = *other
		}
	default:
		return a
	}
	if ant.Prefix != 0 {
		a.Prefix = ant.Prefix
	}
	if ant.PrefixColumns != nil {
		if a.PrefixColumns == nil {
			a.PrefixColumns = make(map[string]uint)
		}
		for column, prefix := range ant.PrefixColumns {
			a.PrefixColumns[column] = prefix
		}
	}
	if ant.OpClass != "" {
		a.OpClass = ant.OpClass
	}
	if ant.OpClassColumns != nil {
		if a.OpClassColumns == nil {
			a.OpClassColumns = make(map[string]string)
		}
		for column, op := range ant.OpClassColumns {
			a.OpClassColumns[column] = op
		}
	}
	if ant.Desc {
		a.Desc = ant.Desc
	}
	if ant.DescColumns != nil {
		if a.DescColumns == nil {
			a.DescColumns = make(map[string]bool)
		}
		for column, desc := range ant.DescColumns {
			a.DescColumns[column] = desc
		}
	}
	if ant.IncludeColumns != nil {
		a.IncludeColumns = append(a.IncludeColumns, ant.IncludeColumns...)
	}
	if ant.Type != "" {
		a.Type = ant.Type
	}
	if ant.Types != nil {
		if a.Types == nil {
			a.Types = make(map[string]string)
		}
		for dialect, t := range ant.Types {
			a.Types[dialect] = t
		}
	}
	if ant.Where != "" {
		a.Where = ant.Where
	}
	return a
}

var _ interface {
	schema.Annotation
	schema.Merger
} = (*IndexAnnotation)(nil)
