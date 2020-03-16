package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mbizmarket/dmp/graphql-aghanim/graph/model"
)

func (r *entityResolver) FindRfqsByID(ctx context.Context, id int) (*model.Rfqs, error) {
	panic(fmt.Errorf("not implemented"))
}

type entityResolver struct{ *Resolver }
