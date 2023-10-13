// Code generated by ent, DO NOT EDIT.

package surveyinvitation

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the surveyinvitation type in the database.
	Label = "survey_invitation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSent holds the string denoting the sent field in the database.
	FieldSent = "sent"
	// FieldSentTimes holds the string denoting the sent_times field in the database.
	FieldSentTimes = "sent_times"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeSurvey holds the string denoting the survey edge name in mutations.
	EdgeSurvey = "survey"
	// CustomerFieldID holds the string denoting the ID field of the Customer.
	CustomerFieldID = "oid"
	// ProjectFieldID holds the string denoting the ID field of the Project.
	ProjectFieldID = "oid"
	// SurveyFieldID holds the string denoting the ID field of the Survey.
	SurveyFieldID = "oid"
	// Table holds the table name of the surveyinvitation in the database.
	Table = "survey_invitations"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "survey_invitations"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_survey_invitations"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "survey_invitations"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_survey_invitations"
	// SurveyTable is the table that holds the survey relation/edge.
	SurveyTable = "survey_invitations"
	// SurveyInverseTable is the table name for the Survey entity.
	// It exists in this package in order to avoid circular dependency with the "survey" package.
	SurveyInverseTable = "surveys"
	// SurveyColumn is the table column denoting the survey relation/edge.
	SurveyColumn = "survey_survey_invitations"
)

// Columns holds all SQL columns for surveyinvitation fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldSent,
	FieldSentTimes,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "survey_invitations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"customer_survey_invitations",
	"project_survey_invitations",
	"survey_survey_invitations",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSent holds the default value on creation for the "sent" field.
	DefaultSent bool
	// DefaultSentTimes holds the default value on creation for the "sent_times" field.
	DefaultSentTimes []string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// StatusUNSEEN is the default value of the Status enum.
const DefaultStatus = StatusUNSEEN

// Status values.
const (
	StatusUNSEEN     Status = "UNSEEN"
	StatusINCOMPLETE Status = "INCOMPLETE"
	StatusCOMPLETED  Status = "COMPLETED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusUNSEEN, StatusINCOMPLETE, StatusCOMPLETED:
		return nil
	default:
		return fmt.Errorf("surveyinvitation: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the SurveyInvitation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// BySent orders the results by the sent field.
func BySent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSent, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCustomerField orders the results by customer field.
func ByCustomerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerStep(), sql.OrderByField(field, opts...))
	}
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}

// BySurveyField orders the results by survey field.
func BySurveyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSurveyStep(), sql.OrderByField(field, opts...))
	}
}
func newCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerInverseTable, CustomerFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
	)
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, ProjectFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
	)
}
func newSurveyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SurveyInverseTable, SurveyFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SurveyTable, SurveyColumn),
	)
}
