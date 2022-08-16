package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Artist holds the schema definition for the Artist entity.
type Artist struct {
	ent.Schema
}

// Fields of the Artist.
func (Artist) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),

		field.String("external_url").
			NotEmpty(),

		field.String("phone_number").
			Unique().
			NotEmpty(),

		field.String("discord").
			Optional(),

		field.String("recommender").
			Optional(),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(entgql.OrderField("CREATED_AT")),
	}
}

// Edges of the Artist.
func (Artist) Edges() []ent.Edge {
	return nil
}

// Annotations of the Artist.
func (Artist) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
