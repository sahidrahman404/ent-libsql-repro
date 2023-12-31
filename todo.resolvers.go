package todo

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"todo/ent"
	"todo/ent/schema/pksuid"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*ent.Todo, error) {
	return ent.FromContext(ctx).Todo.Create().SetInput(input).Save(ctx)
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id pksuid.ID, input ent.UpdateTodoInput) (*ent.Todo, error) {
	return ent.FromContext(ctx).Todo.UpdateOneID(id).SetInput(input).Save(ctx)
}

// CreateMusclesGroup is the resolver for the createMusclesGroup field.
func (r *mutationResolver) CreateMusclesGroup(ctx context.Context, input ent.CreateMusclesGroupInput) (*ent.MusclesGroup, error) {
	return ent.FromContext(ctx).MusclesGroup.Create().SetInput(input).Save(ctx)
}

// CreateExercise is the resolver for the createExercise field.
func (r *mutationResolver) CreateExercise(ctx context.Context, input ent.CreateExerciseInput) (*ent.Exercise, error) {
	return ent.FromContext(ctx).Exercise.Create().SetInput(input).Save(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
