package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Partnership holds the schema definition for the Partnership entity.
type Partnership struct {
	ent.Schema
}

// Fields of the Partnership.
func (Partnership) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),

		field.String("company").
			NotEmpty(),

		field.String("email").
			NotEmpty(),

		field.String("content").
			Optional().
			Default(""),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Partnership.
func (Partnership) Edges() []ent.Edge {
	return nil
}
