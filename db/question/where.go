// Code generated by ent, DO NOT EDIT.

package question

import (
	"test/db/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldID, id))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldText, v))
}

// SubText applies equality check predicate on the "sub_text" field. It's identical to SubTextEQ.
func SubText(v string) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldSubText, v))
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v int) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldWeight, v))
}

// Required applies equality check predicate on the "required" field. It's identical to RequiredEQ.
func Required(v bool) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldRequired, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldUpdatedAt, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Question {
	return predicate.Question(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Question {
	return predicate.Question(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Question {
	return predicate.Question(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Question {
	return predicate.Question(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Question {
	return predicate.Question(sql.FieldContainsFold(FieldText, v))
}

// SubTextEQ applies the EQ predicate on the "sub_text" field.
func SubTextEQ(v string) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldSubText, v))
}

// SubTextNEQ applies the NEQ predicate on the "sub_text" field.
func SubTextNEQ(v string) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldSubText, v))
}

// SubTextIn applies the In predicate on the "sub_text" field.
func SubTextIn(vs ...string) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldSubText, vs...))
}

// SubTextNotIn applies the NotIn predicate on the "sub_text" field.
func SubTextNotIn(vs ...string) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldSubText, vs...))
}

// SubTextGT applies the GT predicate on the "sub_text" field.
func SubTextGT(v string) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldSubText, v))
}

// SubTextGTE applies the GTE predicate on the "sub_text" field.
func SubTextGTE(v string) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldSubText, v))
}

// SubTextLT applies the LT predicate on the "sub_text" field.
func SubTextLT(v string) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldSubText, v))
}

// SubTextLTE applies the LTE predicate on the "sub_text" field.
func SubTextLTE(v string) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldSubText, v))
}

// SubTextContains applies the Contains predicate on the "sub_text" field.
func SubTextContains(v string) predicate.Question {
	return predicate.Question(sql.FieldContains(FieldSubText, v))
}

// SubTextHasPrefix applies the HasPrefix predicate on the "sub_text" field.
func SubTextHasPrefix(v string) predicate.Question {
	return predicate.Question(sql.FieldHasPrefix(FieldSubText, v))
}

// SubTextHasSuffix applies the HasSuffix predicate on the "sub_text" field.
func SubTextHasSuffix(v string) predicate.Question {
	return predicate.Question(sql.FieldHasSuffix(FieldSubText, v))
}

// SubTextIsNil applies the IsNil predicate on the "sub_text" field.
func SubTextIsNil() predicate.Question {
	return predicate.Question(sql.FieldIsNull(FieldSubText))
}

// SubTextNotNil applies the NotNil predicate on the "sub_text" field.
func SubTextNotNil() predicate.Question {
	return predicate.Question(sql.FieldNotNull(FieldSubText))
}

// SubTextEqualFold applies the EqualFold predicate on the "sub_text" field.
func SubTextEqualFold(v string) predicate.Question {
	return predicate.Question(sql.FieldEqualFold(FieldSubText, v))
}

// SubTextContainsFold applies the ContainsFold predicate on the "sub_text" field.
func SubTextContainsFold(v string) predicate.Question {
	return predicate.Question(sql.FieldContainsFold(FieldSubText, v))
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v int) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldWeight, v))
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v int) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldWeight, v))
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...int) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldWeight, vs...))
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...int) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldWeight, vs...))
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v int) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldWeight, v))
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v int) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldWeight, v))
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v int) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldWeight, v))
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v int) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldWeight, v))
}

// WeightIsNil applies the IsNil predicate on the "weight" field.
func WeightIsNil() predicate.Question {
	return predicate.Question(sql.FieldIsNull(FieldWeight))
}

// WeightNotNil applies the NotNil predicate on the "weight" field.
func WeightNotNil() predicate.Question {
	return predicate.Question(sql.FieldNotNull(FieldWeight))
}

// RequiredEQ applies the EQ predicate on the "required" field.
func RequiredEQ(v bool) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldRequired, v))
}

// RequiredNEQ applies the NEQ predicate on the "required" field.
func RequiredNEQ(v bool) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldRequired, v))
}

// RequiredIsNil applies the IsNil predicate on the "required" field.
func RequiredIsNil() predicate.Question {
	return predicate.Question(sql.FieldIsNull(FieldRequired))
}

// RequiredNotNil applies the NotNil predicate on the "required" field.
func RequiredNotNil() predicate.Question {
	return predicate.Question(sql.FieldNotNull(FieldRequired))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasSurveyQuestionAnswers applies the HasEdge predicate on the "survey_question_answers" edge.
func HasSurveyQuestionAnswers() predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SurveyQuestionAnswersTable, SurveyQuestionAnswersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSurveyQuestionAnswersWith applies the HasEdge predicate on the "survey_question_answers" edge with a given conditions (other predicates).
func HasSurveyQuestionAnswersWith(preds ...predicate.SurveyQuestionAnswers) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := newSurveyQuestionAnswersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasQuestionType applies the HasEdge predicate on the "questionType" edge.
func HasQuestionType() predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, QuestionTypeTable, QuestionTypePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionTypeWith applies the HasEdge predicate on the "questionType" edge with a given conditions (other predicates).
func HasQuestionTypeWith(preds ...predicate.QuestionType) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := newQuestionTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChoices applies the HasEdge predicate on the "choices" edge.
func HasChoices() predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChoicesTable, ChoicesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChoicesWith applies the HasEdge predicate on the "choices" edge with a given conditions (other predicates).
func HasChoicesWith(preds ...predicate.Choice) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := newChoicesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Question) predicate.Question {
	return predicate.Question(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Question) predicate.Question {
	return predicate.Question(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Question) predicate.Question {
	return predicate.Question(sql.NotPredicates(p))
}