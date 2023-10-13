// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"test/db/answer"
	"test/db/choice"
	"test/db/predicate"
	"test/db/surveyquestionanswers"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AnswerQuery is the builder for querying Answer entities.
type AnswerQuery struct {
	config
	ctx                       *QueryContext
	order                     []answer.OrderOption
	inters                    []Interceptor
	predicates                []predicate.Answer
	withChoices               *ChoiceQuery
	withSurveyQuestionAnswers *SurveyQuestionAnswersQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AnswerQuery builder.
func (aq *AnswerQuery) Where(ps ...predicate.Answer) *AnswerQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AnswerQuery) Limit(limit int) *AnswerQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AnswerQuery) Offset(offset int) *AnswerQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AnswerQuery) Unique(unique bool) *AnswerQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AnswerQuery) Order(o ...answer.OrderOption) *AnswerQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryChoices chains the current query on the "choices" edge.
func (aq *AnswerQuery) QueryChoices() *ChoiceQuery {
	query := (&ChoiceClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(answer.Table, answer.FieldID, selector),
			sqlgraph.To(choice.Table, choice.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, answer.ChoicesTable, answer.ChoicesColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySurveyQuestionAnswers chains the current query on the "survey_question_answers" edge.
func (aq *AnswerQuery) QuerySurveyQuestionAnswers() *SurveyQuestionAnswersQuery {
	query := (&SurveyQuestionAnswersClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(answer.Table, answer.FieldID, selector),
			sqlgraph.To(surveyquestionanswers.Table, surveyquestionanswers.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, answer.SurveyQuestionAnswersTable, answer.SurveyQuestionAnswersColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Answer entity from the query.
// Returns a *NotFoundError when no Answer was found.
func (aq *AnswerQuery) First(ctx context.Context) (*Answer, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{answer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AnswerQuery) FirstX(ctx context.Context) *Answer {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Answer ID from the query.
// Returns a *NotFoundError when no Answer ID was found.
func (aq *AnswerQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{answer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AnswerQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Answer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Answer entity is found.
// Returns a *NotFoundError when no Answer entities are found.
func (aq *AnswerQuery) Only(ctx context.Context) (*Answer, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{answer.Label}
	default:
		return nil, &NotSingularError{answer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AnswerQuery) OnlyX(ctx context.Context) *Answer {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Answer ID in the query.
// Returns a *NotSingularError when more than one Answer ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AnswerQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{answer.Label}
	default:
		err = &NotSingularError{answer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AnswerQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Answers.
func (aq *AnswerQuery) All(ctx context.Context) ([]*Answer, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Answer, *AnswerQuery]()
	return withInterceptors[[]*Answer](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AnswerQuery) AllX(ctx context.Context) []*Answer {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Answer IDs.
func (aq *AnswerQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(answer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AnswerQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AnswerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AnswerQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AnswerQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AnswerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AnswerQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AnswerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AnswerQuery) Clone() *AnswerQuery {
	if aq == nil {
		return nil
	}
	return &AnswerQuery{
		config:                    aq.config,
		ctx:                       aq.ctx.Clone(),
		order:                     append([]answer.OrderOption{}, aq.order...),
		inters:                    append([]Interceptor{}, aq.inters...),
		predicates:                append([]predicate.Answer{}, aq.predicates...),
		withChoices:               aq.withChoices.Clone(),
		withSurveyQuestionAnswers: aq.withSurveyQuestionAnswers.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithChoices tells the query-builder to eager-load the nodes that are connected to
// the "choices" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AnswerQuery) WithChoices(opts ...func(*ChoiceQuery)) *AnswerQuery {
	query := (&ChoiceClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withChoices = query
	return aq
}

// WithSurveyQuestionAnswers tells the query-builder to eager-load the nodes that are connected to
// the "survey_question_answers" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AnswerQuery) WithSurveyQuestionAnswers(opts ...func(*SurveyQuestionAnswersQuery)) *AnswerQuery {
	query := (&SurveyQuestionAnswersClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withSurveyQuestionAnswers = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Answer.Query().
//		GroupBy(answer.FieldText).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (aq *AnswerQuery) GroupBy(field string, fields ...string) *AnswerGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AnswerGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = answer.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//	}
//
//	client.Answer.Query().
//		Select(answer.FieldText).
//		Scan(ctx, &v)
func (aq *AnswerQuery) Select(fields ...string) *AnswerSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AnswerSelect{AnswerQuery: aq}
	sbuild.label = answer.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AnswerSelect configured with the given aggregations.
func (aq *AnswerQuery) Aggregate(fns ...AggregateFunc) *AnswerSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AnswerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !answer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AnswerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Answer, error) {
	var (
		nodes       = []*Answer{}
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withChoices != nil,
			aq.withSurveyQuestionAnswers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Answer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Answer{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withChoices; query != nil {
		if err := aq.loadChoices(ctx, query, nodes,
			func(n *Answer) { n.Edges.Choices = []*Choice{} },
			func(n *Answer, e *Choice) { n.Edges.Choices = append(n.Edges.Choices, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withSurveyQuestionAnswers; query != nil {
		if err := aq.loadSurveyQuestionAnswers(ctx, query, nodes,
			func(n *Answer) { n.Edges.SurveyQuestionAnswers = []*SurveyQuestionAnswers{} },
			func(n *Answer, e *SurveyQuestionAnswers) {
				n.Edges.SurveyQuestionAnswers = append(n.Edges.SurveyQuestionAnswers, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AnswerQuery) loadChoices(ctx context.Context, query *ChoiceQuery, nodes []*Answer, init func(*Answer), assign func(*Answer, *Choice)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Answer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Choice(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(answer.ChoicesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.answer_choices
		if fk == nil {
			return fmt.Errorf(`foreign-key "answer_choices" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "answer_choices" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AnswerQuery) loadSurveyQuestionAnswers(ctx context.Context, query *SurveyQuestionAnswersQuery, nodes []*Answer, init func(*Answer), assign func(*Answer, *SurveyQuestionAnswers)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Answer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.SurveyQuestionAnswers(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(answer.SurveyQuestionAnswersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.answer_survey_question_answers
		if fk == nil {
			return fmt.Errorf(`foreign-key "answer_survey_question_answers" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "answer_survey_question_answers" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aq *AnswerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AnswerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(answer.Table, answer.Columns, sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, answer.FieldID)
		for i := range fields {
			if fields[i] != answer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AnswerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(answer.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = answer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AnswerGroupBy is the group-by builder for Answer entities.
type AnswerGroupBy struct {
	selector
	build *AnswerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AnswerGroupBy) Aggregate(fns ...AggregateFunc) *AnswerGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AnswerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AnswerQuery, *AnswerGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AnswerGroupBy) sqlScan(ctx context.Context, root *AnswerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AnswerSelect is the builder for selecting fields of Answer entities.
type AnswerSelect struct {
	*AnswerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AnswerSelect) Aggregate(fns ...AggregateFunc) *AnswerSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AnswerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AnswerQuery, *AnswerSelect](ctx, as.AnswerQuery, as, as.inters, v)
}

func (as *AnswerSelect) sqlScan(ctx context.Context, root *AnswerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}