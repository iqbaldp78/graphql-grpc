// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
package main

import (
	"context"
	"time"

	"github.com/mbizmarket/dmp/graphql/graph/generated"
	"github.com/mbizmarket/dmp/graphql/proto/users/pb"
)

type Resolver struct{}

func (r *mutationResolver) CreateTodo(ctx context.Context, input string) (string, error) {
	panic("not implemented")
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*pb.Users, error) {
	panic("not implemented")
}

func (r *usersResolver) LastLogin(ctx context.Context, obj *pb.Users) (*time.Time, error) {
	panic("not implemented")
}

func (r *usersResolver) CreatedAt(ctx context.Context, obj *pb.Users) (*time.Time, error) {
	panic("not implemented")
}

func (r *usersResolver) UpdatedAt(ctx context.Context, obj *pb.Users) (*time.Time, error) {
	panic("not implemented")
}

func (r *usersResolver) DeletedAt(ctx context.Context, obj *pb.Users) (*time.Time, error) {
	panic("not implemented")
}

func (r *usersResolver) LoginInToken(ctx context.Context, obj *pb.Users) (*string, error) {
	panic("not implemented")
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }
func (r *Resolver) Users() generated.UsersResolver       { return &usersResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type usersResolver struct{ *Resolver }
