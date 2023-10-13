// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"test/db/answer"
	"test/db/choice"
	"test/db/surveyquestionanswers"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AnswerCreate is the builder for creating a Answer entity.
type AnswerCreate struct {
	config
	mutation *AnswerMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetText sets the "text" field.
func (ac *AnswerCreate) SetText(s string) *AnswerCreate {
	ac.mutation.SetText(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AnswerCreate) SetCreatedAt(t time.Time) *AnswerCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AnswerCreate) SetNillableCreatedAt(t *time.Time) *AnswerCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AnswerCreate) SetID(u uuid.UUID) *AnswerCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AnswerCreate) SetNillableID(u *uuid.UUID) *AnswerCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// AddChoiceIDs adds the "choices" edge to the Choice entity by IDs.
func (ac *AnswerCreate) AddChoiceIDs(ids ...int) *AnswerCreate {
	ac.mutation.AddChoiceIDs(ids...)
	return ac
}

// AddChoices adds the "choices" edges to the Choice entity.
func (ac *AnswerCreate) AddChoices(c ...*Choice) *AnswerCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ac.AddChoiceIDs(ids...)
}

// AddSurveyQuestionAnswerIDs adds the "survey_question_answers" edge to the SurveyQuestionAnswers entity by IDs.
func (ac *AnswerCreate) AddSurveyQuestionAnswerIDs(ids ...int) *AnswerCreate {
	ac.mutation.AddSurveyQuestionAnswerIDs(ids...)
	return ac
}

// AddSurveyQuestionAnswers adds the "survey_question_answers" edges to the SurveyQuestionAnswers entity.
func (ac *AnswerCreate) AddSurveyQuestionAnswers(s ...*SurveyQuestionAnswers) *AnswerCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ac.AddSurveyQuestionAnswerIDs(ids...)
}

// Mutation returns the AnswerMutation object of the builder.
func (ac *AnswerCreate) Mutation() *AnswerMutation {
	return ac.mutation
}

// Save creates the Answer in the database.
func (ac *AnswerCreate) Save(ctx context.Context) (*Answer, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AnswerCreate) SaveX(ctx context.Context) *Answer {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AnswerCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AnswerCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AnswerCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := answer.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := answer.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AnswerCreate) check() error {
	if _, ok := ac.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`db: missing required field "Answer.text"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "Answer.created_at"`)}
	}
	return nil
}

func (ac *AnswerCreate) sqlSave(ctx context.Context) (*Answer, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AnswerCreate) createSpec() (*Answer, *sqlgraph.CreateSpec) {
	var (
		_node = &Answer{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(answer.Table, sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Text(); ok {
		_spec.SetField(answer.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(answer.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ac.mutation.ChoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   answer.ChoicesTable,
			Columns: []string{answer.ChoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.SurveyQuestionAnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   answer.SurveyQuestionAnswersTable,
			Columns: []string{answer.SurveyQuestionAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyquestionanswers.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Answer.Create().
//		SetText(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AnswerUpsert) {
//			SetText(v+v).
//		}).
//		Exec(ctx)
func (ac *AnswerCreate) OnConflict(opts ...sql.ConflictOption) *AnswerUpsertOne {
	ac.conflict = opts
	return &AnswerUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Answer.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AnswerCreate) OnConflictColumns(columns ...string) *AnswerUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AnswerUpsertOne{
		create: ac,
	}
}

type (
	// AnswerUpsertOne is the builder for "upsert"-ing
	//  one Answer node.
	AnswerUpsertOne struct {
		create *AnswerCreate
	}

	// AnswerUpsert is the "OnConflict" setter.
	AnswerUpsert struct {
		*sql.UpdateSet
	}
)

// SetText sets the "text" field.
func (u *AnswerUpsert) SetText(v string) *AnswerUpsert {
	u.Set(answer.FieldText, v)
	return u
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *AnswerUpsert) UpdateText() *AnswerUpsert {
	u.SetExcluded(answer.FieldText)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Answer.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(answer.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AnswerUpsertOne) UpdateNewValues() *AnswerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(answer.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(answer.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Answer.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AnswerUpsertOne) Ignore() *AnswerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AnswerUpsertOne) DoNothing() *AnswerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AnswerCreate.OnConflict
// documentation for more info.
func (u *AnswerUpsertOne) Update(set func(*AnswerUpsert)) *AnswerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AnswerUpsert{UpdateSet: update})
	}))
	return u
}

// SetText sets the "text" field.
func (u *AnswerUpsertOne) SetText(v string) *AnswerUpsertOne {
	return u.Update(func(s *AnswerUpsert) {
		s.SetText(v)
	})
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *AnswerUpsertOne) UpdateText() *AnswerUpsertOne {
	return u.Update(func(s *AnswerUpsert) {
		s.UpdateText()
	})
}

// Exec executes the query.
func (u *AnswerUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for AnswerCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AnswerUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AnswerUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("db: AnswerUpsertOne.ID is not supported by MySQL driver. Use AnswerUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AnswerUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AnswerCreateBulk is the builder for creating many Answer entities in bulk.
type AnswerCreateBulk struct {
	config
	err      error
	builders []*AnswerCreate
	conflict []sql.ConflictOption
}

// Save creates the Answer entities in the database.
func (acb *AnswerCreateBulk) Save(ctx context.Context) ([]*Answer, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Answer, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AnswerMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AnswerCreateBulk) SaveX(ctx context.Context) []*Answer {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AnswerCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AnswerCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Answer.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AnswerUpsert) {
//			SetText(v+v).
//		}).
//		Exec(ctx)
func (acb *AnswerCreateBulk) OnConflict(opts ...sql.ConflictOption) *AnswerUpsertBulk {
	acb.conflict = opts
	return &AnswerUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Answer.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AnswerCreateBulk) OnConflictColumns(columns ...string) *AnswerUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AnswerUpsertBulk{
		create: acb,
	}
}

// AnswerUpsertBulk is the builder for "upsert"-ing
// a bulk of Answer nodes.
type AnswerUpsertBulk struct {
	create *AnswerCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Answer.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(answer.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AnswerUpsertBulk) UpdateNewValues() *AnswerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(answer.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(answer.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Answer.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AnswerUpsertBulk) Ignore() *AnswerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AnswerUpsertBulk) DoNothing() *AnswerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AnswerCreateBulk.OnConflict
// documentation for more info.
func (u *AnswerUpsertBulk) Update(set func(*AnswerUpsert)) *AnswerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AnswerUpsert{UpdateSet: update})
	}))
	return u
}

// SetText sets the "text" field.
func (u *AnswerUpsertBulk) SetText(v string) *AnswerUpsertBulk {
	return u.Update(func(s *AnswerUpsert) {
		s.SetText(v)
	})
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *AnswerUpsertBulk) UpdateText() *AnswerUpsertBulk {
	return u.Update(func(s *AnswerUpsert) {
		s.UpdateText()
	})
}

// Exec executes the query.
func (u *AnswerUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the AnswerCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for AnswerCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AnswerUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}