// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ExercisesColumns holds the columns for the "exercises" table.
	ExercisesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "how_to", Type: field.TypeString, Nullable: true},
	}
	// ExercisesTable holds the schema information for the "exercises" table.
	ExercisesTable = &schema.Table{
		Name:       "exercises",
		Columns:    ExercisesColumns,
		PrimaryKey: []*schema.Column{ExercisesColumns[0]},
	}
	// MusclesGroupsColumns holds the columns for the "muscles_groups" table.
	MusclesGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "image", Type: field.TypeString},
	}
	// MusclesGroupsTable holds the schema information for the "muscles_groups" table.
	MusclesGroupsTable = &schema.Table{
		Name:       "muscles_groups",
		Columns:    MusclesGroupsColumns,
		PrimaryKey: []*schema.Column{MusclesGroupsColumns[0]},
	}
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}, Default: "IN_PROGRESS"},
		{Name: "priority", Type: field.TypeInt, Default: 0},
		{Name: "todo_parent", Type: field.TypeString, Nullable: true},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "todos_todos_parent",
				Columns:    []*schema.Column{TodosColumns[5]},
				RefColumns: []*schema.Column{TodosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ExerciseMusclesGroupsColumns holds the columns for the "exercise_muscles_groups" table.
	ExerciseMusclesGroupsColumns = []*schema.Column{
		{Name: "exercise_id", Type: field.TypeString},
		{Name: "muscles_group_id", Type: field.TypeString},
	}
	// ExerciseMusclesGroupsTable holds the schema information for the "exercise_muscles_groups" table.
	ExerciseMusclesGroupsTable = &schema.Table{
		Name:       "exercise_muscles_groups",
		Columns:    ExerciseMusclesGroupsColumns,
		PrimaryKey: []*schema.Column{ExerciseMusclesGroupsColumns[0], ExerciseMusclesGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "exercise_muscles_groups_exercise_id",
				Columns:    []*schema.Column{ExerciseMusclesGroupsColumns[0]},
				RefColumns: []*schema.Column{ExercisesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "exercise_muscles_groups_muscles_group_id",
				Columns:    []*schema.Column{ExerciseMusclesGroupsColumns[1]},
				RefColumns: []*schema.Column{MusclesGroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ExercisesTable,
		MusclesGroupsTable,
		TodosTable,
		ExerciseMusclesGroupsTable,
	}
)

func init() {
	TodosTable.ForeignKeys[0].RefTable = TodosTable
	ExerciseMusclesGroupsTable.ForeignKeys[0].RefTable = ExercisesTable
	ExerciseMusclesGroupsTable.ForeignKeys[1].RefTable = MusclesGroupsTable
}
