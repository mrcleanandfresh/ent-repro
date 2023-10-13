package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SurveyQuestionAnswers holds the schema definition for the SurveyQuestionAnswers entity.
type SurveyQuestionAnswers struct {
	ent.Schema
}

// Fields of the SurveyQuestionAnswers.
func (SurveyQuestionAnswers) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("question_num").
			Default(1),
	}
}

// Edges of the SurveyQuestionAnswers.
func (SurveyQuestionAnswers) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("survey", Survey.Type).
			Ref("survey_question_answers").
			Required().
			Unique(),
		edge.From("question", Question.Type).
			Ref("survey_question_answers").
			Required().
			Unique(),
		edge.From("answer", Answer.Type).
			Ref("survey_question_answers").
			Unique(),
	}
}
