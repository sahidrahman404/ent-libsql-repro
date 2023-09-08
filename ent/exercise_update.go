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

// ExerciseUpdate is the builder for updating Exercise entities.
type ExerciseUpdate struct {
	config
	hooks    []Hook
	mutation *ExerciseMutation
}

// Where appends a list predicates to the ExerciseUpdate builder.
func (eu *ExerciseUpdate) Where(ps ...predicate.Exercise) *ExerciseUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *ExerciseUpdate) SetName(s string) *ExerciseUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetImage sets the "image" field.
func (eu *ExerciseUpdate) SetImage(s string) *ExerciseUpdate {
	eu.mutation.SetImage(s)
	return eu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (eu *ExerciseUpdate) SetNillableImage(s *string) *ExerciseUpdate {
	if s != nil {
		eu.SetImage(*s)
	}
	return eu
}

// ClearImage clears the value of the "image" field.
func (eu *ExerciseUpdate) ClearImage() *ExerciseUpdate {
	eu.mutation.ClearImage()
	return eu
}

// SetHowTo sets the "how_to" field.
func (eu *ExerciseUpdate) SetHowTo(s string) *ExerciseUpdate {
	eu.mutation.SetHowTo(s)
	return eu
}

// SetNillableHowTo sets the "how_to" field if the given value is not nil.
func (eu *ExerciseUpdate) SetNillableHowTo(s *string) *ExerciseUpdate {
	if s != nil {
		eu.SetHowTo(*s)
	}
	return eu
}

// ClearHowTo clears the value of the "how_to" field.
func (eu *ExerciseUpdate) ClearHowTo() *ExerciseUpdate {
	eu.mutation.ClearHowTo()
	return eu
}

// AddMusclesGroupIDs adds the "muscles_groups" edge to the MusclesGroup entity by IDs.
func (eu *ExerciseUpdate) AddMusclesGroupIDs(ids ...pksuid.ID) *ExerciseUpdate {
	eu.mutation.AddMusclesGroupIDs(ids...)
	return eu
}

// AddMusclesGroups adds the "muscles_groups" edges to the MusclesGroup entity.
func (eu *ExerciseUpdate) AddMusclesGroups(m ...*MusclesGroup) *ExerciseUpdate {
	ids := make([]pksuid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return eu.AddMusclesGroupIDs(ids...)
}

// Mutation returns the ExerciseMutation object of the builder.
func (eu *ExerciseUpdate) Mutation() *ExerciseMutation {
	return eu.mutation
}

// ClearMusclesGroups clears all "muscles_groups" edges to the MusclesGroup entity.
func (eu *ExerciseUpdate) ClearMusclesGroups() *ExerciseUpdate {
	eu.mutation.ClearMusclesGroups()
	return eu
}

// RemoveMusclesGroupIDs removes the "muscles_groups" edge to MusclesGroup entities by IDs.
func (eu *ExerciseUpdate) RemoveMusclesGroupIDs(ids ...pksuid.ID) *ExerciseUpdate {
	eu.mutation.RemoveMusclesGroupIDs(ids...)
	return eu
}

// RemoveMusclesGroups removes "muscles_groups" edges to MusclesGroup entities.
func (eu *ExerciseUpdate) RemoveMusclesGroups(m ...*MusclesGroup) *ExerciseUpdate {
	ids := make([]pksuid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return eu.RemoveMusclesGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *ExerciseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *ExerciseUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *ExerciseUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *ExerciseUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *ExerciseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(exercise.Table, exercise.Columns, sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(exercise.FieldName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Image(); ok {
		_spec.SetField(exercise.FieldImage, field.TypeString, value)
	}
	if eu.mutation.ImageCleared() {
		_spec.ClearField(exercise.FieldImage, field.TypeString)
	}
	if value, ok := eu.mutation.HowTo(); ok {
		_spec.SetField(exercise.FieldHowTo, field.TypeString, value)
	}
	if eu.mutation.HowToCleared() {
		_spec.ClearField(exercise.FieldHowTo, field.TypeString)
	}
	if eu.mutation.MusclesGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   exercise.MusclesGroupsTable,
			Columns: exercise.MusclesGroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musclesgroup.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedMusclesGroupsIDs(); len(nodes) > 0 && !eu.mutation.MusclesGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   exercise.MusclesGroupsTable,
			Columns: exercise.MusclesGroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musclesgroup.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.MusclesGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   exercise.MusclesGroupsTable,
			Columns: exercise.MusclesGroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musclesgroup.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{exercise.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// ExerciseUpdateOne is the builder for updating a single Exercise entity.
type ExerciseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExerciseMutation
}

// SetName sets the "name" field.
func (euo *ExerciseUpdateOne) SetName(s string) *ExerciseUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetImage sets the "image" field.
func (euo *ExerciseUpdateOne) SetImage(s string) *ExerciseUpdateOne {
	euo.mutation.SetImage(s)
	return euo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (euo *ExerciseUpdateOne) SetNillableImage(s *string) *ExerciseUpdateOne {
	if s != nil {
		euo.SetImage(*s)
	}
	return euo
}

// ClearImage clears the value of the "image" field.
func (euo *ExerciseUpdateOne) ClearImage() *ExerciseUpdateOne {
	euo.mutation.ClearImage()
	return euo
}

// SetHowTo sets the "how_to" field.
func (euo *ExerciseUpdateOne) SetHowTo(s string) *ExerciseUpdateOne {
	euo.mutation.SetHowTo(s)
	return euo
}

// SetNillableHowTo sets the "how_to" field if the given value is not nil.
func (euo *ExerciseUpdateOne) SetNillableHowTo(s *string) *ExerciseUpdateOne {
	if s != nil {
		euo.SetHowTo(*s)
	}
	return euo
}

// ClearHowTo clears the value of the "how_to" field.
func (euo *ExerciseUpdateOne) ClearHowTo() *ExerciseUpdateOne {
	euo.mutation.ClearHowTo()
	return euo
}

// AddMusclesGroupIDs adds the "muscles_groups" edge to the MusclesGroup entity by IDs.
func (euo *ExerciseUpdateOne) AddMusclesGroupIDs(ids ...pksuid.ID) *ExerciseUpdateOne {
	euo.mutation.AddMusclesGroupIDs(ids...)
	return euo
}

// AddMusclesGroups adds the "muscles_groups" edges to the MusclesGroup entity.
func (euo *ExerciseUpdateOne) AddMusclesGroups(m ...*MusclesGroup) *ExerciseUpdateOne {
	ids := make([]pksuid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return euo.AddMusclesGroupIDs(ids...)
}

// Mutation returns the ExerciseMutation object of the builder.
func (euo *ExerciseUpdateOne) Mutation() *ExerciseMutation {
	return euo.mutation
}

// ClearMusclesGroups clears all "muscles_groups" edges to the MusclesGroup entity.
func (euo *ExerciseUpdateOne) ClearMusclesGroups() *ExerciseUpdateOne {
	euo.mutation.ClearMusclesGroups()
	return euo
}

// RemoveMusclesGroupIDs removes the "muscles_groups" edge to MusclesGroup entities by IDs.
func (euo *ExerciseUpdateOne) RemoveMusclesGroupIDs(ids ...pksuid.ID) *ExerciseUpdateOne {
	euo.mutation.RemoveMusclesGroupIDs(ids...)
	return euo
}

// RemoveMusclesGroups removes "muscles_groups" edges to MusclesGroup entities.
func (euo *ExerciseUpdateOne) RemoveMusclesGroups(m ...*MusclesGroup) *ExerciseUpdateOne {
	ids := make([]pksuid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return euo.RemoveMusclesGroupIDs(ids...)
}

// Where appends a list predicates to the ExerciseUpdate builder.
func (euo *ExerciseUpdateOne) Where(ps ...predicate.Exercise) *ExerciseUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *ExerciseUpdateOne) Select(field string, fields ...string) *ExerciseUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Exercise entity.
func (euo *ExerciseUpdateOne) Save(ctx context.Context) (*Exercise, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *ExerciseUpdateOne) SaveX(ctx context.Context) *Exercise {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *ExerciseUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *ExerciseUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *ExerciseUpdateOne) sqlSave(ctx context.Context) (_node *Exercise, err error) {
	_spec := sqlgraph.NewUpdateSpec(exercise.Table, exercise.Columns, sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Exercise.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, exercise.FieldID)
		for _, f := range fields {
			if !exercise.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != exercise.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(exercise.FieldName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Image(); ok {
		_spec.SetField(exercise.FieldImage, field.TypeString, value)
	}
	if euo.mutation.ImageCleared() {
		_spec.ClearField(exercise.FieldImage, field.TypeString)
	}
	if value, ok := euo.mutation.HowTo(); ok {
		_spec.SetField(exercise.FieldHowTo, field.TypeString, value)
	}
	if euo.mutation.HowToCleared() {
		_spec.ClearField(exercise.FieldHowTo, field.TypeString)
	}
	if euo.mutation.MusclesGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   exercise.MusclesGroupsTable,
			Columns: exercise.MusclesGroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musclesgroup.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedMusclesGroupsIDs(); len(nodes) > 0 && !euo.mutation.MusclesGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   exercise.MusclesGroupsTable,
			Columns: exercise.MusclesGroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musclesgroup.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.MusclesGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   exercise.MusclesGroupsTable,
			Columns: exercise.MusclesGroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musclesgroup.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Exercise{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{exercise.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}