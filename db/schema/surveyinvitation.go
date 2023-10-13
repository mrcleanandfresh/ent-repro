package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SurveyInvitation holds the schema definition for the SurveyInvitation entity.
type SurveyInvitation struct {
	ent.Schema
}

// Fields of the SurveyInvitation.
func (SurveyInvitation) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			Values("UNSEEN", "INCOMPLETE", "COMPLETED").
			Default("UNSEEN"),
		field.Bool("sent").
			Default(false),
		field.Strings("sent_times").
			Default([]string{}),
		field.Time("created_at").
			Default(time.Now().UTC),
		field.Time("updated_at").
			Default(time.Now().UTC),
	}
}

// Edges of the SurveyInvitation.
func (SurveyInvitation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("survey_invitations").
			Unique(),
		edge.From("project", Project.Type).
			Ref("survey_invitations").
			Unique(),
		edge.From("survey", Survey.Type).
			Ref("survey_invitations").
			Unique(),
	}
}
