package schema

import (
	"context"
	gen "test/db"
	"test/db/hook"
	"test/db/schema/schematype"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

// Fields of the Question.
func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.String("text").
			Unique(),
		field.String("sub_text").
			Optional().
			Nillable(),
		field.Int("weight").
			Optional().
			Nillable().
			Default(0),
		field.Bool("required").
			Optional().
			Nillable().
			Default(true),
		field.JSON("meta", &schematype.Metadata{}).
			Default(&schematype.Metadata{}),
		field.Time("created_at").
			Default(time.Now().UTC).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now().UTC),
	}
}

// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("survey_question_answers", SurveyQuestionAnswers.Type),
		edge.From("questionType", QuestionType.Type).
			Ref("questions"),
		edge.To("choices", Choice.Type),
	}
}

func (Question) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.QuestionFunc(func(ctx context.Context, mutation *gen.QuestionMutation) (gen.Value, error) {
					// set the updated at date
					mutation.SetUpdatedAt(time.Now().UTC())
					return next.Mutate(ctx, mutation)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
