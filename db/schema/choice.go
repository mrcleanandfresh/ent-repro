package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Choice holds the schema definition for the Choice entity.
type Choice struct {
	ent.Schema
}

// Fields of the Choice.
func (Choice) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").
			Optional().
			Nillable(),
		field.String("value"),
	}
}

// Edges of the Choice.
func (Choice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("answer", Answer.Type).
			Ref("choices").
			Unique(),
		edge.From("question", Question.Type).
			Ref("choices").
			Unique(),
	}
}
