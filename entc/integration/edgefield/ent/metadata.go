// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgefield/ent/metadata"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
)

// Metadata is the model entity for the Metadata schema.
type Metadata struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MetadataQuery when eager-loading is set.
	Edges MetadataEdges `json:"edges"`
}

// MetadataEdges holds the relations/edges for other nodes in the graph.
type MetadataEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MetadataEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Metadata) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case metadata.FieldID, metadata.FieldAge:
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Metadata", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Metadata fields.
func (m *Metadata) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case metadata.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case metadata.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				m.Age = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Metadata entity.
func (m *Metadata) QueryUser() *UserQuery {
	return (&MetadataClient{config: m.config}).QueryUser(m)
}

// Update returns a builder for updating this Metadata.
// Note that you need to call Metadata.Unwrap() before calling this method if this Metadata
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Metadata) Update() *MetadataUpdateOne {
	return (&MetadataClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Metadata entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Metadata) Unwrap() *Metadata {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Metadata is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// TableName returns the table name of the Metadata in the database.
func (m *Metadata) TableName() string {
	return metadata.Table
}

// String implements the fmt.Stringer.
func (m *Metadata) String() string {
	var builder strings.Builder
	builder.WriteString("Metadata(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", m.Age))
	builder.WriteByte(')')
	return builder.String()
}

// MetadataSlice is a parsable slice of Metadata.
type MetadataSlice []*Metadata

func (m MetadataSlice) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
