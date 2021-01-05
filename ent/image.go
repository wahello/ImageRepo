// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/JustinHaTran/ImageRepo/ent/image"
	"github.com/JustinHaTran/ImageRepo/ent/user"
	"github.com/facebook/ent/dialect/sql"
)

// Image is the model entity for the Image schema.
type Image struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// RegisteredAt holds the value of the "registered_at" field.
	RegisteredAt time.Time `json:"registered_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImageQuery when eager-loading is set.
	Edges       ImageEdges `json:"edges"`
	user_images *int
}

// ImageEdges holds the relations/edges for other nodes in the graph.
type ImageEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User
	// Imagerepos holds the value of the imagerepos edge.
	Imagerepos []*ImageRepo
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageEdges) OwnerOrErr() (*User, error) {
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

// ImagereposOrErr returns the Imagerepos value or an error if the edge
// was not loaded in eager-loading.
func (e ImageEdges) ImagereposOrErr() ([]*ImageRepo, error) {
	if e.loadedTypes[1] {
		return e.Imagerepos, nil
	}
	return nil, &NotLoadedError{edge: "imagerepos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Image) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // model
		&sql.NullTime{},   // registered_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Image) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_images
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Image fields.
func (i *Image) assignValues(values ...interface{}) error {
	if m, n := len(values), len(image.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	i.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field model", values[0])
	} else if value.Valid {
		i.Model = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field registered_at", values[1])
	} else if value.Valid {
		i.RegisteredAt = value.Time
	}
	values = values[2:]
	if len(values) == len(image.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_images", value)
		} else if value.Valid {
			i.user_images = new(int)
			*i.user_images = int(value.Int64)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the Image.
func (i *Image) QueryOwner() *UserQuery {
	return (&ImageClient{config: i.config}).QueryOwner(i)
}

// QueryImagerepos queries the imagerepos edge of the Image.
func (i *Image) QueryImagerepos() *ImageRepoQuery {
	return (&ImageClient{config: i.config}).QueryImagerepos(i)
}

// Update returns a builder for updating this Image.
// Note that, you need to call Image.Unwrap() before calling this method, if this Image
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Image) Update() *ImageUpdateOne {
	return (&ImageClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (i *Image) Unwrap() *Image {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Image is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Image) String() string {
	var builder strings.Builder
	builder.WriteString("Image(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", model=")
	builder.WriteString(i.Model)
	builder.WriteString(", registered_at=")
	builder.WriteString(i.RegisteredAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Images is a parsable slice of Image.
type Images []*Image

func (i Images) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}