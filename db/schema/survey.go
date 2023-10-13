package schema

import (
	"context"
	"fmt"
	gen "test/db"
	"test/db/hook"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mazen160/go-random"
)

// Survey holds the schema definition for the Survey entity.
type Survey struct {
	ent.Schema
}

// Fields of the Survey.
func (Survey) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.String("identifier").
			Unique().
			DefaultFunc(func() string {
				randomStr, _ := random.String(16)
				return fmt.Sprintf("SUR-%s", randomStr)
			}),
		field.String("name").
			MaxLen(50),
		field.String("description").
			MaxLen(50),
		field.Time("created_at").
			Default(time.Now().UTC).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now().UTC),
	}
}

// Edges of the Survey.
func (Survey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("survey_question_answers", SurveyQuestionAnswers.Type),
		edge.To("survey_invitations", SurveyInvitation.Type),
	}
}

func (Survey) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.SurveyFunc(func(ctx context.Context, mutation *gen.SurveyMutation) (gen.Value, error) {
					// set the updated at date
					mutation.SetUpdatedAt(time.Now().UTC())
					return next.Mutate(ctx, mutation)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
