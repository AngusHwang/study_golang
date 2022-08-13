package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name").
			Default("unknown"),
		field.String("nick").
			Default("unknown"),
		field.String("team").
			Default("unknown"),
		field.String("detail").
			Default("unknown"),
		field.String("img").
			Default("unknown"),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}
