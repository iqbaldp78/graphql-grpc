// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"fmt"

	"github.com/mbizmarket/dmp/graphql/graph/generated"
	"github.com/mbizmarket/dmp/graphql/graph/model"
)

func (r *entityResolver) FindRfqByID(ctx context.Context, id int) (*model.Rfq, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindUsersByID(ctx context.Context, id int) (*model.Users, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
