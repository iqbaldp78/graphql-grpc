package graph

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mbizmarket/dmp/graphql/proto/users/pb"
)

//Resolver --
type Resolver struct {
	ResolverClientRadiance pb.RadianceServicesService
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*pb.Users, error) {
	res, err := r.ResolverClientRadiance.GetAllUsers(ctx, &empty.Empty{})
	// log.Println("res nya ", res)
	if err != nil {
		log.Println("err nya ", err)
		return []*pb.Users{}, err
	}
	return res.Data, nil
}
