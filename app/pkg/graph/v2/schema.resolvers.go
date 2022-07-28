package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/roderm/example-plugin-system/app/ent/entdb"
	"github.com/roderm/example-plugin-system/app/pkg/graph/v2/generated"
	"github.com/roderm/example-plugin-system/app/pkg/graph/v2/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input entdb.CreateUserInput) (*entdb.User, error) {
	return r.DB.User.Create().CbSave(ctx, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input entdb.UpdateUserInput) (*entdb.User, error) {
	user, err := r.DB.User.Get(ctx, id)
	if err != nil {
		return user, err
	}
	return user.Update().SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*entdb.User, error) {
	user, err := r.DB.User.Get(ctx, id)
	if err != nil {
		return user, err
	}
	err = r.DB.User.DeleteOne(user).Exec(ctx)
	return user, err
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input entdb.CreateTodoInput) (*entdb.Todo, error) {
	return r.DB.Todo.Create().CbSave(ctx, input)
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id int, input entdb.UpdateTodoInput) (*entdb.Todo, error) {
	todo, err := r.DB.Todo.Get(ctx, id)
	if err != nil {
		return todo, err
	}
	return todo.Update().SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id int) (*entdb.Todo, error) {
	todo, err := r.DB.Todo.Get(ctx, id)
	if err != nil {
		return todo, err
	}
	err = r.DB.Todo.DeleteOne(todo).Exec(ctx)
	return todo, err
}

func (r *queryResolver) Users(ctx context.Context, filter *entdb.UserWhereInput, paging *model.Paging) (*entdb.UserConnection, error) {
	paging = CheckPaging(paging)
	opts := Options[entdb.UserPaginateOption]()
	opts = AppendIfNotNil(opts, filter, func() entdb.UserPaginateOption {
		return entdb.WithUserFilter(filter.Filter)
	})
	return r.DB.User.Query().Paginate(ctx, paging.After, paging.First, paging.Before, paging.Last, opts...)
}

func (r *queryResolver) Todos(ctx context.Context, filter *entdb.TodoWhereInput, paging *model.Paging) (*entdb.TodoConnection, error) {
	paging = CheckPaging(paging)
	opts := Options[entdb.TodoPaginateOption]()
	opts = AppendIfNotNil(opts, filter, func() entdb.TodoPaginateOption {
		return entdb.WithTodoFilter(filter.Filter)
	})
	return r.DB.Todo.Query().Paginate(ctx, paging.After, paging.First, paging.Before, paging.Last, opts...)
}

func (r *todoResolver) Desciption(ctx context.Context, obj *entdb.Todo) (string, error) {
	if obj == nil {
		return "", errors.New("not a todo")
	}
	return obj.Description, nil
}

func (r *userResolver) Todos(ctx context.Context, obj *entdb.User, filter *entdb.TodoWhereInput, paging *model.Paging) (*entdb.TodoConnection, error) {
	paging = CheckPaging(paging)
	opts := Options[entdb.TodoPaginateOption]()
	opts = AppendIfNotNil(opts, filter, func() entdb.TodoPaginateOption {
		return entdb.WithTodoFilter(filter.Filter)
	})
	return obj.QueryTodos().Paginate(ctx, paging.After, paging.First, paging.Before, paging.Last, opts...)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
