package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	gen "test/db"
	"test/db/hook"
	"time"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.String("prospective_id").
			Unique(),
		field.String("user_id").
			Unique().
			Optional().
			Nillable(),
		field.String("first_name"),
		field.String("last_name").
			Optional().
			Nillable(),
		field.String("email"),
		field.Time("created_at").
			Default(time.Now().UTC).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now().UTC),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("survey_invitations", SurveyInvitation.Type),
		edge.To("projects", Project.Type),
	}
}

func (Customer) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.CustomerFunc(func(ctx context.Context, mutation *gen.CustomerMutation) (gen.Value, error) {
					// set the updated at date
					mutation.SetUpdatedAt(time.Now().UTC())
					return next.Mutate(ctx, mutation)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
