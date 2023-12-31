// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"todo/ent/exercise"
	"todo/ent/schema/pksuid"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Exercise is the model entity for the Exercise schema.
type Exercise struct {
	config `json:"-"`
	// ID of the ent.
	ID pksuid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Image holds the value of the "image" field.
	Image *string `json:"image,omitempty"`
	// HowTo holds the value of the "how_to" field.
	HowTo *string `json:"how_to,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExerciseQuery when eager-loading is set.
	Edges        ExerciseEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ExerciseEdges holds the relations/edges for other nodes in the graph.
type ExerciseEdges struct {
	// MusclesGroups holds the value of the muscles_groups edge.
	MusclesGroups []*MusclesGroup `json:"muscles_groups,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedMusclesGroups map[string][]*MusclesGroup
}

// MusclesGroupsOrErr returns the MusclesGroups value or an error if the edge
// was not loaded in eager-loading.
func (e ExerciseEdges) MusclesGroupsOrErr() ([]*MusclesGroup, error) {
	if e.loadedTypes[0] {
		return e.MusclesGroups, nil
	}
	return nil, &NotLoadedError{edge: "muscles_groups"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Exercise) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case exercise.FieldID:
			values[i] = new(pksuid.ID)
		case exercise.FieldName, exercise.FieldImage, exercise.FieldHowTo:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Exercise fields.
func (e *Exercise) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case exercise.FieldID:
			if value, ok := values[i].(*pksuid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case exercise.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case exercise.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				e.Image = new(string)
				*e.Image = value.String
			}
		case exercise.FieldHowTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field how_to", values[i])
			} else if value.Valid {
				e.HowTo = new(string)
				*e.HowTo = value.String
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Exercise.
// This includes values selected through modifiers, order, etc.
func (e *Exercise) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryMusclesGroups queries the "muscles_groups" edge of the Exercise entity.
func (e *Exercise) QueryMusclesGroups() *MusclesGroupQuery {
	return NewExerciseClient(e.config).QueryMusclesGroups(e)
}

// Update returns a builder for updating this Exercise.
// Note that you need to call Exercise.Unwrap() before calling this method if this Exercise
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Exercise) Update() *ExerciseUpdateOne {
	return NewExerciseClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Exercise entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Exercise) Unwrap() *Exercise {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Exercise is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Exercise) String() string {
	var builder strings.Builder
	builder.WriteString("Exercise(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	if v := e.Image; v != nil {
		builder.WriteString("image=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := e.HowTo; v != nil {
		builder.WriteString("how_to=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// NamedMusclesGroups returns the MusclesGroups named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Exercise) NamedMusclesGroups(name string) ([]*MusclesGroup, error) {
	if e.Edges.namedMusclesGroups == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedMusclesGroups[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Exercise) appendNamedMusclesGroups(name string, edges ...*MusclesGroup) {
	if e.Edges.namedMusclesGroups == nil {
		e.Edges.namedMusclesGroups = make(map[string][]*MusclesGroup)
	}
	if len(edges) == 0 {
		e.Edges.namedMusclesGroups[name] = []*MusclesGroup{}
	} else {
		e.Edges.namedMusclesGroups[name] = append(e.Edges.namedMusclesGroups[name], edges...)
	}
}

// Exercises is a parsable slice of Exercise.
type Exercises []*Exercise
