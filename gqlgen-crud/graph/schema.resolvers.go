package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen-crud/ent"
	"gqlgen-crud/graph/generated"
	"gqlgen-crud/graph/model"
	"math/rand"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

// CreateMember is the resolver for the createMember field.
func (r *mutationResolver) CreateMember(ctx context.Context, input model.NewMember) (*model.Member, error) {
	member := &model.Member{
		Name:   input.Name,
		Nick:   input.Nick,
		Team:   input.Team,
		Detail: input.Team,
		Img:    input.Img,
	}
	fmt.Println(member)
	conn := ent.ConnectDB()
	defer conn.Close()
	ent.CreateMember(ctx, conn, ent.Member{ID: 0, Name: member.Name, Nick: member.Nick, Team: member.Team, Detail: member.Detail, Img: member.Img})
	return member, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Members is the resolver for the members field.
func (r *queryResolver) Members(ctx context.Context) ([]*model.Member, error) {
	conn := ent.ConnectDB()
	defer conn.Close()
	//members := ent.QueryAllMember(ctx, conn)
	return r.members, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
