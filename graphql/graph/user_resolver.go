package graph

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/mbizmarket/dmp/graphql/graph/generated"
	"github.com/mbizmarket/dmp/graphql/proto/users/pb"
)

type usersResolver struct{ *Resolver }

//Users --
func (r *Resolver) Users() generated.UsersResolver { return &usersResolver{r} }

func (r *usersResolver) LastLogin(ctx context.Context, obj *pb.Users) (*time.Time, error) {
	if obj.LastLogin == nil {
		return nil, nil
	}
	data, err := ptypes.Timestamp(obj.LastLogin)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *usersResolver) CreatedAt(ctx context.Context, obj *pb.Users) (*time.Time, error) {
	if obj.CreatedAt == nil {
		return nil, nil
	}
	data, err := ptypes.Timestamp(obj.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *usersResolver) UpdatedAt(ctx context.Context, obj *pb.Users) (*time.Time, error) {
	if obj.UpdatedAt == nil {
		return nil, nil
	}
	data, err := ptypes.Timestamp(obj.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *usersResolver) DeletedAt(ctx context.Context, obj *pb.Users) (*time.Time, error) {
	if obj.DeletedAt == nil {
		return nil, nil
	}
	data, err := ptypes.Timestamp(obj.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *usersResolver) LoginInToken(ctx context.Context, obj *pb.Users) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}
