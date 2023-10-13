// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"test/db/customer"
	"test/db/predicate"
	"test/db/project"
	"test/db/surveyinvitation"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CustomerUpdate is the builder for updating Customer entities.
type CustomerUpdate struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// Where appends a list predicates to the CustomerUpdate builder.
func (cu *CustomerUpdate) Where(ps ...predicate.Customer) *CustomerUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetProspectiveID sets the "prospective_id" field.
func (cu *CustomerUpdate) SetProspectiveID(s string) *CustomerUpdate {
	cu.mutation.SetProspectiveID(s)
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CustomerUpdate) SetUserID(s string) *CustomerUpdate {
	cu.mutation.SetUserID(s)
	return cu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableUserID(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetUserID(*s)
	}
	return cu
}

// ClearUserID clears the value of the "user_id" field.
func (cu *CustomerUpdate) ClearUserID() *CustomerUpdate {
	cu.mutation.ClearUserID()
	return cu
}

// SetFirstName sets the "first_name" field.
func (cu *CustomerUpdate) SetFirstName(s string) *CustomerUpdate {
	cu.mutation.SetFirstName(s)
	return cu
}

// SetLastName sets the "last_name" field.
func (cu *CustomerUpdate) SetLastName(s string) *CustomerUpdate {
	cu.mutation.SetLastName(s)
	return cu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableLastName(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetLastName(*s)
	}
	return cu
}

// ClearLastName clears the value of the "last_name" field.
func (cu *CustomerUpdate) ClearLastName() *CustomerUpdate {
	cu.mutation.ClearLastName()
	return cu
}

// SetEmail sets the "email" field.
func (cu *CustomerUpdate) SetEmail(s string) *CustomerUpdate {
	cu.mutation.SetEmail(s)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CustomerUpdate) SetUpdatedAt(t time.Time) *CustomerUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableUpdatedAt(t *time.Time) *CustomerUpdate {
	if t != nil {
		cu.SetUpdatedAt(*t)
	}
	return cu
}

// AddSurveyInvitationIDs adds the "survey_invitations" edge to the SurveyInvitation entity by IDs.
func (cu *CustomerUpdate) AddSurveyInvitationIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddSurveyInvitationIDs(ids...)
	return cu
}

// AddSurveyInvitations adds the "survey_invitations" edges to the SurveyInvitation entity.
func (cu *CustomerUpdate) AddSurveyInvitations(s ...*SurveyInvitation) *CustomerUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddSurveyInvitationIDs(ids...)
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (cu *CustomerUpdate) AddProjectIDs(ids ...uuid.UUID) *CustomerUpdate {
	cu.mutation.AddProjectIDs(ids...)
	return cu
}

// AddProjects adds the "projects" edges to the Project entity.
func (cu *CustomerUpdate) AddProjects(p ...*Project) *CustomerUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddProjectIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cu *CustomerUpdate) Mutation() *CustomerMutation {
	return cu.mutation
}

// ClearSurveyInvitations clears all "survey_invitations" edges to the SurveyInvitation entity.
func (cu *CustomerUpdate) ClearSurveyInvitations() *CustomerUpdate {
	cu.mutation.ClearSurveyInvitations()
	return cu
}

// RemoveSurveyInvitationIDs removes the "survey_invitations" edge to SurveyInvitation entities by IDs.
func (cu *CustomerUpdate) RemoveSurveyInvitationIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemoveSurveyInvitationIDs(ids...)
	return cu
}

// RemoveSurveyInvitations removes "survey_invitations" edges to SurveyInvitation entities.
func (cu *CustomerUpdate) RemoveSurveyInvitations(s ...*SurveyInvitation) *CustomerUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveSurveyInvitationIDs(ids...)
}

// ClearProjects clears all "projects" edges to the Project entity.
func (cu *CustomerUpdate) ClearProjects() *CustomerUpdate {
	cu.mutation.ClearProjects()
	return cu
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (cu *CustomerUpdate) RemoveProjectIDs(ids ...uuid.UUID) *CustomerUpdate {
	cu.mutation.RemoveProjectIDs(ids...)
	return cu
}

// RemoveProjects removes "projects" edges to Project entities.
func (cu *CustomerUpdate) RemoveProjects(p ...*Project) *CustomerUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemoveProjectIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CustomerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CustomerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CustomerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(customer.Table, customer.Columns, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.ProspectiveID(); ok {
		_spec.SetField(customer.FieldProspectiveID, field.TypeString, value)
	}
	if value, ok := cu.mutation.UserID(); ok {
		_spec.SetField(customer.FieldUserID, field.TypeString, value)
	}
	if cu.mutation.UserIDCleared() {
		_spec.ClearField(customer.FieldUserID, field.TypeString)
	}
	if value, ok := cu.mutation.FirstName(); ok {
		_spec.SetField(customer.FieldFirstName, field.TypeString, value)
	}
	if value, ok := cu.mutation.LastName(); ok {
		_spec.SetField(customer.FieldLastName, field.TypeString, value)
	}
	if cu.mutation.LastNameCleared() {
		_spec.ClearField(customer.FieldLastName, field.TypeString)
	}
	if value, ok := cu.mutation.Email(); ok {
		_spec.SetField(customer.FieldEmail, field.TypeString, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(customer.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.SurveyInvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.SurveyInvitationsTable,
			Columns: []string{customer.SurveyInvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyinvitation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedSurveyInvitationsIDs(); len(nodes) > 0 && !cu.mutation.SurveyInvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.SurveyInvitationsTable,
			Columns: []string{customer.SurveyInvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyinvitation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.SurveyInvitationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.SurveyInvitationsTable,
			Columns: []string{customer.SurveyInvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyinvitation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ProjectsTable,
			Columns: []string{customer.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !cu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ProjectsTable,
			Columns: []string{customer.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ProjectsTable,
			Columns: []string{customer.ProjectsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CustomerUpdateOne is the builder for updating a single Customer entity.
type CustomerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CustomerMutation
}

// SetProspectiveID sets the "prospective_id" field.
func (cuo *CustomerUpdateOne) SetProspectiveID(s string) *CustomerUpdateOne {
	cuo.mutation.SetProspectiveID(s)
	return cuo
}

// SetUserID sets the "user_id" field.
func (cuo *CustomerUpdateOne) SetUserID(s string) *CustomerUpdateOne {
	cuo.mutation.SetUserID(s)
	return cuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableUserID(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetUserID(*s)
	}
	return cuo
}

// ClearUserID clears the value of the "user_id" field.
func (cuo *CustomerUpdateOne) ClearUserID() *CustomerUpdateOne {
	cuo.mutation.ClearUserID()
	return cuo
}

// SetFirstName sets the "first_name" field.
func (cuo *CustomerUpdateOne) SetFirstName(s string) *CustomerUpdateOne {
	cuo.mutation.SetFirstName(s)
	return cuo
}

// SetLastName sets the "last_name" field.
func (cuo *CustomerUpdateOne) SetLastName(s string) *CustomerUpdateOne {
	cuo.mutation.SetLastName(s)
	return cuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableLastName(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetLastName(*s)
	}
	return cuo
}

// ClearLastName clears the value of the "last_name" field.
func (cuo *CustomerUpdateOne) ClearLastName() *CustomerUpdateOne {
	cuo.mutation.ClearLastName()
	return cuo
}

// SetEmail sets the "email" field.
func (cuo *CustomerUpdateOne) SetEmail(s string) *CustomerUpdateOne {
	cuo.mutation.SetEmail(s)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CustomerUpdateOne) SetUpdatedAt(t time.Time) *CustomerUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableUpdatedAt(t *time.Time) *CustomerUpdateOne {
	if t != nil {
		cuo.SetUpdatedAt(*t)
	}
	return cuo
}

// AddSurveyInvitationIDs adds the "survey_invitations" edge to the SurveyInvitation entity by IDs.
func (cuo *CustomerUpdateOne) AddSurveyInvitationIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddSurveyInvitationIDs(ids...)
	return cuo
}

// AddSurveyInvitations adds the "survey_invitations" edges to the SurveyInvitation entity.
func (cuo *CustomerUpdateOne) AddSurveyInvitations(s ...*SurveyInvitation) *CustomerUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddSurveyInvitationIDs(ids...)
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (cuo *CustomerUpdateOne) AddProjectIDs(ids ...uuid.UUID) *CustomerUpdateOne {
	cuo.mutation.AddProjectIDs(ids...)
	return cuo
}

// AddProjects adds the "projects" edges to the Project entity.
func (cuo *CustomerUpdateOne) AddProjects(p ...*Project) *CustomerUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddProjectIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cuo *CustomerUpdateOne) Mutation() *CustomerMutation {
	return cuo.mutation
}

// ClearSurveyInvitations clears all "survey_invitations" edges to the SurveyInvitation entity.
func (cuo *CustomerUpdateOne) ClearSurveyInvitations() *CustomerUpdateOne {
	cuo.mutation.ClearSurveyInvitations()
	return cuo
}

// RemoveSurveyInvitationIDs removes the "survey_invitations" edge to SurveyInvitation entities by IDs.
func (cuo *CustomerUpdateOne) RemoveSurveyInvitationIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemoveSurveyInvitationIDs(ids...)
	return cuo
}

// RemoveSurveyInvitations removes "survey_invitations" edges to SurveyInvitation entities.
func (cuo *CustomerUpdateOne) RemoveSurveyInvitations(s ...*SurveyInvitation) *CustomerUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveSurveyInvitationIDs(ids...)
}

// ClearProjects clears all "projects" edges to the Project entity.
func (cuo *CustomerUpdateOne) ClearProjects() *CustomerUpdateOne {
	cuo.mutation.ClearProjects()
	return cuo
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (cuo *CustomerUpdateOne) RemoveProjectIDs(ids ...uuid.UUID) *CustomerUpdateOne {
	cuo.mutation.RemoveProjectIDs(ids...)
	return cuo
}

// RemoveProjects removes "projects" edges to Project entities.
func (cuo *CustomerUpdateOne) RemoveProjects(p ...*Project) *CustomerUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemoveProjectIDs(ids...)
}

// Where appends a list predicates to the CustomerUpdate builder.
func (cuo *CustomerUpdateOne) Where(ps ...predicate.Customer) *CustomerUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CustomerUpdateOne) Select(field string, fields ...string) *CustomerUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Customer entity.
func (cuo *CustomerUpdateOne) Save(ctx context.Context) (*Customer, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CustomerUpdateOne) SaveX(ctx context.Context) *Customer {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CustomerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CustomerUpdateOne) sqlSave(ctx context.Context) (_node *Customer, err error) {
	_spec := sqlgraph.NewUpdateSpec(customer.Table, customer.Columns, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Customer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customer.FieldID)
		for _, f := range fields {
			if !customer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != customer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.ProspectiveID(); ok {
		_spec.SetField(customer.FieldProspectiveID, field.TypeString, value)
	}
	if value, ok := cuo.mutation.UserID(); ok {
		_spec.SetField(customer.FieldUserID, field.TypeString, value)
	}
	if cuo.mutation.UserIDCleared() {
		_spec.ClearField(customer.FieldUserID, field.TypeString)
	}
	if value, ok := cuo.mutation.FirstName(); ok {
		_spec.SetField(customer.FieldFirstName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.LastName(); ok {
		_spec.SetField(customer.FieldLastName, field.TypeString, value)
	}
	if cuo.mutation.LastNameCleared() {
		_spec.ClearField(customer.FieldLastName, field.TypeString)
	}
	if value, ok := cuo.mutation.Email(); ok {
		_spec.SetField(customer.FieldEmail, field.TypeString, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(customer.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.SurveyInvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.SurveyInvitationsTable,
			Columns: []string{customer.SurveyInvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyinvitation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedSurveyInvitationsIDs(); len(nodes) > 0 && !cuo.mutation.SurveyInvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.SurveyInvitationsTable,
			Columns: []string{customer.SurveyInvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyinvitation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.SurveyInvitationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.SurveyInvitationsTable,
			Columns: []string{customer.SurveyInvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyinvitation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ProjectsTable,
			Columns: []string{customer.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !cuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ProjectsTable,
			Columns: []string{customer.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ProjectsTable,
			Columns: []string{customer.ProjectsColumn},
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
	_node = &Customer{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}