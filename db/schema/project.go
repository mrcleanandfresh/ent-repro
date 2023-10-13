package schema

import (
	"context"
	"errors"
	gen "test/db"
	"test/db/hook"
	"test/db/schema/schematype"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.String("prospective_id").
			Unique(),
		field.String("existing_id").
			Unique().
			Optional().
			Nillable(),
		field.String("name").
			MaxLen(50),
		field.String("description").
			MaxLen(255),
		field.Float("est_budget").
			GoType(schematype.Amount(0)).
			Validate(func(f float64) error {
				if f < 0 {
					return errors.New("cannot have a negative estimate")
				}
				return nil
			}),
		field.Time("created_at").
			Default(time.Now().UTC).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now().UTC),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("projects").
			Unique(),
		edge.To("survey_invitations", SurveyInvitation.Type),
	}
}

func (Project) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.ProjectFunc(func(ctx context.Context, mutation *gen.ProjectMutation) (gen.Value, error) {
					// set the updated at date
					mutation.SetUpdatedAt(time.Now().UTC())
					return next.Mutate(ctx, mutation)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
