// Code generated by ent, DO NOT EDIT.

package customer

import (
	"test/db/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldID, id))
}

// ProspectiveID applies equality check predicate on the "prospective_id" field. It's identical to ProspectiveIDEQ.
func ProspectiveID(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldProspectiveID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldUserID, v))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldFirstName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldLastName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldEmail, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldUpdatedAt, v))
}

// ProspectiveIDEQ applies the EQ predicate on the "prospective_id" field.
func ProspectiveIDEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldProspectiveID, v))
}

// ProspectiveIDNEQ applies the NEQ predicate on the "prospective_id" field.
func ProspectiveIDNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldProspectiveID, v))
}

// ProspectiveIDIn applies the In predicate on the "prospective_id" field.
func ProspectiveIDIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldProspectiveID, vs...))
}

// ProspectiveIDNotIn applies the NotIn predicate on the "prospective_id" field.
func ProspectiveIDNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldProspectiveID, vs...))
}

// ProspectiveIDGT applies the GT predicate on the "prospective_id" field.
func ProspectiveIDGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldProspectiveID, v))
}

// ProspectiveIDGTE applies the GTE predicate on the "prospective_id" field.
func ProspectiveIDGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldProspectiveID, v))
}

// ProspectiveIDLT applies the LT predicate on the "prospective_id" field.
func ProspectiveIDLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldProspectiveID, v))
}

// ProspectiveIDLTE applies the LTE predicate on the "prospective_id" field.
func ProspectiveIDLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldProspectiveID, v))
}

// ProspectiveIDContains applies the Contains predicate on the "prospective_id" field.
func ProspectiveIDContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldProspectiveID, v))
}

// ProspectiveIDHasPrefix applies the HasPrefix predicate on the "prospective_id" field.
func ProspectiveIDHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldProspectiveID, v))
}

// ProspectiveIDHasSuffix applies the HasSuffix predicate on the "prospective_id" field.
func ProspectiveIDHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldProspectiveID, v))
}

// ProspectiveIDEqualFold applies the EqualFold predicate on the "prospective_id" field.
func ProspectiveIDEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldProspectiveID, v))
}

// ProspectiveIDContainsFold applies the ContainsFold predicate on the "prospective_id" field.
func ProspectiveIDContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldProspectiveID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Customer {
	return predicate.Customer(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Customer {
	return predicate.Customer(sql.FieldNotNull(FieldUserID))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldUserID, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldFirstName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameIsNil applies the IsNil predicate on the "last_name" field.
func LastNameIsNil() predicate.Customer {
	return predicate.Customer(sql.FieldIsNull(FieldLastName))
}

// LastNameNotNil applies the NotNil predicate on the "last_name" field.
func LastNameNotNil() predicate.Customer {
	return predicate.Customer(sql.FieldNotNull(FieldLastName))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldLastName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldEmail, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasSurveyInvitations applies the HasEdge predicate on the "survey_invitations" edge.
func HasSurveyInvitations() predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SurveyInvitationsTable, SurveyInvitationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSurveyInvitationsWith applies the HasEdge predicate on the "survey_invitations" edge with a given conditions (other predicates).
func HasSurveyInvitationsWith(preds ...predicate.SurveyInvitation) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := newSurveyInvitationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProjects applies the HasEdge predicate on the "projects" edge.
func HasProjects() predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectsTable, ProjectsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectsWith applies the HasEdge predicate on the "projects" edge with a given conditions (other predicates).
func HasProjectsWith(preds ...predicate.Project) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := newProjectsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Customer) predicate.Customer {
	return predicate.Customer(sql.NotPredicates(p))
}
