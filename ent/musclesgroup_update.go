// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"todo/ent/exercise"
	"todo/ent/musclesgroup"
	"todo/ent/predicate"
	"todo/ent/schema/pksuid"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MusclesGroupUpdate is the builder for updating MusclesGroup entities.
type MusclesGroupUpdate struct {
	config
	hooks    []Hook
	mutation *MusclesGroupMutation
}

// Where appends a list predicates to the MusclesGroupUpdate builder.
func (mgu *MusclesGroupUpdate) Where(ps ...predicate.MusclesGroup) *MusclesGroupUpdate {
	mgu.mutation.Where(ps...)
	return mgu
}

// SetName sets the "name" field.
func (mgu *MusclesGroupUpdate) SetName(s string) *MusclesGroupUpdate {
	mgu.mutation.SetName(s)
	return mgu
}

// SetImage sets the "image" field.
func (mgu *MusclesGroupUpdate) SetImage(s string) *MusclesGroupUpdate {
	mgu.mutation.SetImage(s)
	return mgu
}

// AddExerciseIDs adds the "exercises" edge to the Exercise entity by IDs.
func (mgu *MusclesGroupUpdate) AddExerciseIDs(ids ...pksuid.ID) *MusclesGroupUpdate {
	mgu.mutation.AddExerciseIDs(ids...)
	return mgu
}

// AddExercises adds the "exercises" edges to the Exercise entity.
func (mgu *MusclesGroupUpdate) AddExercises(e ...*Exercise) *MusclesGroupUpdate {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return mgu.AddExerciseIDs(ids...)
}

// Mutation returns the MusclesGroupMutation object of the builder.
func (mgu *MusclesGroupUpdate) Mutation() *MusclesGroupMutation {
	return mgu.mutation
}

// ClearExercises clears all "exercises" edges to the Exercise entity.
func (mgu *MusclesGroupUpdate) ClearExercises() *MusclesGroupUpdate {
	mgu.mutation.ClearExercises()
	return mgu
}

// RemoveExerciseIDs removes the "exercises" edge to Exercise entities by IDs.
func (mgu *MusclesGroupUpdate) RemoveExerciseIDs(ids ...pksuid.ID) *MusclesGroupUpdate {
	mgu.mutation.RemoveExerciseIDs(ids...)
	return mgu
}

// RemoveExercises removes "exercises" edges to Exercise entities.
func (mgu *MusclesGroupUpdate) RemoveExercises(e ...*Exercise) *MusclesGroupUpdate {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return mgu.RemoveExerciseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mgu *MusclesGroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mgu.sqlSave, mgu.mutation, mgu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mgu *MusclesGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := mgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mgu *MusclesGroupUpdate) Exec(ctx context.Context) error {
	_, err := mgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mgu *MusclesGroupUpdate) ExecX(ctx context.Context) {
	if err := mgu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mgu *MusclesGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(musclesgroup.Table, musclesgroup.Columns, sqlgraph.NewFieldSpec(musclesgroup.FieldID, field.TypeString))
	if ps := mgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mgu.mutation.Name(); ok {
		_spec.SetField(musclesgroup.FieldName, field.TypeString, value)
	}
	if value, ok := mgu.mutation.Image(); ok {
		_spec.SetField(musclesgroup.FieldImage, field.TypeString, value)
	}
	if mgu.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   musclesgroup.ExercisesTable,
			Columns: musclesgroup.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mgu.mutation.RemovedExercisesIDs(); len(nodes) > 0 && !mgu.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   musclesgroup.ExercisesTable,
			Columns: musclesgroup.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mgu.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   musclesgroup.ExercisesTable,
			Columns: musclesgroup.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{musclesgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mgu.mutation.done = true
	return n, nil
}

// MusclesGroupUpdateOne is the builder for updating a single MusclesGroup entity.
type MusclesGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MusclesGroupMutation
}

// SetName sets the "name" field.
func (mguo *MusclesGroupUpdateOne) SetName(s string) *MusclesGroupUpdateOne {
	mguo.mutation.SetName(s)
	return mguo
}

// SetImage sets the "image" field.
func (mguo *MusclesGroupUpdateOne) SetImage(s string) *MusclesGroupUpdateOne {
	mguo.mutation.SetImage(s)
	return mguo
}

// AddExerciseIDs adds the "exercises" edge to the Exercise entity by IDs.
func (mguo *MusclesGroupUpdateOne) AddExerciseIDs(ids ...pksuid.ID) *MusclesGroupUpdateOne {
	mguo.mutation.AddExerciseIDs(ids...)
	return mguo
}

// AddExercises adds the "exercises" edges to the Exercise entity.
func (mguo *MusclesGroupUpdateOne) AddExercises(e ...*Exercise) *MusclesGroupUpdateOne {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return mguo.AddExerciseIDs(ids...)
}

// Mutation returns the MusclesGroupMutation object of the builder.
func (mguo *MusclesGroupUpdateOne) Mutation() *MusclesGroupMutation {
	return mguo.mutation
}

// ClearExercises clears all "exercises" edges to the Exercise entity.
func (mguo *MusclesGroupUpdateOne) ClearExercises() *MusclesGroupUpdateOne {
	mguo.mutation.ClearExercises()
	return mguo
}

// RemoveExerciseIDs removes the "exercises" edge to Exercise entities by IDs.
func (mguo *MusclesGroupUpdateOne) RemoveExerciseIDs(ids ...pksuid.ID) *MusclesGroupUpdateOne {
	mguo.mutation.RemoveExerciseIDs(ids...)
	return mguo
}

// RemoveExercises removes "exercises" edges to Exercise entities.
func (mguo *MusclesGroupUpdateOne) RemoveExercises(e ...*Exercise) *MusclesGroupUpdateOne {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return mguo.RemoveExerciseIDs(ids...)
}

// Where appends a list predicates to the MusclesGroupUpdate builder.
func (mguo *MusclesGroupUpdateOne) Where(ps ...predicate.MusclesGroup) *MusclesGroupUpdateOne {
	mguo.mutation.Where(ps...)
	return mguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mguo *MusclesGroupUpdateOne) Select(field string, fields ...string) *MusclesGroupUpdateOne {
	mguo.fields = append([]string{field}, fields...)
	return mguo
}

// Save executes the query and returns the updated MusclesGroup entity.
func (mguo *MusclesGroupUpdateOne) Save(ctx context.Context) (*MusclesGroup, error) {
	return withHooks(ctx, mguo.sqlSave, mguo.mutation, mguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mguo *MusclesGroupUpdateOne) SaveX(ctx context.Context) *MusclesGroup {
	node, err := mguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mguo *MusclesGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := mguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mguo *MusclesGroupUpdateOne) ExecX(ctx context.Context) {
	if err := mguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mguo *MusclesGroupUpdateOne) sqlSave(ctx context.Context) (_node *MusclesGroup, err error) {
	_spec := sqlgraph.NewUpdateSpec(musclesgroup.Table, musclesgroup.Columns, sqlgraph.NewFieldSpec(musclesgroup.FieldID, field.TypeString))
	id, ok := mguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MusclesGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, musclesgroup.FieldID)
		for _, f := range fields {
			if !musclesgroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != musclesgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mguo.mutation.Name(); ok {
		_spec.SetField(musclesgroup.FieldName, field.TypeString, value)
	}
	if value, ok := mguo.mutation.Image(); ok {
		_spec.SetField(musclesgroup.FieldImage, field.TypeString, value)
	}
	if mguo.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   musclesgroup.ExercisesTable,
			Columns: musclesgroup.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mguo.mutation.RemovedExercisesIDs(); len(nodes) > 0 && !mguo.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   musclesgroup.ExercisesTable,
			Columns: musclesgroup.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mguo.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   musclesgroup.ExercisesTable,
			Columns: musclesgroup.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MusclesGroup{config: mguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{musclesgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mguo.mutation.done = true
	return _node, nil
}
