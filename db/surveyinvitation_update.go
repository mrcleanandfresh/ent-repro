// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"test/db/customer"
	"test/db/predicate"
	"test/db/project"
	"test/db/survey"
	"test/db/surveyinvitation"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SurveyInvitationUpdate is the builder for updating SurveyInvitation entities.
type SurveyInvitationUpdate struct {
	config
	hooks    []Hook
	mutation *SurveyInvitationMutation
}

// Where appends a list predicates to the SurveyInvitationUpdate builder.
func (siu *SurveyInvitationUpdate) Where(ps ...predicate.SurveyInvitation) *SurveyInvitationUpdate {
	siu.mutation.Where(ps...)
	return siu
}

// SetStatus sets the "status" field.
func (siu *SurveyInvitationUpdate) SetStatus(s surveyinvitation.Status) *SurveyInvitationUpdate {
	siu.mutation.SetStatus(s)
	return siu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (siu *SurveyInvitationUpdate) SetNillableStatus(s *surveyinvitation.Status) *SurveyInvitationUpdate {
	if s != nil {
		siu.SetStatus(*s)
	}
	return siu
}

// SetSent sets the "sent" field.
func (siu *SurveyInvitationUpdate) SetSent(b bool) *SurveyInvitationUpdate {
	siu.mutation.SetSent(b)
	return siu
}

// SetNillableSent sets the "sent" field if the given value is not nil.
func (siu *SurveyInvitationUpdate) SetNillableSent(b *bool) *SurveyInvitationUpdate {
	if b != nil {
		siu.SetSent(*b)
	}
	return siu
}

// SetSentTimes sets the "sent_times" field.
func (siu *SurveyInvitationUpdate) SetSentTimes(s []string) *SurveyInvitationUpdate {
	siu.mutation.SetSentTimes(s)
	return siu
}

// AppendSentTimes appends s to the "sent_times" field.
func (siu *SurveyInvitationUpdate) AppendSentTimes(s []string) *SurveyInvitationUpdate {
	siu.mutation.AppendSentTimes(s)
	return siu
}

// SetCreatedAt sets the "created_at" field.
func (siu *SurveyInvitationUpdate) SetCreatedAt(t time.Time) *SurveyInvitationUpdate {
	siu.mutation.SetCreatedAt(t)
	return siu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (siu *SurveyInvitationUpdate) SetNillableCreatedAt(t *time.Time) *SurveyInvitationUpdate {
	if t != nil {
		siu.SetCreatedAt(*t)
	}
	return siu
}

// SetUpdatedAt sets the "updated_at" field.
func (siu *SurveyInvitationUpdate) SetUpdatedAt(t time.Time) *SurveyInvitationUpdate {
	siu.mutation.SetUpdatedAt(t)
	return siu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (siu *SurveyInvitationUpdate) SetNillableUpdatedAt(t *time.Time) *SurveyInvitationUpdate {
	if t != nil {
		siu.SetUpdatedAt(*t)
	}
	return siu
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (siu *SurveyInvitationUpdate) SetCustomerID(id uuid.UUID) *SurveyInvitationUpdate {
	siu.mutation.SetCustomerID(id)
	return siu
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (siu *SurveyInvitationUpdate) SetNillableCustomerID(id *uuid.UUID) *SurveyInvitationUpdate {
	if id != nil {
		siu = siu.SetCustomerID(*id)
	}
	return siu
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (siu *SurveyInvitationUpdate) SetCustomer(c *Customer) *SurveyInvitationUpdate {
	return siu.SetCustomerID(c.ID)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (siu *SurveyInvitationUpdate) SetProjectID(id uuid.UUID) *SurveyInvitationUpdate {
	siu.mutation.SetProjectID(id)
	return siu
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (siu *SurveyInvitationUpdate) SetNillableProjectID(id *uuid.UUID) *SurveyInvitationUpdate {
	if id != nil {
		siu = siu.SetProjectID(*id)
	}
	return siu
}

// SetProject sets the "project" edge to the Project entity.
func (siu *SurveyInvitationUpdate) SetProject(p *Project) *SurveyInvitationUpdate {
	return siu.SetProjectID(p.ID)
}

// SetSurveyID sets the "survey" edge to the Survey entity by ID.
func (siu *SurveyInvitationUpdate) SetSurveyID(id uuid.UUID) *SurveyInvitationUpdate {
	siu.mutation.SetSurveyID(id)
	return siu
}

// SetNillableSurveyID sets the "survey" edge to the Survey entity by ID if the given value is not nil.
func (siu *SurveyInvitationUpdate) SetNillableSurveyID(id *uuid.UUID) *SurveyInvitationUpdate {
	if id != nil {
		siu = siu.SetSurveyID(*id)
	}
	return siu
}

// SetSurvey sets the "survey" edge to the Survey entity.
func (siu *SurveyInvitationUpdate) SetSurvey(s *Survey) *SurveyInvitationUpdate {
	return siu.SetSurveyID(s.ID)
}

// Mutation returns the SurveyInvitationMutation object of the builder.
func (siu *SurveyInvitationUpdate) Mutation() *SurveyInvitationMutation {
	return siu.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (siu *SurveyInvitationUpdate) ClearCustomer() *SurveyInvitationUpdate {
	siu.mutation.ClearCustomer()
	return siu
}

// ClearProject clears the "project" edge to the Project entity.
func (siu *SurveyInvitationUpdate) ClearProject() *SurveyInvitationUpdate {
	siu.mutation.ClearProject()
	return siu
}

// ClearSurvey clears the "survey" edge to the Survey entity.
func (siu *SurveyInvitationUpdate) ClearSurvey() *SurveyInvitationUpdate {
	siu.mutation.ClearSurvey()
	return siu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (siu *SurveyInvitationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, siu.sqlSave, siu.mutation, siu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (siu *SurveyInvitationUpdate) SaveX(ctx context.Context) int {
	affected, err := siu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (siu *SurveyInvitationUpdate) Exec(ctx context.Context) error {
	_, err := siu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (siu *SurveyInvitationUpdate) ExecX(ctx context.Context) {
	if err := siu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (siu *SurveyInvitationUpdate) check() error {
	if v, ok := siu.mutation.Status(); ok {
		if err := surveyinvitation.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`db: validator failed for field "SurveyInvitation.status": %w`, err)}
		}
	}
	return nil
}

func (siu *SurveyInvitationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := siu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(surveyinvitation.Table, surveyinvitation.Columns, sqlgraph.NewFieldSpec(surveyinvitation.FieldID, field.TypeInt))
	if ps := siu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := siu.mutation.Status(); ok {
		_spec.SetField(surveyinvitation.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := siu.mutation.Sent(); ok {
		_spec.SetField(surveyinvitation.FieldSent, field.TypeBool, value)
	}
	if value, ok := siu.mutation.SentTimes(); ok {
		_spec.SetField(surveyinvitation.FieldSentTimes, field.TypeJSON, value)
	}
	if value, ok := siu.mutation.AppendedSentTimes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, surveyinvitation.FieldSentTimes, value)
		})
	}
	if value, ok := siu.mutation.CreatedAt(); ok {
		_spec.SetField(surveyinvitation.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := siu.mutation.UpdatedAt(); ok {
		_spec.SetField(surveyinvitation.FieldUpdatedAt, field.TypeTime, value)
	}
	if siu.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyinvitation.CustomerTable,
			Columns: []string{surveyinvitation.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := siu.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyinvitation.CustomerTable,
			Columns: []string{surveyinvitation.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if siu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyinvitation.ProjectTable,
			Columns: []string{surveyinvitation.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := siu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyinvitation.ProjectTable,
			Columns: []string{surveyinvitation.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if siu.mutation.SurveyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyinvitation.SurveyTable,
			Columns: []string{surveyinvitation.SurveyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(survey.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := siu.mutation.SurveyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyinvitation.SurveyTable,
			Columns: []string{surveyinvitation.SurveyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(survey.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, siu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surveyinvitation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	siu.mutation.done = true
	return n, nil
}

// SurveyInvitationUpdateOne is the builder for updating a single SurveyInvitation entity.
type SurveyInvitationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SurveyInvitationMutation
}

// SetStatus sets the "status" field.
func (siuo *SurveyInvitationUpdateOne) SetStatus(s surveyinvitation.Status) *SurveyInvitationUpdateOne {
	siuo.mutation.SetStatus(s)
	return siuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (siuo *SurveyInvitationUpdateOne) SetNillableStatus(s *surveyinvitation.Status) *SurveyInvitationUpdateOne {
	if s != nil {
		siuo.SetStatus(*s)
	}
	return siuo
}

// SetSent sets the "sent" field.
func (siuo *SurveyInvitationUpdateOne) SetSent(b bool) *SurveyInvitationUpdateOne {
	siuo.mutation.SetSent(b)
	return siuo
}

// SetNillableSent sets the "sent" field if the given value is not nil.
func (siuo *SurveyInvitationUpdateOne) SetNillableSent(b *bool) *SurveyInvitationUpdateOne {
	if b != nil {
		siuo.SetSent(*b)
	}
	return siuo
}

// SetSentTimes sets the "sent_times" field.
func (siuo *SurveyInvitationUpdateOne) SetSentTimes(s []string) *SurveyInvitationUpdateOne {
	siuo.mutation.SetSentTimes(s)
	return siuo
}

// AppendSentTimes appends s to the "sent_times" field.
func (siuo *SurveyInvitationUpdateOne) AppendSentTimes(s []string) *SurveyInvitationUpdateOne {
	siuo.mutation.AppendSentTimes(s)
	return siuo
}

// SetCreatedAt sets the "created_at" field.
func (siuo *SurveyInvitationUpdateOne) SetCreatedAt(t time.Time) *SurveyInvitationUpdateOne {
	siuo.mutation.SetCreatedAt(t)
	return siuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (siuo *SurveyInvitationUpdateOne) SetNillableCreatedAt(t *time.Time) *SurveyInvitationUpdateOne {
	if t != nil {
		siuo.SetCreatedAt(*t)
	}
	return siuo
}

// SetUpdatedAt sets the "updated_at" field.
func (siuo *SurveyInvitationUpdateOne) SetUpdatedAt(t time.Time) *SurveyInvitationUpdateOne {
	siuo.mutation.SetUpdatedAt(t)
	return siuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (siuo *SurveyInvitationUpdateOne) SetNillableUpdatedAt(t *time.Time) *SurveyInvitationUpdateOne {
	if t != nil {
		siuo.SetUpdatedAt(*t)
	}
	return siuo
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (siuo *SurveyInvitationUpdateOne) SetCustomerID(id uuid.UUID) *SurveyInvitationUpdateOne {
	siuo.mutation.SetCustomerID(id)
	return siuo
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (siuo *SurveyInvitationUpdateOne) SetNillableCustomerID(id *uuid.UUID) *SurveyInvitationUpdateOne {
	if id != nil {
		siuo = siuo.SetCustomerID(*id)
	}
	return siuo
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (siuo *SurveyInvitationUpdateOne) SetCustomer(c *Customer) *SurveyInvitationUpdateOne {
	return siuo.SetCustomerID(c.ID)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (siuo *SurveyInvitationUpdateOne) SetProjectID(id uuid.UUID) *SurveyInvitationUpdateOne {
	siuo.mutation.SetProjectID(id)
	return siuo
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (siuo *SurveyInvitationUpdateOne) SetNillableProjectID(id *uuid.UUID) *SurveyInvitationUpdateOne {
	if id != nil {
		siuo = siuo.SetProjectID(*id)
	}
	return siuo
}

// SetProject sets the "project" edge to the Project entity.
func (siuo *SurveyInvitationUpdateOne) SetProject(p *Project) *SurveyInvitationUpdateOne {
	return siuo.SetProjectID(p.ID)
}

// SetSurveyID sets the "survey" edge to the Survey entity by ID.
func (siuo *SurveyInvitationUpdateOne) SetSurveyID(id uuid.UUID) *SurveyInvitationUpdateOne {
	siuo.mutation.SetSurveyID(id)
	return siuo
}

// SetNillableSurveyID sets the "survey" edge to the Survey entity by ID if the given value is not nil.
func (siuo *SurveyInvitationUpdateOne) SetNillableSurveyID(id *uuid.UUID) *SurveyInvitationUpdateOne {
	if id != nil {
		siuo = siuo.SetSurveyID(*id)
	}
	return siuo
}

// SetSurvey sets the "survey" edge to the Survey entity.
func (siuo *SurveyInvitationUpdateOne) SetSurvey(s *Survey) *SurveyInvitationUpdateOne {
	return siuo.SetSurveyID(s.ID)
}

// Mutation returns the SurveyInvitationMutation object of the builder.
func (siuo *SurveyInvitationUpdateOne) Mutation() *SurveyInvitationMutation {
	return siuo.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (siuo *SurveyInvitationUpdateOne) ClearCustomer() *SurveyInvitationUpdateOne {
	siuo.mutation.ClearCustomer()
	return siuo
}

// ClearProject clears the "project" edge to the Project entity.
func (siuo *SurveyInvitationUpdateOne) ClearProject() *SurveyInvitationUpdateOne {
	siuo.mutation.ClearProject()
	return siuo
}

// ClearSurvey clears the "survey" edge to the Survey entity.
func (siuo *SurveyInvitationUpdateOne) ClearSurvey() *SurveyInvitationUpdateOne {
	siuo.mutation.ClearSurvey()
	return siuo
}

// Where appends a list predicates to the SurveyInvitationUpdate builder.
func (siuo *SurveyInvitationUpdateOne) Where(ps ...predicate.SurveyInvitation) *SurveyInvitationUpdateOne {
	siuo.mutation.Where(ps...)
	return siuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (siuo *SurveyInvitationUpdateOne) Select(field string, fields ...string) *SurveyInvitationUpdateOne {
	siuo.fields = append([]string{field}, fields...)
	return siuo
}

// Save executes the query and returns the updated SurveyInvitation entity.
func (siuo *SurveyInvitationUpdateOne) Save(ctx context.Context) (*SurveyInvitation, error) {
	return withHooks(ctx, siuo.sqlSave, siuo.mutation, siuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (siuo *SurveyInvitationUpdateOne) SaveX(ctx context.Context) *SurveyInvitation {
	node, err := siuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (siuo *SurveyInvitationUpdateOne) Exec(ctx context.Context) error {
	_, err := siuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (siuo *SurveyInvitationUpdateOne) ExecX(ctx context.Context) {
	if err := siuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (siuo *SurveyInvitationUpdateOne) check() error {
	if v, ok := siuo.mutation.Status(); ok {
		if err := surveyinvitation.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`db: validator failed for field "SurveyInvitation.status": %w`, err)}
		}
	}
	return nil
}

func (siuo *SurveyInvitationUpdateOne) sqlSave(ctx context.Context) (_node *SurveyInvitation, err error) {
	if err := siuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(surveyinvitation.Table, surveyinvitation.Columns, sqlgraph.NewFieldSpec(surveyinvitation.FieldID, field.TypeInt))
	id, ok := siuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "SurveyInvitation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := siuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, surveyinvitation.FieldID)
		for _, f := range fields {
			if !surveyinvitation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != surveyinvitation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := siuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := siuo.mutation.Status(); ok {
		_spec.SetField(surveyinvitation.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := siuo.mutation.Sent(); ok {
		_spec.SetField(surveyinvitation.FieldSent, field.TypeBool, value)
	}
	if value, ok := siuo.mutation.SentTimes(); ok {
		_spec.SetField(surveyinvitation.FieldSentTimes, field.TypeJSON, value)
	}
	if value, ok := siuo.mutation.AppendedSentTimes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, surveyinvitation.FieldSentTimes, value)
		})
	}
	if value, ok := siuo.mutation.CreatedAt(); ok {
		_spec.SetField(surveyinvitation.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := siuo.mutation.UpdatedAt(); ok {
		_spec.SetField(surveyinvitation.FieldUpdatedAt, field.TypeTime, value)
	}
	if siuo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyinvitation.CustomerTable,
			Columns: []string{surveyinvitation.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := siuo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyinvitation.CustomerTable,
			Columns: []string{surveyinvitation.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if siuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyinvitation.ProjectTable,
			Columns: []string{surveyinvitation.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := siuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyinvitation.ProjectTable,
			Columns: []string{surveyinvitation.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if siuo.mutation.SurveyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyinvitation.SurveyTable,
			Columns: []string{surveyinvitation.SurveyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(survey.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := siuo.mutation.SurveyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyinvitation.SurveyTable,
			Columns: []string{surveyinvitation.SurveyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(survey.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SurveyInvitation{config: siuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, siuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surveyinvitation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	siuo.mutation.done = true
	return _node, nil
}