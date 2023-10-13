// Code generated by ent, DO NOT EDIT.

package questiontype

import (
	"test/db/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldLTE(FieldID, id))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldEQ(FieldDescription, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldContainsFold(FieldDescription, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.QuestionType {
	return predicate.QuestionType(sql.FieldNotIn(FieldType, vs...))
}

// HasQuestions applies the HasEdge predicate on the "questions" edge.
func HasQuestions() predicate.QuestionType {
	return predicate.QuestionType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, QuestionsTable, QuestionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionsWith applies the HasEdge predicate on the "questions" edge with a given conditions (other predicates).
func HasQuestionsWith(preds ...predicate.Question) predicate.QuestionType {
	return predicate.QuestionType(func(s *sql.Selector) {
		step := newQuestionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.QuestionType) predicate.QuestionType {
	return predicate.QuestionType(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.QuestionType) predicate.QuestionType {
	return predicate.QuestionType(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.QuestionType) predicate.QuestionType {
	return predicate.QuestionType(sql.NotPredicates(p))
}