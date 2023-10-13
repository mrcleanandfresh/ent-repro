package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// QuestionType holds the schema definition for the QuestionType entity.
type QuestionType struct {
	ent.Schema
}

// Fields of the QuestionType.
func (QuestionType) Fields() []ent.Field {
	return []ent.Field{
		field.String("description"),
		field.Enum("type").
			Values(
				"MULTIPLE_CHOICE",
				"NUMBER",
				"MAP",
				"SHORT",
				"ESSAY",
				"KEY_VALUE",
			),
	}
}

// Edges of the QuestionType.
func (QuestionType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questions", Question.Type),
	}
}
