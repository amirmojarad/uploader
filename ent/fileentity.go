// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"uploader/ent/fileentity"
	"uploader/ent/user"

	"entgo.io/ent/dialect/sql"
)

// FileEntity is the model entity for the FileEntity schema.
type FileEntity struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Size holds the value of the "size" field.
	Size int64 `json:"size,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileEntityQuery when eager-loading is set.
	Edges      FileEntityEdges `json:"edges"`
	user_files *int
}

// FileEntityEdges holds the relations/edges for other nodes in the graph.
type FileEntityEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEntityEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FileEntity) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case fileentity.FieldID, fileentity.FieldSize:
			values[i] = new(sql.NullInt64)
		case fileentity.FieldName, fileentity.FieldType, fileentity.FieldURL:
			values[i] = new(sql.NullString)
		case fileentity.ForeignKeys[0]: // user_files
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FileEntity", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FileEntity fields.
func (fe *FileEntity) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fileentity.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fe.ID = int(value.Int64)
		case fileentity.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				fe.Name = value.String
			}
		case fileentity.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				fe.Type = value.String
			}
		case fileentity.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				fe.Size = value.Int64
			}
		case fileentity.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				fe.URL = value.String
			}
		case fileentity.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_files", value)
			} else if value.Valid {
				fe.user_files = new(int)
				*fe.user_files = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the FileEntity entity.
func (fe *FileEntity) QueryOwner() *UserQuery {
	return (&FileEntityClient{config: fe.config}).QueryOwner(fe)
}

// Update returns a builder for updating this FileEntity.
// Note that you need to call FileEntity.Unwrap() before calling this method if this FileEntity
// was returned from a transaction, and the transaction was committed or rolled back.
func (fe *FileEntity) Update() *FileEntityUpdateOne {
	return (&FileEntityClient{config: fe.config}).UpdateOne(fe)
}

// Unwrap unwraps the FileEntity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fe *FileEntity) Unwrap() *FileEntity {
	_tx, ok := fe.config.driver.(*txDriver)
	if !ok {
		panic("ent: FileEntity is not a transactional entity")
	}
	fe.config.driver = _tx.drv
	return fe
}

// String implements the fmt.Stringer.
func (fe *FileEntity) String() string {
	var builder strings.Builder
	builder.WriteString("FileEntity(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fe.ID))
	builder.WriteString("name=")
	builder.WriteString(fe.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fe.Type)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", fe.Size))
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(fe.URL)
	builder.WriteByte(')')
	return builder.String()
}

// FileEntities is a parsable slice of FileEntity.
type FileEntities []*FileEntity

func (fe FileEntities) config(cfg config) {
	for _i := range fe {
		fe[_i].config = cfg
	}
}
