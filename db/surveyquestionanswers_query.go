// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"
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

// SurveyQuestionAnswersQuery is the builder for querying SurveyQuestionAnswers entities.
type SurveyQuestionAnswersQuery struct {
	config
	ctx          *QueryContext
	order        []surveyquestionanswers.OrderOption
	inters       []Interceptor
	predicates   []predicate.SurveyQuestionAnswers
	withSurvey   *SurveyQuery
	withQuestion *QuestionQuery
	withAnswer   *AnswerQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SurveyQuestionAnswersQuery builder.
func (sqaq *SurveyQuestionAnswersQuery) Where(ps ...predicate.SurveyQuestionAnswers) *SurveyQuestionAnswersQuery {
	sqaq.predicates = append(sqaq.predicates, ps...)
	return sqaq
}

// Limit the number of records to be returned by this query.
func (sqaq *SurveyQuestionAnswersQuery) Limit(limit int) *SurveyQuestionAnswersQuery {
	sqaq.ctx.Limit = &limit
	return sqaq
}

// Offset to start from.
func (sqaq *SurveyQuestionAnswersQuery) Offset(offset int) *SurveyQuestionAnswersQuery {
	sqaq.ctx.Offset = &offset
	return sqaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sqaq *SurveyQuestionAnswersQuery) Unique(unique bool) *SurveyQuestionAnswersQuery {
	sqaq.ctx.Unique = &unique
	return sqaq
}

// Order specifies how the records should be ordered.
func (sqaq *SurveyQuestionAnswersQuery) Order(o ...surveyquestionanswers.OrderOption) *SurveyQuestionAnswersQuery {
	sqaq.order = append(sqaq.order, o...)
	return sqaq
}

// QuerySurvey chains the current query on the "survey" edge.
func (sqaq *SurveyQuestionAnswersQuery) QuerySurvey() *SurveyQuery {
	query := (&SurveyClient{config: sqaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sqaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sqaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(surveyquestionanswers.Table, surveyquestionanswers.FieldID, selector),
			sqlgraph.To(survey.Table, survey.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, surveyquestionanswers.SurveyTable, surveyquestionanswers.SurveyColumn),
		)
		fromU = sqlgraph.SetNeighbors(sqaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryQuestion chains the current query on the "question" edge.
func (sqaq *SurveyQuestionAnswersQuery) QueryQuestion() *QuestionQuery {
	query := (&QuestionClient{config: sqaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sqaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sqaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(surveyquestionanswers.Table, surveyquestionanswers.FieldID, selector),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, surveyquestionanswers.QuestionTable, surveyquestionanswers.QuestionColumn),
		)
		fromU = sqlgraph.SetNeighbors(sqaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAnswer chains the current query on the "answer" edge.
func (sqaq *SurveyQuestionAnswersQuery) QueryAnswer() *AnswerQuery {
	query := (&AnswerClient{config: sqaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sqaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sqaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(surveyquestionanswers.Table, surveyquestionanswers.FieldID, selector),
			sqlgraph.To(answer.Table, answer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, surveyquestionanswers.AnswerTable, surveyquestionanswers.AnswerColumn),
		)
		fromU = sqlgraph.SetNeighbors(sqaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SurveyQuestionAnswers entity from the query.
// Returns a *NotFoundError when no SurveyQuestionAnswers was found.
func (sqaq *SurveyQuestionAnswersQuery) First(ctx context.Context) (*SurveyQuestionAnswers, error) {
	nodes, err := sqaq.Limit(1).All(setContextOp(ctx, sqaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{surveyquestionanswers.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sqaq *SurveyQuestionAnswersQuery) FirstX(ctx context.Context) *SurveyQuestionAnswers {
	node, err := sqaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SurveyQuestionAnswers ID from the query.
// Returns a *NotFoundError when no SurveyQuestionAnswers ID was found.
func (sqaq *SurveyQuestionAnswersQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sqaq.Limit(1).IDs(setContextOp(ctx, sqaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{surveyquestionanswers.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sqaq *SurveyQuestionAnswersQuery) FirstIDX(ctx context.Context) int {
	id, err := sqaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SurveyQuestionAnswers entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SurveyQuestionAnswers entity is found.
// Returns a *NotFoundError when no SurveyQuestionAnswers entities are found.
func (sqaq *SurveyQuestionAnswersQuery) Only(ctx context.Context) (*SurveyQuestionAnswers, error) {
	nodes, err := sqaq.Limit(2).All(setContextOp(ctx, sqaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{surveyquestionanswers.Label}
	default:
		return nil, &NotSingularError{surveyquestionanswers.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sqaq *SurveyQuestionAnswersQuery) OnlyX(ctx context.Context) *SurveyQuestionAnswers {
	node, err := sqaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SurveyQuestionAnswers ID in the query.
// Returns a *NotSingularError when more than one SurveyQuestionAnswers ID is found.
// Returns a *NotFoundError when no entities are found.
func (sqaq *SurveyQuestionAnswersQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sqaq.Limit(2).IDs(setContextOp(ctx, sqaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{surveyquestionanswers.Label}
	default:
		err = &NotSingularError{surveyquestionanswers.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sqaq *SurveyQuestionAnswersQuery) OnlyIDX(ctx context.Context) int {
	id, err := sqaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SurveyQuestionAnswersSlice.
func (sqaq *SurveyQuestionAnswersQuery) All(ctx context.Context) ([]*SurveyQuestionAnswers, error) {
	ctx = setContextOp(ctx, sqaq.ctx, "All")
	if err := sqaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SurveyQuestionAnswers, *SurveyQuestionAnswersQuery]()
	return withInterceptors[[]*SurveyQuestionAnswers](ctx, sqaq, qr, sqaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sqaq *SurveyQuestionAnswersQuery) AllX(ctx context.Context) []*SurveyQuestionAnswers {
	nodes, err := sqaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SurveyQuestionAnswers IDs.
func (sqaq *SurveyQuestionAnswersQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sqaq.ctx.Unique == nil && sqaq.path != nil {
		sqaq.Unique(true)
	}
	ctx = setContextOp(ctx, sqaq.ctx, "IDs")
	if err = sqaq.Select(surveyquestionanswers.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sqaq *SurveyQuestionAnswersQuery) IDsX(ctx context.Context) []int {
	ids, err := sqaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sqaq *SurveyQuestionAnswersQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sqaq.ctx, "Count")
	if err := sqaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sqaq, querierCount[*SurveyQuestionAnswersQuery](), sqaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sqaq *SurveyQuestionAnswersQuery) CountX(ctx context.Context) int {
	count, err := sqaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sqaq *SurveyQuestionAnswersQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sqaq.ctx, "Exist")
	switch _, err := sqaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sqaq *SurveyQuestionAnswersQuery) ExistX(ctx context.Context) bool {
	exist, err := sqaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SurveyQuestionAnswersQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sqaq *SurveyQuestionAnswersQuery) Clone() *SurveyQuestionAnswersQuery {
	if sqaq == nil {
		return nil
	}
	return &SurveyQuestionAnswersQuery{
		config:       sqaq.config,
		ctx:          sqaq.ctx.Clone(),
		order:        append([]surveyquestionanswers.OrderOption{}, sqaq.order...),
		inters:       append([]Interceptor{}, sqaq.inters...),
		predicates:   append([]predicate.SurveyQuestionAnswers{}, sqaq.predicates...),
		withSurvey:   sqaq.withSurvey.Clone(),
		withQuestion: sqaq.withQuestion.Clone(),
		withAnswer:   sqaq.withAnswer.Clone(),
		// clone intermediate query.
		sql:  sqaq.sql.Clone(),
		path: sqaq.path,
	}
}

// WithSurvey tells the query-builder to eager-load the nodes that are connected to
// the "survey" edge. The optional arguments are used to configure the query builder of the edge.
func (sqaq *SurveyQuestionAnswersQuery) WithSurvey(opts ...func(*SurveyQuery)) *SurveyQuestionAnswersQuery {
	query := (&SurveyClient{config: sqaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sqaq.withSurvey = query
	return sqaq
}

// WithQuestion tells the query-builder to eager-load the nodes that are connected to
// the "question" edge. The optional arguments are used to configure the query builder of the edge.
func (sqaq *SurveyQuestionAnswersQuery) WithQuestion(opts ...func(*QuestionQuery)) *SurveyQuestionAnswersQuery {
	query := (&QuestionClient{config: sqaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sqaq.withQuestion = query
	return sqaq
}

// WithAnswer tells the query-builder to eager-load the nodes that are connected to
// the "answer" edge. The optional arguments are used to configure the query builder of the edge.
func (sqaq *SurveyQuestionAnswersQuery) WithAnswer(opts ...func(*AnswerQuery)) *SurveyQuestionAnswersQuery {
	query := (&AnswerClient{config: sqaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sqaq.withAnswer = query
	return sqaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		QuestionNum int64 `json:"question_num,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SurveyQuestionAnswers.Query().
//		GroupBy(surveyquestionanswers.FieldQuestionNum).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (sqaq *SurveyQuestionAnswersQuery) GroupBy(field string, fields ...string) *SurveyQuestionAnswersGroupBy {
	sqaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SurveyQuestionAnswersGroupBy{build: sqaq}
	grbuild.flds = &sqaq.ctx.Fields
	grbuild.label = surveyquestionanswers.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		QuestionNum int64 `json:"question_num,omitempty"`
//	}
//
//	client.SurveyQuestionAnswers.Query().
//		Select(surveyquestionanswers.FieldQuestionNum).
//		Scan(ctx, &v)
func (sqaq *SurveyQuestionAnswersQuery) Select(fields ...string) *SurveyQuestionAnswersSelect {
	sqaq.ctx.Fields = append(sqaq.ctx.Fields, fields...)
	sbuild := &SurveyQuestionAnswersSelect{SurveyQuestionAnswersQuery: sqaq}
	sbuild.label = surveyquestionanswers.Label
	sbuild.flds, sbuild.scan = &sqaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SurveyQuestionAnswersSelect configured with the given aggregations.
func (sqaq *SurveyQuestionAnswersQuery) Aggregate(fns ...AggregateFunc) *SurveyQuestionAnswersSelect {
	return sqaq.Select().Aggregate(fns...)
}

func (sqaq *SurveyQuestionAnswersQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sqaq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sqaq); err != nil {
				return err
			}
		}
	}
	for _, f := range sqaq.ctx.Fields {
		if !surveyquestionanswers.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if sqaq.path != nil {
		prev, err := sqaq.path(ctx)
		if err != nil {
			return err
		}
		sqaq.sql = prev
	}
	return nil
}

func (sqaq *SurveyQuestionAnswersQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SurveyQuestionAnswers, error) {
	var (
		nodes       = []*SurveyQuestionAnswers{}
		withFKs     = sqaq.withFKs
		_spec       = sqaq.querySpec()
		loadedTypes = [3]bool{
			sqaq.withSurvey != nil,
			sqaq.withQuestion != nil,
			sqaq.withAnswer != nil,
		}
	)
	if sqaq.withSurvey != nil || sqaq.withQuestion != nil || sqaq.withAnswer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, surveyquestionanswers.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SurveyQuestionAnswers).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SurveyQuestionAnswers{config: sqaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sqaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sqaq.withSurvey; query != nil {
		if err := sqaq.loadSurvey(ctx, query, nodes, nil,
			func(n *SurveyQuestionAnswers, e *Survey) { n.Edges.Survey = e }); err != nil {
			return nil, err
		}
	}
	if query := sqaq.withQuestion; query != nil {
		if err := sqaq.loadQuestion(ctx, query, nodes, nil,
			func(n *SurveyQuestionAnswers, e *Question) { n.Edges.Question = e }); err != nil {
			return nil, err
		}
	}
	if query := sqaq.withAnswer; query != nil {
		if err := sqaq.loadAnswer(ctx, query, nodes, nil,
			func(n *SurveyQuestionAnswers, e *Answer) { n.Edges.Answer = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sqaq *SurveyQuestionAnswersQuery) loadSurvey(ctx context.Context, query *SurveyQuery, nodes []*SurveyQuestionAnswers, init func(*SurveyQuestionAnswers), assign func(*SurveyQuestionAnswers, *Survey)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*SurveyQuestionAnswers)
	for i := range nodes {
		if nodes[i].survey_survey_question_answers == nil {
			continue
		}
		fk := *nodes[i].survey_survey_question_answers
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(survey.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "survey_survey_question_answers" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sqaq *SurveyQuestionAnswersQuery) loadQuestion(ctx context.Context, query *QuestionQuery, nodes []*SurveyQuestionAnswers, init func(*SurveyQuestionAnswers), assign func(*SurveyQuestionAnswers, *Question)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*SurveyQuestionAnswers)
	for i := range nodes {
		if nodes[i].question_survey_question_answers == nil {
			continue
		}
		fk := *nodes[i].question_survey_question_answers
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(question.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "question_survey_question_answers" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sqaq *SurveyQuestionAnswersQuery) loadAnswer(ctx context.Context, query *AnswerQuery, nodes []*SurveyQuestionAnswers, init func(*SurveyQuestionAnswers), assign func(*SurveyQuestionAnswers, *Answer)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*SurveyQuestionAnswers)
	for i := range nodes {
		if nodes[i].answer_survey_question_answers == nil {
			continue
		}
		fk := *nodes[i].answer_survey_question_answers
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(answer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "answer_survey_question_answers" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sqaq *SurveyQuestionAnswersQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sqaq.querySpec()
	_spec.Node.Columns = sqaq.ctx.Fields
	if len(sqaq.ctx.Fields) > 0 {
		_spec.Unique = sqaq.ctx.Unique != nil && *sqaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sqaq.driver, _spec)
}

func (sqaq *SurveyQuestionAnswersQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(surveyquestionanswers.Table, surveyquestionanswers.Columns, sqlgraph.NewFieldSpec(surveyquestionanswers.FieldID, field.TypeInt))
	_spec.From = sqaq.sql
	if unique := sqaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sqaq.path != nil {
		_spec.Unique = true
	}
	if fields := sqaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, surveyquestionanswers.FieldID)
		for i := range fields {
			if fields[i] != surveyquestionanswers.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sqaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sqaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sqaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sqaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sqaq *SurveyQuestionAnswersQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sqaq.driver.Dialect())
	t1 := builder.Table(surveyquestionanswers.Table)
	columns := sqaq.ctx.Fields
	if len(columns) == 0 {
		columns = surveyquestionanswers.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sqaq.sql != nil {
		selector = sqaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sqaq.ctx.Unique != nil && *sqaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sqaq.predicates {
		p(selector)
	}
	for _, p := range sqaq.order {
		p(selector)
	}
	if offset := sqaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sqaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SurveyQuestionAnswersGroupBy is the group-by builder for SurveyQuestionAnswers entities.
type SurveyQuestionAnswersGroupBy struct {
	selector
	build *SurveyQuestionAnswersQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sqagb *SurveyQuestionAnswersGroupBy) Aggregate(fns ...AggregateFunc) *SurveyQuestionAnswersGroupBy {
	sqagb.fns = append(sqagb.fns, fns...)
	return sqagb
}

// Scan applies the selector query and scans the result into the given value.
func (sqagb *SurveyQuestionAnswersGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sqagb.build.ctx, "GroupBy")
	if err := sqagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SurveyQuestionAnswersQuery, *SurveyQuestionAnswersGroupBy](ctx, sqagb.build, sqagb, sqagb.build.inters, v)
}

func (sqagb *SurveyQuestionAnswersGroupBy) sqlScan(ctx context.Context, root *SurveyQuestionAnswersQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sqagb.fns))
	for _, fn := range sqagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sqagb.flds)+len(sqagb.fns))
		for _, f := range *sqagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sqagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sqagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SurveyQuestionAnswersSelect is the builder for selecting fields of SurveyQuestionAnswers entities.
type SurveyQuestionAnswersSelect struct {
	*SurveyQuestionAnswersQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sqas *SurveyQuestionAnswersSelect) Aggregate(fns ...AggregateFunc) *SurveyQuestionAnswersSelect {
	sqas.fns = append(sqas.fns, fns...)
	return sqas
}

// Scan applies the selector query and scans the result into the given value.
func (sqas *SurveyQuestionAnswersSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sqas.ctx, "Select")
	if err := sqas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SurveyQuestionAnswersQuery, *SurveyQuestionAnswersSelect](ctx, sqas.SurveyQuestionAnswersQuery, sqas, sqas.inters, v)
}

func (sqas *SurveyQuestionAnswersSelect) sqlScan(ctx context.Context, root *SurveyQuestionAnswersQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sqas.fns))
	for _, fn := range sqas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sqas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sqas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
