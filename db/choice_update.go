// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"test/db/answer"
	"test/db/choice"
	"test/db/predicate"
	"test/db/question"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ChoiceUpdate is the builder for updating Choice entities.
type ChoiceUpdate struct {
	config
	hooks    []Hook
	mutation *ChoiceMutation
}

// Where appends a list predicates to the ChoiceUpdate builder.
func (cu *ChoiceUpdate) Where(ps ...predicate.Choice) *ChoiceUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetLabel sets the "label" field.
func (cu *ChoiceUpdate) SetLabel(s string) *ChoiceUpdate {
	cu.mutation.SetLabel(s)
	return cu
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (cu *ChoiceUpdate) SetNillableLabel(s *string) *ChoiceUpdate {
	if s != nil {
		cu.SetLabel(*s)
	}
	return cu
}

// ClearLabel clears the value of the "label" field.
func (cu *ChoiceUpdate) ClearLabel() *ChoiceUpdate {
	cu.mutation.ClearLabel()
	return cu
}

// SetValue sets the "value" field.
func (cu *ChoiceUpdate) SetValue(s string) *ChoiceUpdate {
	cu.mutation.SetValue(s)
	return cu
}

// SetAnswerID sets the "answer" edge to the Answer entity by ID.
func (cu *ChoiceUpdate) SetAnswerID(id uuid.UUID) *ChoiceUpdate {
	cu.mutation.SetAnswerID(id)
	return cu
}

// SetNillableAnswerID sets the "answer" edge to the Answer entity by ID if the given value is not nil.
func (cu *ChoiceUpdate) SetNillableAnswerID(id *uuid.UUID) *ChoiceUpdate {
	if id != nil {
		cu = cu.SetAnswerID(*id)
	}
	return cu
}

// SetAnswer sets the "answer" edge to the Answer entity.
func (cu *ChoiceUpdate) SetAnswer(a *Answer) *ChoiceUpdate {
	return cu.SetAnswerID(a.ID)
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (cu *ChoiceUpdate) SetQuestionID(id uuid.UUID) *ChoiceUpdate {
	cu.mutation.SetQuestionID(id)
	return cu
}

// SetNillableQuestionID sets the "question" edge to the Question entity by ID if the given value is not nil.
func (cu *ChoiceUpdate) SetNillableQuestionID(id *uuid.UUID) *ChoiceUpdate {
	if id != nil {
		cu = cu.SetQuestionID(*id)
	}
	return cu
}

// SetQuestion sets the "question" edge to the Question entity.
func (cu *ChoiceUpdate) SetQuestion(q *Question) *ChoiceUpdate {
	return cu.SetQuestionID(q.ID)
}

// Mutation returns the ChoiceMutation object of the builder.
func (cu *ChoiceUpdate) Mutation() *ChoiceMutation {
	return cu.mutation
}

// ClearAnswer clears the "answer" edge to the Answer entity.
func (cu *ChoiceUpdate) ClearAnswer() *ChoiceUpdate {
	cu.mutation.ClearAnswer()
	return cu
}

// ClearQuestion clears the "question" edge to the Question entity.
func (cu *ChoiceUpdate) ClearQuestion() *ChoiceUpdate {
	cu.mutation.ClearQuestion()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChoiceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChoiceUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChoiceUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChoiceUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ChoiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(choice.Table, choice.Columns, sqlgraph.NewFieldSpec(choice.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Label(); ok {
		_spec.SetField(choice.FieldLabel, field.TypeString, value)
	}
	if cu.mutation.LabelCleared() {
		_spec.ClearField(choice.FieldLabel, field.TypeString)
	}
	if value, ok := cu.mutation.Value(); ok {
		_spec.SetField(choice.FieldValue, field.TypeString, value)
	}
	if cu.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   choice.AnswerTable,
			Columns: []string{choice.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AnswerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   choice.AnswerTable,
			Columns: []string{choice.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   choice.QuestionTable,
			Columns: []string{choice.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   choice.QuestionTable,
			Columns: []string{choice.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{choice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChoiceUpdateOne is the builder for updating a single Choice entity.
type ChoiceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChoiceMutation
}

// SetLabel sets the "label" field.
func (cuo *ChoiceUpdateOne) SetLabel(s string) *ChoiceUpdateOne {
	cuo.mutation.SetLabel(s)
	return cuo
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (cuo *ChoiceUpdateOne) SetNillableLabel(s *string) *ChoiceUpdateOne {
	if s != nil {
		cuo.SetLabel(*s)
	}
	return cuo
}

// ClearLabel clears the value of the "label" field.
func (cuo *ChoiceUpdateOne) ClearLabel() *ChoiceUpdateOne {
	cuo.mutation.ClearLabel()
	return cuo
}

// SetValue sets the "value" field.
func (cuo *ChoiceUpdateOne) SetValue(s string) *ChoiceUpdateOne {
	cuo.mutation.SetValue(s)
	return cuo
}

// SetAnswerID sets the "answer" edge to the Answer entity by ID.
func (cuo *ChoiceUpdateOne) SetAnswerID(id uuid.UUID) *ChoiceUpdateOne {
	cuo.mutation.SetAnswerID(id)
	return cuo
}

// SetNillableAnswerID sets the "answer" edge to the Answer entity by ID if the given value is not nil.
func (cuo *ChoiceUpdateOne) SetNillableAnswerID(id *uuid.UUID) *ChoiceUpdateOne {
	if id != nil {
		cuo = cuo.SetAnswerID(*id)
	}
	return cuo
}

// SetAnswer sets the "answer" edge to the Answer entity.
func (cuo *ChoiceUpdateOne) SetAnswer(a *Answer) *ChoiceUpdateOne {
	return cuo.SetAnswerID(a.ID)
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (cuo *ChoiceUpdateOne) SetQuestionID(id uuid.UUID) *ChoiceUpdateOne {
	cuo.mutation.SetQuestionID(id)
	return cuo
}

// SetNillableQuestionID sets the "question" edge to the Question entity by ID if the given value is not nil.
func (cuo *ChoiceUpdateOne) SetNillableQuestionID(id *uuid.UUID) *ChoiceUpdateOne {
	if id != nil {
		cuo = cuo.SetQuestionID(*id)
	}
	return cuo
}

// SetQuestion sets the "question" edge to the Question entity.
func (cuo *ChoiceUpdateOne) SetQuestion(q *Question) *ChoiceUpdateOne {
	return cuo.SetQuestionID(q.ID)
}

// Mutation returns the ChoiceMutation object of the builder.
func (cuo *ChoiceUpdateOne) Mutation() *ChoiceMutation {
	return cuo.mutation
}

// ClearAnswer clears the "answer" edge to the Answer entity.
func (cuo *ChoiceUpdateOne) ClearAnswer() *ChoiceUpdateOne {
	cuo.mutation.ClearAnswer()
	return cuo
}

// ClearQuestion clears the "question" edge to the Question entity.
func (cuo *ChoiceUpdateOne) ClearQuestion() *ChoiceUpdateOne {
	cuo.mutation.ClearQuestion()
	return cuo
}

// Where appends a list predicates to the ChoiceUpdate builder.
func (cuo *ChoiceUpdateOne) Where(ps ...predicate.Choice) *ChoiceUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChoiceUpdateOne) Select(field string, fields ...string) *ChoiceUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Choice entity.
func (cuo *ChoiceUpdateOne) Save(ctx context.Context) (*Choice, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChoiceUpdateOne) SaveX(ctx context.Context) *Choice {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChoiceUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChoiceUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ChoiceUpdateOne) sqlSave(ctx context.Context) (_node *Choice, err error) {
	_spec := sqlgraph.NewUpdateSpec(choice.Table, choice.Columns, sqlgraph.NewFieldSpec(choice.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Choice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, choice.FieldID)
		for _, f := range fields {
			if !choice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != choice.FieldID {
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
	if value, ok := cuo.mutation.Label(); ok {
		_spec.SetField(choice.FieldLabel, field.TypeString, value)
	}
	if cuo.mutation.LabelCleared() {
		_spec.ClearField(choice.FieldLabel, field.TypeString)
	}
	if value, ok := cuo.mutation.Value(); ok {
		_spec.SetField(choice.FieldValue, field.TypeString, value)
	}
	if cuo.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   choice.AnswerTable,
			Columns: []string{choice.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AnswerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   choice.AnswerTable,
			Columns: []string{choice.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   choice.QuestionTable,
			Columns: []string{choice.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   choice.QuestionTable,
			Columns: []string{choice.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Choice{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{choice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}