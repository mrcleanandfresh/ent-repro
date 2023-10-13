// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"test/db/answer"
	"test/db/question"
	"test/db/survey"
	"test/db/surveyquestionanswers"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SurveyQuestionAnswersCreate is the builder for creating a SurveyQuestionAnswers entity.
type SurveyQuestionAnswersCreate struct {
	config
	mutation *SurveyQuestionAnswersMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetQuestionNum sets the "question_num" field.
func (sqac *SurveyQuestionAnswersCreate) SetQuestionNum(i int64) *SurveyQuestionAnswersCreate {
	sqac.mutation.SetQuestionNum(i)
	return sqac
}

// SetNillableQuestionNum sets the "question_num" field if the given value is not nil.
func (sqac *SurveyQuestionAnswersCreate) SetNillableQuestionNum(i *int64) *SurveyQuestionAnswersCreate {
	if i != nil {
		sqac.SetQuestionNum(*i)
	}
	return sqac
}

// SetSurveyID sets the "survey" edge to the Survey entity by ID.
func (sqac *SurveyQuestionAnswersCreate) SetSurveyID(id uuid.UUID) *SurveyQuestionAnswersCreate {
	sqac.mutation.SetSurveyID(id)
	return sqac
}

// SetSurvey sets the "survey" edge to the Survey entity.
func (sqac *SurveyQuestionAnswersCreate) SetSurvey(s *Survey) *SurveyQuestionAnswersCreate {
	return sqac.SetSurveyID(s.ID)
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (sqac *SurveyQuestionAnswersCreate) SetQuestionID(id uuid.UUID) *SurveyQuestionAnswersCreate {
	sqac.mutation.SetQuestionID(id)
	return sqac
}

// SetQuestion sets the "question" edge to the Question entity.
func (sqac *SurveyQuestionAnswersCreate) SetQuestion(q *Question) *SurveyQuestionAnswersCreate {
	return sqac.SetQuestionID(q.ID)
}

// SetAnswerID sets the "answer" edge to the Answer entity by ID.
func (sqac *SurveyQuestionAnswersCreate) SetAnswerID(id uuid.UUID) *SurveyQuestionAnswersCreate {
	sqac.mutation.SetAnswerID(id)
	return sqac
}

// SetNillableAnswerID sets the "answer" edge to the Answer entity by ID if the given value is not nil.
func (sqac *SurveyQuestionAnswersCreate) SetNillableAnswerID(id *uuid.UUID) *SurveyQuestionAnswersCreate {
	if id != nil {
		sqac = sqac.SetAnswerID(*id)
	}
	return sqac
}

// SetAnswer sets the "answer" edge to the Answer entity.
func (sqac *SurveyQuestionAnswersCreate) SetAnswer(a *Answer) *SurveyQuestionAnswersCreate {
	return sqac.SetAnswerID(a.ID)
}

// Mutation returns the SurveyQuestionAnswersMutation object of the builder.
func (sqac *SurveyQuestionAnswersCreate) Mutation() *SurveyQuestionAnswersMutation {
	return sqac.mutation
}

// Save creates the SurveyQuestionAnswers in the database.
func (sqac *SurveyQuestionAnswersCreate) Save(ctx context.Context) (*SurveyQuestionAnswers, error) {
	sqac.defaults()
	return withHooks(ctx, sqac.sqlSave, sqac.mutation, sqac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sqac *SurveyQuestionAnswersCreate) SaveX(ctx context.Context) *SurveyQuestionAnswers {
	v, err := sqac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sqac *SurveyQuestionAnswersCreate) Exec(ctx context.Context) error {
	_, err := sqac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sqac *SurveyQuestionAnswersCreate) ExecX(ctx context.Context) {
	if err := sqac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sqac *SurveyQuestionAnswersCreate) defaults() {
	if _, ok := sqac.mutation.QuestionNum(); !ok {
		v := surveyquestionanswers.DefaultQuestionNum
		sqac.mutation.SetQuestionNum(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sqac *SurveyQuestionAnswersCreate) check() error {
	if _, ok := sqac.mutation.QuestionNum(); !ok {
		return &ValidationError{Name: "question_num", err: errors.New(`db: missing required field "SurveyQuestionAnswers.question_num"`)}
	}
	if _, ok := sqac.mutation.SurveyID(); !ok {
		return &ValidationError{Name: "survey", err: errors.New(`db: missing required edge "SurveyQuestionAnswers.survey"`)}
	}
	if _, ok := sqac.mutation.QuestionID(); !ok {
		return &ValidationError{Name: "question", err: errors.New(`db: missing required edge "SurveyQuestionAnswers.question"`)}
	}
	return nil
}

func (sqac *SurveyQuestionAnswersCreate) sqlSave(ctx context.Context) (*SurveyQuestionAnswers, error) {
	if err := sqac.check(); err != nil {
		return nil, err
	}
	_node, _spec := sqac.createSpec()
	if err := sqlgraph.CreateNode(ctx, sqac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sqac.mutation.id = &_node.ID
	sqac.mutation.done = true
	return _node, nil
}

func (sqac *SurveyQuestionAnswersCreate) createSpec() (*SurveyQuestionAnswers, *sqlgraph.CreateSpec) {
	var (
		_node = &SurveyQuestionAnswers{config: sqac.config}
		_spec = sqlgraph.NewCreateSpec(surveyquestionanswers.Table, sqlgraph.NewFieldSpec(surveyquestionanswers.FieldID, field.TypeInt))
	)
	_spec.OnConflict = sqac.conflict
	if value, ok := sqac.mutation.QuestionNum(); ok {
		_spec.SetField(surveyquestionanswers.FieldQuestionNum, field.TypeInt64, value)
		_node.QuestionNum = value
	}
	if nodes := sqac.mutation.SurveyIDs(); len(nodes) > 0 {
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
		_node.survey_survey_question_answers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sqac.mutation.QuestionIDs(); len(nodes) > 0 {
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
		_node.question_survey_question_answers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sqac.mutation.AnswerIDs(); len(nodes) > 0 {
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
		_node.answer_survey_question_answers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SurveyQuestionAnswers.Create().
//		SetQuestionNum(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SurveyQuestionAnswersUpsert) {
//			SetQuestionNum(v+v).
//		}).
//		Exec(ctx)
func (sqac *SurveyQuestionAnswersCreate) OnConflict(opts ...sql.ConflictOption) *SurveyQuestionAnswersUpsertOne {
	sqac.conflict = opts
	return &SurveyQuestionAnswersUpsertOne{
		create: sqac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SurveyQuestionAnswers.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sqac *SurveyQuestionAnswersCreate) OnConflictColumns(columns ...string) *SurveyQuestionAnswersUpsertOne {
	sqac.conflict = append(sqac.conflict, sql.ConflictColumns(columns...))
	return &SurveyQuestionAnswersUpsertOne{
		create: sqac,
	}
}

type (
	// SurveyQuestionAnswersUpsertOne is the builder for "upsert"-ing
	//  one SurveyQuestionAnswers node.
	SurveyQuestionAnswersUpsertOne struct {
		create *SurveyQuestionAnswersCreate
	}

	// SurveyQuestionAnswersUpsert is the "OnConflict" setter.
	SurveyQuestionAnswersUpsert struct {
		*sql.UpdateSet
	}
)

// SetQuestionNum sets the "question_num" field.
func (u *SurveyQuestionAnswersUpsert) SetQuestionNum(v int64) *SurveyQuestionAnswersUpsert {
	u.Set(surveyquestionanswers.FieldQuestionNum, v)
	return u
}

// UpdateQuestionNum sets the "question_num" field to the value that was provided on create.
func (u *SurveyQuestionAnswersUpsert) UpdateQuestionNum() *SurveyQuestionAnswersUpsert {
	u.SetExcluded(surveyquestionanswers.FieldQuestionNum)
	return u
}

// AddQuestionNum adds v to the "question_num" field.
func (u *SurveyQuestionAnswersUpsert) AddQuestionNum(v int64) *SurveyQuestionAnswersUpsert {
	u.Add(surveyquestionanswers.FieldQuestionNum, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.SurveyQuestionAnswers.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SurveyQuestionAnswersUpsertOne) UpdateNewValues() *SurveyQuestionAnswersUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SurveyQuestionAnswers.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SurveyQuestionAnswersUpsertOne) Ignore() *SurveyQuestionAnswersUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SurveyQuestionAnswersUpsertOne) DoNothing() *SurveyQuestionAnswersUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SurveyQuestionAnswersCreate.OnConflict
// documentation for more info.
func (u *SurveyQuestionAnswersUpsertOne) Update(set func(*SurveyQuestionAnswersUpsert)) *SurveyQuestionAnswersUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SurveyQuestionAnswersUpsert{UpdateSet: update})
	}))
	return u
}

// SetQuestionNum sets the "question_num" field.
func (u *SurveyQuestionAnswersUpsertOne) SetQuestionNum(v int64) *SurveyQuestionAnswersUpsertOne {
	return u.Update(func(s *SurveyQuestionAnswersUpsert) {
		s.SetQuestionNum(v)
	})
}

// AddQuestionNum adds v to the "question_num" field.
func (u *SurveyQuestionAnswersUpsertOne) AddQuestionNum(v int64) *SurveyQuestionAnswersUpsertOne {
	return u.Update(func(s *SurveyQuestionAnswersUpsert) {
		s.AddQuestionNum(v)
	})
}

// UpdateQuestionNum sets the "question_num" field to the value that was provided on create.
func (u *SurveyQuestionAnswersUpsertOne) UpdateQuestionNum() *SurveyQuestionAnswersUpsertOne {
	return u.Update(func(s *SurveyQuestionAnswersUpsert) {
		s.UpdateQuestionNum()
	})
}

// Exec executes the query.
func (u *SurveyQuestionAnswersUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for SurveyQuestionAnswersCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SurveyQuestionAnswersUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SurveyQuestionAnswersUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SurveyQuestionAnswersUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SurveyQuestionAnswersCreateBulk is the builder for creating many SurveyQuestionAnswers entities in bulk.
type SurveyQuestionAnswersCreateBulk struct {
	config
	err      error
	builders []*SurveyQuestionAnswersCreate
	conflict []sql.ConflictOption
}

// Save creates the SurveyQuestionAnswers entities in the database.
func (sqacb *SurveyQuestionAnswersCreateBulk) Save(ctx context.Context) ([]*SurveyQuestionAnswers, error) {
	if sqacb.err != nil {
		return nil, sqacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sqacb.builders))
	nodes := make([]*SurveyQuestionAnswers, len(sqacb.builders))
	mutators := make([]Mutator, len(sqacb.builders))
	for i := range sqacb.builders {
		func(i int, root context.Context) {
			builder := sqacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SurveyQuestionAnswersMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sqacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sqacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sqacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sqacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sqacb *SurveyQuestionAnswersCreateBulk) SaveX(ctx context.Context) []*SurveyQuestionAnswers {
	v, err := sqacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sqacb *SurveyQuestionAnswersCreateBulk) Exec(ctx context.Context) error {
	_, err := sqacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sqacb *SurveyQuestionAnswersCreateBulk) ExecX(ctx context.Context) {
	if err := sqacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SurveyQuestionAnswers.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SurveyQuestionAnswersUpsert) {
//			SetQuestionNum(v+v).
//		}).
//		Exec(ctx)
func (sqacb *SurveyQuestionAnswersCreateBulk) OnConflict(opts ...sql.ConflictOption) *SurveyQuestionAnswersUpsertBulk {
	sqacb.conflict = opts
	return &SurveyQuestionAnswersUpsertBulk{
		create: sqacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SurveyQuestionAnswers.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sqacb *SurveyQuestionAnswersCreateBulk) OnConflictColumns(columns ...string) *SurveyQuestionAnswersUpsertBulk {
	sqacb.conflict = append(sqacb.conflict, sql.ConflictColumns(columns...))
	return &SurveyQuestionAnswersUpsertBulk{
		create: sqacb,
	}
}

// SurveyQuestionAnswersUpsertBulk is the builder for "upsert"-ing
// a bulk of SurveyQuestionAnswers nodes.
type SurveyQuestionAnswersUpsertBulk struct {
	create *SurveyQuestionAnswersCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SurveyQuestionAnswers.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SurveyQuestionAnswersUpsertBulk) UpdateNewValues() *SurveyQuestionAnswersUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SurveyQuestionAnswers.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SurveyQuestionAnswersUpsertBulk) Ignore() *SurveyQuestionAnswersUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SurveyQuestionAnswersUpsertBulk) DoNothing() *SurveyQuestionAnswersUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SurveyQuestionAnswersCreateBulk.OnConflict
// documentation for more info.
func (u *SurveyQuestionAnswersUpsertBulk) Update(set func(*SurveyQuestionAnswersUpsert)) *SurveyQuestionAnswersUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SurveyQuestionAnswersUpsert{UpdateSet: update})
	}))
	return u
}

// SetQuestionNum sets the "question_num" field.
func (u *SurveyQuestionAnswersUpsertBulk) SetQuestionNum(v int64) *SurveyQuestionAnswersUpsertBulk {
	return u.Update(func(s *SurveyQuestionAnswersUpsert) {
		s.SetQuestionNum(v)
	})
}

// AddQuestionNum adds v to the "question_num" field.
func (u *SurveyQuestionAnswersUpsertBulk) AddQuestionNum(v int64) *SurveyQuestionAnswersUpsertBulk {
	return u.Update(func(s *SurveyQuestionAnswersUpsert) {
		s.AddQuestionNum(v)
	})
}

// UpdateQuestionNum sets the "question_num" field to the value that was provided on create.
func (u *SurveyQuestionAnswersUpsertBulk) UpdateQuestionNum() *SurveyQuestionAnswersUpsertBulk {
	return u.Update(func(s *SurveyQuestionAnswersUpsert) {
		s.UpdateQuestionNum()
	})
}

// Exec executes the query.
func (u *SurveyQuestionAnswersUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the SurveyQuestionAnswersCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for SurveyQuestionAnswersCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SurveyQuestionAnswersUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
