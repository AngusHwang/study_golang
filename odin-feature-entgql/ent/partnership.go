// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"odin/ent/partnership"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Partnership is the model entity for the Partnership schema.
type Partnership struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Company holds the value of the "company" field.
	Company string `json:"company,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Partnership) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case partnership.FieldID:
			values[i] = new(sql.NullInt64)
		case partnership.FieldName, partnership.FieldCompany, partnership.FieldEmail, partnership.FieldContent:
			values[i] = new(sql.NullString)
		case partnership.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Partnership", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Partnership fields.
func (pa *Partnership) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case partnership.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case partnership.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pa.Name = value.String
			}
		case partnership.FieldCompany:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field company", values[i])
			} else if value.Valid {
				pa.Company = value.String
			}
		case partnership.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				pa.Email = value.String
			}
		case partnership.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				pa.Content = value.String
			}
		case partnership.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pa.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Partnership.
// Note that you need to call Partnership.Unwrap() before calling this method if this Partnership
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Partnership) Update() *PartnershipUpdateOne {
	return (&PartnershipClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the Partnership entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Partnership) Unwrap() *Partnership {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Partnership is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Partnership) String() string {
	var builder strings.Builder
	builder.WriteString("Partnership(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("name=")
	builder.WriteString(pa.Name)
	builder.WriteString(", ")
	builder.WriteString("company=")
	builder.WriteString(pa.Company)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(pa.Email)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(pa.Content)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pa.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Partnerships is a parsable slice of Partnership.
type Partnerships []*Partnership

func (pa Partnerships) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}