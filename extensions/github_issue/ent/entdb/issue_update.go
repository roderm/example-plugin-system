// Code generated by entc, DO NOT EDIT.

package entdb

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/roderm/example-plugin-system/extensions/github_issues/ent/entdb/issue"
	"github.com/roderm/example-plugin-system/extensions/github_issues/ent/entdb/predicate"
	"github.com/roderm/example-plugin-system/extensions/github_issues/ent/entdb/todo"
)

// IssueUpdate is the builder for updating Issue entities.
type IssueUpdate struct {
	config
	hooks    []Hook
	mutation *IssueMutation
}

// Where appends a list predicates to the IssueUpdate builder.
func (iu *IssueUpdate) Where(ps ...predicate.Issue) *IssueUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetURL sets the "url" field.
func (iu *IssueUpdate) SetURL(s string) *IssueUpdate {
	iu.mutation.SetURL(s)
	return iu
}

// SetIsPr sets the "is_pr" field.
func (iu *IssueUpdate) SetIsPr(b bool) *IssueUpdate {
	iu.mutation.SetIsPr(b)
	return iu
}

// SetNillableIsPr sets the "is_pr" field if the given value is not nil.
func (iu *IssueUpdate) SetNillableIsPr(b *bool) *IssueUpdate {
	if b != nil {
		iu.SetIsPr(*b)
	}
	return iu
}

// SetStatus sets the "status" field.
func (iu *IssueUpdate) SetStatus(s string) *IssueUpdate {
	iu.mutation.SetStatus(s)
	return iu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (iu *IssueUpdate) SetNillableStatus(s *string) *IssueUpdate {
	if s != nil {
		iu.SetStatus(*s)
	}
	return iu
}

// ClearStatus clears the value of the "status" field.
func (iu *IssueUpdate) ClearStatus() *IssueUpdate {
	iu.mutation.ClearStatus()
	return iu
}

// SetIssueID sets the "issue" edge to the Todo entity by ID.
func (iu *IssueUpdate) SetIssueID(id int) *IssueUpdate {
	iu.mutation.SetIssueID(id)
	return iu
}

// SetNillableIssueID sets the "issue" edge to the Todo entity by ID if the given value is not nil.
func (iu *IssueUpdate) SetNillableIssueID(id *int) *IssueUpdate {
	if id != nil {
		iu = iu.SetIssueID(*id)
	}
	return iu
}

// SetIssue sets the "issue" edge to the Todo entity.
func (iu *IssueUpdate) SetIssue(t *Todo) *IssueUpdate {
	return iu.SetIssueID(t.ID)
}

// Mutation returns the IssueMutation object of the builder.
func (iu *IssueUpdate) Mutation() *IssueMutation {
	return iu.mutation
}

// ClearIssue clears the "issue" edge to the Todo entity.
func (iu *IssueUpdate) ClearIssue() *IssueUpdate {
	iu.mutation.ClearIssue()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IssueUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IssueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("entdb: uninitialized hook (forgotten import entdb/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IssueUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IssueUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IssueUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *IssueUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   issue.Table,
			Columns: issue.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: issue.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: issue.FieldURL,
		})
	}
	if value, ok := iu.mutation.IsPr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: issue.FieldIsPr,
		})
	}
	if value, ok := iu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: issue.FieldStatus,
		})
	}
	if iu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: issue.FieldStatus,
		})
	}
	if iu.mutation.IssueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   issue.IssueTable,
			Columns: []string{issue.IssueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.IssueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   issue.IssueTable,
			Columns: []string{issue.IssueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{issue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// IssueUpdateOne is the builder for updating a single Issue entity.
type IssueUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IssueMutation
}

// SetURL sets the "url" field.
func (iuo *IssueUpdateOne) SetURL(s string) *IssueUpdateOne {
	iuo.mutation.SetURL(s)
	return iuo
}

// SetIsPr sets the "is_pr" field.
func (iuo *IssueUpdateOne) SetIsPr(b bool) *IssueUpdateOne {
	iuo.mutation.SetIsPr(b)
	return iuo
}

// SetNillableIsPr sets the "is_pr" field if the given value is not nil.
func (iuo *IssueUpdateOne) SetNillableIsPr(b *bool) *IssueUpdateOne {
	if b != nil {
		iuo.SetIsPr(*b)
	}
	return iuo
}

// SetStatus sets the "status" field.
func (iuo *IssueUpdateOne) SetStatus(s string) *IssueUpdateOne {
	iuo.mutation.SetStatus(s)
	return iuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (iuo *IssueUpdateOne) SetNillableStatus(s *string) *IssueUpdateOne {
	if s != nil {
		iuo.SetStatus(*s)
	}
	return iuo
}

// ClearStatus clears the value of the "status" field.
func (iuo *IssueUpdateOne) ClearStatus() *IssueUpdateOne {
	iuo.mutation.ClearStatus()
	return iuo
}

// SetIssueID sets the "issue" edge to the Todo entity by ID.
func (iuo *IssueUpdateOne) SetIssueID(id int) *IssueUpdateOne {
	iuo.mutation.SetIssueID(id)
	return iuo
}

// SetNillableIssueID sets the "issue" edge to the Todo entity by ID if the given value is not nil.
func (iuo *IssueUpdateOne) SetNillableIssueID(id *int) *IssueUpdateOne {
	if id != nil {
		iuo = iuo.SetIssueID(*id)
	}
	return iuo
}

// SetIssue sets the "issue" edge to the Todo entity.
func (iuo *IssueUpdateOne) SetIssue(t *Todo) *IssueUpdateOne {
	return iuo.SetIssueID(t.ID)
}

// Mutation returns the IssueMutation object of the builder.
func (iuo *IssueUpdateOne) Mutation() *IssueMutation {
	return iuo.mutation
}

// ClearIssue clears the "issue" edge to the Todo entity.
func (iuo *IssueUpdateOne) ClearIssue() *IssueUpdateOne {
	iuo.mutation.ClearIssue()
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IssueUpdateOne) Select(field string, fields ...string) *IssueUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Issue entity.
func (iuo *IssueUpdateOne) Save(ctx context.Context) (*Issue, error) {
	var (
		err  error
		node *Issue
	)
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IssueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("entdb: uninitialized hook (forgotten import entdb/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, iuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Issue)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from IssueMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IssueUpdateOne) SaveX(ctx context.Context) *Issue {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IssueUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IssueUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *IssueUpdateOne) sqlSave(ctx context.Context) (_node *Issue, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   issue.Table,
			Columns: issue.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: issue.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entdb: missing "Issue.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, issue.FieldID)
		for _, f := range fields {
			if !issue.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entdb: invalid field %q for query", f)}
			}
			if f != issue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: issue.FieldURL,
		})
	}
	if value, ok := iuo.mutation.IsPr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: issue.FieldIsPr,
		})
	}
	if value, ok := iuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: issue.FieldStatus,
		})
	}
	if iuo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: issue.FieldStatus,
		})
	}
	if iuo.mutation.IssueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   issue.IssueTable,
			Columns: []string{issue.IssueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.IssueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   issue.IssueTable,
			Columns: []string{issue.IssueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Issue{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{issue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
