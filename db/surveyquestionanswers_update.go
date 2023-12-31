// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"test/db/answer"
	"test/db/predicate"
	"test/db/question"
	"test/db/survey"
	"test/db/surveyquestionanswers"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SurveyQuestionAnswersUpdate is the builder for updating SurveyQuestionAnswers entities.
type SurveyQuestionAnswersUpdate struct {
	config
	hooks    []Hook
	mutation *SurveyQuestionAnswersMutation
}

// Where appends a list predicates to the SurveyQuestionAnswersUpdate builder.
func (sqau *SurveyQuestionAnswersUpdate) Where(ps ...predicate.SurveyQuestionAnswers) *SurveyQuestionAnswersUpdate {
	sqau.mutation.Where(ps...)
	return sqau
}

// SetQuestionNum sets the "question_num" field.
func (sqau *SurveyQuestionAnswersUpdate) SetQuestionNum(i int64) *SurveyQuestionAnswersUpdate {
	sqau.mutation.ResetQuestionNum()
	sqau.mutation.SetQuestionNum(i)
	return sqau
}

// SetNillableQuestionNum sets the "question_num" field if the given value is not nil.
func (sqau *SurveyQuestionAnswersUpdate) SetNillableQuestionNum(i *int64) *SurveyQuestionAnswersUpdate {
	if i != nil {
		sqau.SetQuestionNum(*i)
	}
	return sqau
}

// AddQuestionNum adds i to the "question_num" field.
func (sqau *SurveyQuestionAnswersUpdate) AddQuestionNum(i int64) *SurveyQuestionAnswersUpdate {
	sqau.mutation.AddQuestionNum(i)
	return sqau
}

// SetSurveyID sets the "survey" edge to the Survey entity by ID.
func (sqau *SurveyQuestionAnswersUpdate) SetSurveyID(id uuid.UUID) *SurveyQuestionAnswersUpdate {
	sqau.mutation.SetSurveyID(id)
	return sqau
}

// SetSurvey sets the "survey" edge to the Survey entity.
func (sqau *SurveyQuestionAnswersUpdate) SetSurvey(s *Survey) *SurveyQuestionAnswersUpdate {
	return sqau.SetSurveyID(s.ID)
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (sqau *SurveyQuestionAnswersUpdate) SetQuestionID(id uuid.UUID) *SurveyQuestionAnswersUpdate {
	sqau.mutation.SetQuestionID(id)
	return sqau
}

// SetQuestion sets the "question" edge to the Question entity.
func (sqau *SurveyQuestionAnswersUpdate) SetQuestion(q *Question) *SurveyQuestionAnswersUpdate {
	return sqau.SetQuestionID(q.ID)
}

// SetAnswerID sets the "answer" edge to the Answer entity by ID.
func (sqau *SurveyQuestionAnswersUpdate) SetAnswerID(id uuid.UUID) *SurveyQuestionAnswersUpdate {
	sqau.mutation.SetAnswerID(id)
	return sqau
}

// SetNillableAnswerID sets the "answer" edge to the Answer entity by ID if the given value is not nil.
func (sqau *SurveyQuestionAnswersUpdate) SetNillableAnswerID(id *uuid.UUID) *SurveyQuestionAnswersUpdate {
	if id != nil {
		sqau = sqau.SetAnswerID(*id)
	}
	return sqau
}

// SetAnswer sets the "answer" edge to the Answer entity.
func (sqau *SurveyQuestionAnswersUpdate) SetAnswer(a *Answer) *SurveyQuestionAnswersUpdate {
	return sqau.SetAnswerID(a.ID)
}

// Mutation returns the SurveyQuestionAnswersMutation object of the builder.
func (sqau *SurveyQuestionAnswersUpdate) Mutation() *SurveyQuestionAnswersMutation {
	return sqau.mutation
}

// ClearSurvey clears the "survey" edge to the Survey entity.
func (sqau *SurveyQuestionAnswersUpdate) ClearSurvey() *SurveyQuestionAnswersUpdate {
	sqau.mutation.ClearSurvey()
	return sqau
}

// ClearQuestion clears the "question" edge to the Question entity.
func (sqau *SurveyQuestionAnswersUpdate) ClearQuestion() *SurveyQuestionAnswersUpdate {
	sqau.mutation.ClearQuestion()
	return sqau
}

// ClearAnswer clears the "answer" edge to the Answer entity.
func (sqau *SurveyQuestionAnswersUpdate) ClearAnswer() *SurveyQuestionAnswersUpdate {
	sqau.mutation.ClearAnswer()
	return sqau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sqau *SurveyQuestionAnswersUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, sqau.sqlSave, sqau.mutation, sqau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sqau *SurveyQuestionAnswersUpdate) SaveX(ctx context.Context) int {
	affected, err := sqau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sqau *SurveyQuestionAnswersUpdate) Exec(ctx context.Context) error {
	_, err := sqau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sqau *SurveyQuestionAnswersUpdate) ExecX(ctx context.Context) {
	if err := sqau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sqau *SurveyQuestionAnswersUpdate) check() error {
	if _, ok := sqau.mutation.SurveyID(); sqau.mutation.SurveyCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "SurveyQuestionAnswers.survey"`)
	}
	if _, ok := sqau.mutation.QuestionID(); sqau.mutation.QuestionCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "SurveyQuestionAnswers.question"`)
	}
	return nil
}

func (sqau *SurveyQuestionAnswersUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sqau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(surveyquestionanswers.Table, surveyquestionanswers.Columns, sqlgraph.NewFieldSpec(surveyquestionanswers.FieldID, field.TypeInt))
	if ps := sqau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sqau.mutation.QuestionNum(); ok {
		_spec.SetField(surveyquestionanswers.FieldQuestionNum, field.TypeInt64, value)
	}
	if value, ok := sqau.mutation.AddedQuestionNum(); ok {
		_spec.AddField(surveyquestionanswers.FieldQuestionNum, field.TypeInt64, value)
	}
	if sqau.mutation.SurveyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestionanswers.SurveyTable,
			Columns: []string{surveyquestionanswers.SurveyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(survey.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sqau.mutation.SurveyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestionanswers.SurveyTable,
			Columns: []string{surveyquestionanswers.SurveyColumn},
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
	if sqau.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestionanswers.QuestionTable,
			Columns: []string{surveyquestionanswers.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sqau.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestionanswers.QuestionTable,
			Columns: []string{surveyquestionanswers.QuestionColumn},
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
	if sqau.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestionanswers.AnswerTable,
			Columns: []string{surveyquestionanswers.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sqau.mutation.AnswerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestionanswers.AnswerTable,
			Columns: []string{surveyquestionanswers.AnswerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, sqau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surveyquestionanswers.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sqau.mutation.done = true
	return n, nil
}

// SurveyQuestionAnswersUpdateOne is the builder for updating a single SurveyQuestionAnswers entity.
type SurveyQuestionAnswersUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SurveyQuestionAnswersMutation
}

// SetQuestionNum sets the "question_num" field.
func (sqauo *SurveyQuestionAnswersUpdateOne) SetQuestionNum(i int64) *SurveyQuestionAnswersUpdateOne {
	sqauo.mutation.ResetQuestionNum()
	sqauo.mutation.SetQuestionNum(i)
	return sqauo
}

// SetNillableQuestionNum sets the "question_num" field if the given value is not nil.
func (sqauo *SurveyQuestionAnswersUpdateOne) SetNillableQuestionNum(i *int64) *SurveyQuestionAnswersUpdateOne {
	if i != nil {
		sqauo.SetQuestionNum(*i)
	}
	return sqauo
}

// AddQuestionNum adds i to the "question_num" field.
func (sqauo *SurveyQuestionAnswersUpdateOne) AddQuestionNum(i int64) *SurveyQuestionAnswersUpdateOne {
	sqauo.mutation.AddQuestionNum(i)
	return sqauo
}

// SetSurveyID sets the "survey" edge to the Survey entity by ID.
func (sqauo *SurveyQuestionAnswersUpdateOne) SetSurveyID(id uuid.UUID) *SurveyQuestionAnswersUpdateOne {
	sqauo.mutation.SetSurveyID(id)
	return sqauo
}

// SetSurvey sets the "survey" edge to the Survey entity.
func (sqauo *SurveyQuestionAnswersUpdateOne) SetSurvey(s *Survey) *SurveyQuestionAnswersUpdateOne {
	return sqauo.SetSurveyID(s.ID)
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (sqauo *SurveyQuestionAnswersUpdateOne) SetQuestionID(id uuid.UUID) *SurveyQuestionAnswersUpdateOne {
	sqauo.mutation.SetQuestionID(id)
	return sqauo
}

// SetQuestion sets the "question" edge to the Question entity.
func (sqauo *SurveyQuestionAnswersUpdateOne) SetQuestion(q *Question) *SurveyQuestionAnswersUpdateOne {
	return sqauo.SetQuestionID(q.ID)
}

// SetAnswerID sets the "answer" edge to the Answer entity by ID.
func (sqauo *SurveyQuestionAnswersUpdateOne) SetAnswerID(id uuid.UUID) *SurveyQuestionAnswersUpdateOne {
	sqauo.mutation.SetAnswerID(id)
	return sqauo
}

// SetNillableAnswerID sets the "answer" edge to the Answer entity by ID if the given value is not nil.
func (sqauo *SurveyQuestionAnswersUpdateOne) SetNillableAnswerID(id *uuid.UUID) *SurveyQuestionAnswersUpdateOne {
	if id != nil {
		sqauo = sqauo.SetAnswerID(*id)
	}
	return sqauo
}

// SetAnswer sets the "answer" edge to the Answer entity.
func (sqauo *SurveyQuestionAnswersUpdateOne) SetAnswer(a *Answer) *SurveyQuestionAnswersUpdateOne {
	return sqauo.SetAnswerID(a.ID)
}

// Mutation returns the SurveyQuestionAnswersMutation object of the builder.
func (sqauo *SurveyQuestionAnswersUpdateOne) Mutation() *SurveyQuestionAnswersMutation {
	return sqauo.mutation
}

// ClearSurvey clears the "survey" edge to the Survey entity.
func (sqauo *SurveyQuestionAnswersUpdateOne) ClearSurvey() *SurveyQuestionAnswersUpdateOne {
	sqauo.mutation.ClearSurvey()
	return sqauo
}

// ClearQuestion clears the "question" edge to the Question entity.
func (sqauo *SurveyQuestionAnswersUpdateOne) ClearQuestion() *SurveyQuestionAnswersUpdateOne {
	sqauo.mutation.ClearQuestion()
	return sqauo
}

// ClearAnswer clears the "answer" edge to the Answer entity.
func (sqauo *SurveyQuestionAnswersUpdateOne) ClearAnswer() *SurveyQuestionAnswersUpdateOne {
	sqauo.mutation.ClearAnswer()
	return sqauo
}

// Where appends a list predicates to the SurveyQuestionAnswersUpdate builder.
func (sqauo *SurveyQuestionAnswersUpdateOne) Where(ps ...predicate.SurveyQuestionAnswers) *SurveyQuestionAnswersUpdateOne {
	sqauo.mutation.Where(ps...)
	return sqauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sqauo *SurveyQuestionAnswersUpdateOne) Select(field string, fields ...string) *SurveyQuestionAnswersUpdateOne {
	sqauo.fields = append([]string{field}, fields...)
	return sqauo
}

// Save executes the query and returns the updated SurveyQuestionAnswers entity.
func (sqauo *SurveyQuestionAnswersUpdateOne) Save(ctx context.Context) (*SurveyQuestionAnswers, error) {
	return withHooks(ctx, sqauo.sqlSave, sqauo.mutation, sqauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sqauo *SurveyQuestionAnswersUpdateOne) SaveX(ctx context.Context) *SurveyQuestionAnswers {
	node, err := sqauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sqauo *SurveyQuestionAnswersUpdateOne) Exec(ctx context.Context) error {
	_, err := sqauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sqauo *SurveyQuestionAnswersUpdateOne) ExecX(ctx context.Context) {
	if err := sqauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sqauo *SurveyQuestionAnswersUpdateOne) check() error {
	if _, ok := sqauo.mutation.SurveyID(); sqauo.mutation.SurveyCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "SurveyQuestionAnswers.survey"`)
	}
	if _, ok := sqauo.mutation.QuestionID(); sqauo.mutation.QuestionCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "SurveyQuestionAnswers.question"`)
	}
	return nil
}

func (sqauo *SurveyQuestionAnswersUpdateOne) sqlSave(ctx context.Context) (_node *SurveyQuestionAnswers, err error) {
	if err := sqauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(surveyquestionanswers.Table, surveyquestionanswers.Columns, sqlgraph.NewFieldSpec(surveyquestionanswers.FieldID, field.TypeInt))
	id, ok := sqauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "SurveyQuestionAnswers.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sqauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, surveyquestionanswers.FieldID)
		for _, f := range fields {
			if !surveyquestionanswers.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != surveyquestionanswers.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sqauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sqauo.mutation.QuestionNum(); ok {
		_spec.SetField(surveyquestionanswers.FieldQuestionNum, field.TypeInt64, value)
	}
	if value, ok := sqauo.mutation.AddedQuestionNum(); ok {
		_spec.AddField(surveyquestionanswers.FieldQuestionNum, field.TypeInt64, value)
	}
	if sqauo.mutation.SurveyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestionanswers.SurveyTable,
			Columns: []string{surveyquestionanswers.SurveyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(survey.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sqauo.mutation.SurveyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestionanswers.SurveyTable,
			Columns: []string{surveyquestionanswers.SurveyColumn},
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
	if sqauo.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestionanswers.QuestionTable,
			Columns: []string{surveyquestionanswers.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sqauo.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestionanswers.QuestionTable,
			Columns: []string{surveyquestionanswers.QuestionColumn},
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
	if sqauo.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestionanswers.AnswerTable,
			Columns: []string{surveyquestionanswers.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sqauo.mutation.AnswerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyquestionanswers.AnswerTable,
			Columns: []string{surveyquestionanswers.AnswerColumn},
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
	_node = &SurveyQuestionAnswers{config: sqauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sqauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surveyquestionanswers.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sqauo.mutation.done = true
	return _node, nil
}
