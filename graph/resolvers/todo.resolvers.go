package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gqlgen-lambda/ent"
	"gqlgen-lambda/graph/generated"
	"gqlgen-lambda/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, todo model.TodoInput) (*ent.Todo, error) {
	return r.client.Todo.Create().
		SetText(todo.Text).
		SetStatus(todo.Status).
		SetNillablePriority(todo.Priority).
		SetNillableParentID(todo.Parent).
		Save(ctx)
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*ent.Todo, error) {
	return r.client.Todo.Query().
		All(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
