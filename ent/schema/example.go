package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Example holds the schema definition for the Example entity.
type Example struct {
	ent.Schema
}

// Fields of the Example.
func (Example) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("name").MaxLen(64),
		field.String("summary"),
		field.Float("price").Min(0.0),
		field.Time("create_time").Optional().Nillable(),
		field.Time("update_time").Optional().Nillable(),
	}
}

// Edges of the Example.
func (Example) Edges() []ent.Edge {
	return nil
}

func (Example) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "el_example"},
	}
}
